package modules

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	StartDateTime *time.Time `xorm:"startDateTime" json:"startDateTime"`
	EndDateTime   *time.Time `xorm:"endDateTime" json:"endDateTime"`
}

func BitcoinRead(c *gin.Context) {

	req := request{}
	find := []BitcoinWallet{}
	res := []BitcoinWallet{}
	btc := 0.0
	var timeRoundHour time.Time
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}

	if req.StartDateTime == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, `end date time should be not null`)
		log.Panicln(`end date time should be not null`)
	}
	if req.EndDateTime == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, `start date time should be not null`)
		log.Panicln(`start date time should be not null`)
	}
	timeStartHour := *req.StartDateTime
	timeEndHour := *req.EndDateTime

	dx, err := DbConnect()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}

	ress, err := dx.QueryString(`SELECT sum(amount) from bitcoin_wallet where date_time < ?`, timeStartHour)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}
	btc, _ = strconv.ParseFloat(ress[0][`sum`], 64)

	err = dx.Table(`bitcoin_wallet`).Where(`date_time >= ? and date_time <= ?`, timeStartHour, timeEndHour).Asc(`date_time`).Find(&find)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}

	if len(find) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf(`not found transaction from [%v] to [%v]`, timeStartHour, timeEndHour))
		log.Panicf(`not found transaction from [%v] to [%v]`, timeStartHour, timeEndHour)
	}

	for k, v := range find {
		if k == 0 {
			timeRoundHour = *v.DateTime
		}
		if len(find) == k+1 {
			//last transaction
			timeResponse := time.Date(v.DateTime.Year(), v.DateTime.Month(), v.DateTime.Day(), v.DateTime.Hour(), 0, 0, 0, v.DateTime.Location())
			btcAmountHour := btc
			res = append(res, BitcoinWallet{
				Amount:   &btcAmountHour,
				DateTime: &timeResponse,
			})

		} else if (v.DateTime.After(timeStartHour) || v.DateTime.Equal(timeStartHour)) && v.DateTime.Hour() == timeRoundHour.Hour() {
			//transaction in same day same hour
			btc += *v.Amount
			timeRoundHour = *v.DateTime

		} else if (v.DateTime.After(timeStartHour) || v.DateTime.Equal(timeStartHour)) && v.DateTime.Hour() != timeRoundHour.Hour() && v.DateTime.Day() == timeRoundHour.Day() {
			//transaction in same day not same hour
			timeResponse := time.Date(v.DateTime.Year(), v.DateTime.Month(), v.DateTime.Day(), v.DateTime.Hour(), 0, 0, 0, v.DateTime.Location())
			btcAmountHour := btc
			res = append(res, BitcoinWallet{
				Amount:   &btcAmountHour,
				DateTime: &timeResponse,
			})
			btc += *v.Amount
			timeRoundHour = *v.DateTime

		} else if (v.DateTime.After(timeStartHour) || v.DateTime.Equal(timeStartHour)) && v.DateTime.Hour() != timeRoundHour.Hour() && v.DateTime.Day() != timeRoundHour.Day() {
			//transaction in next day
			timeResponse := time.Date(timeRoundHour.Year(), timeRoundHour.Month(), timeRoundHour.Day(), timeRoundHour.Hour()+1, 0, 0, 0, timeRoundHour.Location())
			btcAmountHour := btc
			res = append(res, BitcoinWallet{
				Amount:   &btcAmountHour,
				DateTime: &timeResponse,
			})
			btc += *v.Amount
			timeRoundHour = *v.DateTime

		}
	}

	c.JSON(http.StatusOK, gin.H{
		`btc`:  btc,
		`data`: res,
		`find`: find,
	})
}
