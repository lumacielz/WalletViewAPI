package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	cacheCurrencies()

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

	destinationCurrencyAddress := getAddressBySymbol(destinationCurrency)
	token, err := getTokenData(destinationCurrencyAddress)
	if err != nil {
		gCtx.Error(err)
		gCtx.Status(500)
		return
	}

	resp := map[string]string{}
	for _, asset := range wallet.Result.Assets {
		amount := calculateValue(asset.BalanceUsd, token.Result.USDPrice)

		resp[asset.TokenSymbol] = fmt.Sprintf("%.3f %s", amount, destinationCurrency)
	}

	gCtx.JSON(200, resp)
}

func calculateValue(balanceUsd, priceUsd string) float64 {
	//TODO handle error
	balance, _ := strconv.ParseFloat(balanceUsd, 64)
	price, _ := strconv.ParseFloat(priceUsd, 64)

	return balance / price
}
