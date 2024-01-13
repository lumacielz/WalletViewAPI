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

var client = http.Client{Timeout: 10 * time.Second}

const (
	token           = "56e84c6bf80710059261f480d36e057e1cd735c191bcffc890a5177141867bd5"
	baseURL         = "https://rpc.ankr.com/multichain/%s/?ankr_%s="
	getWalletMethod = "getAccountBalance"
	getTokenMethod  = "getTokenPrice"
)

type GetWalletResponse struct {
	Result WalletResult `json:"result"`
}

type GetTokenPriceResponse struct {
	Result TokenResult `json:"result"`
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
	body := []byte(`{
	  "jsonrpc": "2.0",
	  "id": 1,
	  "method": "ankr_getAccountBalance",
	  "params": {
		"blockchain": "eth",
		"walletAddress": "0x43a000734d61083c05e59355f7e40e2bd434577c"
	  }
	}`)

	resp := GetWalletResponse{}

	err := post(getWalletMethod, body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func getTokenData(contractAddress string) (*GetTokenPriceResponse, error) {
	body := []byte(`{
	  "jsonrpc": "2.0",
      "id":1,
	  "method": "ankr_getTokenPrice",
	  "params": {
		"blockchain": "eth",
		"contractAddress": "0xdac17f958d2ee523a2206206994597c13d831ec7"
	  }
	}`)
	resp := GetTokenPriceResponse{}

	err := post(getTokenMethod, body, &resp)
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
