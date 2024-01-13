package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/walletBalance", walletBalanceHandler)

	router.Run()
}

func walletBalanceHandler(gCtx *gin.Context) {
	id := gCtx.Query("address")
	destinationCurrency := gCtx.Query("currency")

	wallet, err := getWallet(id)
	if err != nil {
		gCtx.Error(err)
		gCtx.Status(500)
		return
	}

	token, err := getTokenData(destinationCurrency)
	if err != nil {
		gCtx.Error(err)
		gCtx.Status(500)
		return
	}

	resp := "%s: %.3f %s"
	for _, asset := range wallet.Result.Assets {
		amount := calculateValue(asset.BalanceUsd, token.Result.USDPrice)

		resp = fmt.Sprintf(resp, asset.TokenSymbol, amount, destinationCurrency)
	}

	gCtx.JSON(200, "OK")
}

func calculateValue(balanceUsd, priceUsd string) float64 {
	//TODO handle error
	balance, _ := strconv.ParseFloat(balanceUsd, 64)
	price, _ := strconv.ParseFloat(priceUsd, 64)

	return balance / price
}
