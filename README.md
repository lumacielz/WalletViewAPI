# WalletViewAPI

This simple API receives a walletAddress and a cryptocurrency in witch to display the assets. Example: /walletBalance?address=0x43a000734d61083c05e59355f7e40e2bd434577c&currency=ETH 

- All data are taken from the <b>Ankr API</b>. 
- First, the method <b>getCurrencies</b> is called to store the addresses of each currency symbol in a map, this data could be persisted at a database in the case but for the purpose and time available it was not. This step is necessary to get the contract address from the destination currency symbol that is needed in the next step.
- The exchange rate is taken from the <b>getTokenPrice</b> endpoint, that receives only the contract address. This rate is changeable but in case of a large scale API it could be cached and updated in appropriate time.
- The wallets are retrieved from <b>getAccountBalance</b> method and for each asset, the <b>usdBalance</b> is divided by the unit usd value of the destination cryptocurrency.

## Running
```
go run main.go
```
The server will be listening at localhost:8080.
