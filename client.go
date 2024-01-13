package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 3 * time.Second}

const (
	token               = "56e84c6bf80710059261f480d36e057e1cd735c191bcffc890a5177141867bd5"
	baseURL             = "https://rpc.ankr.com/multichain/%s/?ankr_%s="
	getWalletMethod     = "getAccountBalance"
	getTokenMethod      = "getTokenPrice"
	getCurrenciesMethod = "getCurrencies"
)

type GetWalletResponse struct {
	Result WalletResult `json:"result"`
}

type GetTokenPriceResponse struct {
	Result TokenResult `json:"result"`
}

type GetCurrenciesResponse struct {
	Result CurrenciesResult `json:"result"`
}

type CurrenciesResult struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
}

type TokenResult struct {
	USDPrice string `json:"usdPrice"`
}

type WalletResult struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	TokenSymbol     string `json:"tokenSymbol"`
	ContractAddress string `json:"contractAddress"`
	Balance         string `json:"balance"`
	BalanceUsd      string `json:"balanceUsd"`
}

func getWallet(address string) (*GetWalletResponse, error) {
	body := []byte(fmt.Sprintf(`{
	  "jsonrpc": "2.0",
	  "id": 1,
	  "method": "ankr_getAccountBalance",
	  "params": {
		"blockchain": "eth",
		"walletAddress": "%s"
	  }
	}`, address))

	resp := GetWalletResponse{}

	err := post(getWalletMethod, body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func getTokenData(contractAddress string) (*GetTokenPriceResponse, error) {
	body := []byte(fmt.Sprintf(`{
	  "jsonrpc": "2.0",
      "id":1,
	  "method": "ankr_getTokenPrice",
	  "params": {
		"blockchain": "eth",
		"contractAddress": "%s"
	  }
	}`, contractAddress))

	resp := GetTokenPriceResponse{}

	err := post(getTokenMethod, body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func getCurrenciesData() (*GetCurrenciesResponse, error) {
	body := []byte(`{
	  "jsonrpc": "2.0",
	  "id": 1,
	  "method": "ankr_getCurrencies",
	  "params": {
		"blockchain": "arbitrum"
	  }
	}`)

	resp := GetCurrenciesResponse{}

	err := post(getCurrenciesMethod, body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func post(method string, payload []byte, output any) error {
	url := fmt.Sprintf(baseURL, token, method)

	body := bytes.NewReader(payload)
	resp, err := client.Post(url, "application/json", body)
	if err != nil {
		return err
	}

	if code := resp.StatusCode; code != http.StatusOK {
		return errors.New(fmt.Sprintf("Ankr API returned unexpected status code: %d", code))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, output)

	if err != nil {
		return err
	}

	return nil
}
