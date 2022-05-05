package modules

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BitcoinWallet struct {
	DateTime *time.Time `xorm:"datetime" json:"datetime"`
	Amount   *float64   `xorm:"amount" json:"amount"`
}

func BitcoinCreate(c *gin.Context) {

	bitcoin_wallet := BitcoinWallet{}
	if err := c.ShouldBindJSON(&bitcoin_wallet); err != nil {
		log.Panicln(err.Error())
	}

	if bitcoin_wallet.Amount == nil {
		log.Panicln(`amount should be not null`)
	}
	if bitcoin_wallet.DateTime == nil {
		log.Panicln(`date time should be not null`)
	}
	dx, err := DbConnect()
	if err != nil {
		log.Panicln(err.Error())
	}

	if _, err := dx.InsertOne(&bitcoin_wallet); err != nil {
		log.Panicln(err.Error())
	}

	c.JSON(http.StatusOK, `Success`)
}
