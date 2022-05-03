package main

import (
	"app/modules"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST(`bitcoin`, modules.BitcoinCreate)
	r.GET(`bitcoin`, modules.BitcoinRead)

	r.Run(`:9997`)
}
