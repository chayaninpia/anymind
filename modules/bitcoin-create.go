package modules

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BitcoinCreate(c *gin.Context) {

	req := gin.H{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Panicln(err.Error())
	}
	dx, err := DbConnect()
	if err != nil {
		log.Panicln(err.Error())
	}

	query := `INSERT INTO bitcoin_wallet VALUES ($1,$2)`
	if _, err = dx.Exec(query, req[`datetime`], req[`amount`]); err != nil {
		log.Panicln(err.Error())
	}

	c.JSON(http.StatusOK, `Success`)
}
