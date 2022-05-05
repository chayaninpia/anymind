package modules

import (
	"fmt"
	"log"
	"net/http"
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
	dx, err := DbConnect()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}

	err = dx.Table(`bitcoin_wallet`).Where(`date_time <= ?`, req.EndDateTime).Asc(`date_time`).Find(&find)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Panicln(err.Error())
	}

	if len(find) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf(`not found transaction from [%v] to [%v]`, req.StartDateTime, req.EndDateTime))
		log.Panicf(`not found transaction from [%v] to [%v]`, req.StartDateTime, req.EndDateTime)
	}

	for k, v := range find {
		if k == 0 {
			timeRoundHour = *req.StartDateTime
			btc = *v.Amount
		}
		if v.DateTime.Hour() < timeRoundHour.Hour() || v.DateTime.Hour() == timeRoundHour.Hour() {
			btc += *v.Amount
		} else if v.DateTime.Hour() > timeRoundHour.Hour() {
			timeResponse := time.Date(v.DateTime.Year(), v.DateTime.Month(), v.DateTime.Day(), v.DateTime.Hour()+1, 0, 0, 0, v.DateTime.Location())
			res = append(res, BitcoinWallet{
				Amount:   &btc,
				DateTime: &timeResponse,
			})
			timeRoundHour = *v.DateTime
		} else if len(find) == k+1 {
			timeResponse := time.Date(v.DateTime.Year(), v.DateTime.Month(), v.DateTime.Day(), v.DateTime.Hour()+1, 0, 0, 0, v.DateTime.Location())
			res = append(res, BitcoinWallet{
				Amount:   &btc,
				DateTime: &timeResponse,
			})
		}
	}

	c.JSON(http.StatusOK, res)
}
