package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(getWallet(""))
	fmt.Println(getTokenData(""))
	//router := gin.Default()
	//
	//router.GET("/walletBalance", walletBalanceHandler)
	//
	//router.Run()
}

func walletBalanceHandler(gCtx *gin.Context) {
	//id := gCtx.Query("address")
	//destinationCurrency := gCtx.Query("currency")

	gCtx.JSON(200, "OK")
}
