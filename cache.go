package main

var currencies = map[string]string{}

func cacheCurrencies() {
	data, _ := getCurrenciesData()
	for _, currency := range data.Result.Currencies {
		currencies[currency.Symbol] = currency.Address
	}
}

func getAddressBySymbol(wantSymbol string) string {
	for symbol, address := range currencies {
		if symbol == wantSymbol {
			return address
		}
	}

	return ""
}
