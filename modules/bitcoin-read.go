package modules

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BitcoinRead(c *gin.Context) {

	req := gin.H{}
	res := []gin.H{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Panicln(err.Error())
	}

	dx, err := DbConnect()
	if err != nil {
		log.Panicln(err.Error())
	}

	query := `SELECT * FROM bitcoin_wallet WHERE datetime >= $1 AND datetime <= $2`
	row, err := dx.Query(query, req[`startDateTime`], req[`endDateTime`])
	if err != nil {
		log.Panicln(err.Error())
	}

	if err = row.Scan(&res); err != nil {
		log.Panicln(err.Error())
	}

	c.JSON(http.StatusOK, res)
}
