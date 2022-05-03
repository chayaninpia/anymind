package modules

import (
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
		log.Panicln(err.Error())
	}

	dx, err := DbConnect()
	if err != nil {
		log.Panicln(err.Error())
	}

	err = dx.Table(`bitcoin_wallet`).Where(`datetime <= ?`, req.EndDateTime).Asc(`datetime`).Find(&find)
	if err != nil {
		log.Panicln(err.Error())
	}

	for k, v := range find {
		if k == 0 {
			timeRoundHour = *req.StartDateTime
		}
		if v.DateTime.Hour() < timeRoundHour.Hour() || v.DateTime.Hour() == timeRoundHour.Hour() {
			btc += *v.Amount
		} else if v.DateTime.Hour() > timeRoundHour.Hour() {
			timeRoundHour = time.Date(timeRoundHour.Year(), timeRoundHour.Month(), timeRoundHour.Day(), timeRoundHour.Hour(), 0, 0, 0, timeRoundHour.Location())
			res = append(res, BitcoinWallet{
				Amount:   &btc,
				DateTime: &timeRoundHour,
			})
			timeRoundHour = *v.DateTime
		} else if len(find) == k+1 {
			timeRoundHour = time.Date(timeRoundHour.Year(), timeRoundHour.Month(), timeRoundHour.Day(), timeRoundHour.Hour(), 0, 0, 0, timeRoundHour.Location())
			res = append(res, BitcoinWallet{
				Amount:   &btc,
				DateTime: &timeRoundHour,
			})
		}
	}

	c.JSON(http.StatusOK, res)
}
