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

	req := BitcoinWallet{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Panicln(err.Error())
	}
	dx, err := DbConnect()
	if err != nil {
		log.Panicln(err.Error())
	}

	if _, err := dx.InsertOne(&req); err != nil {
		log.Panicln(err.Error())
	}

	c.JSON(http.StatusOK, `Success`)
}
