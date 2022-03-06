package main

import (
// "context"
// "fmt"
// "log"
// "math/big"
// "github.com/ethereum/go-ethereum/ethclient"
)

/*

Watches potential liquidations on TraderJoe
	- Learn about the Compound DeFi protocol.
	- Model expected return of liquidating underwater accounts.
	- Make some m-m-m-money.
	** Read https://medium.com/efficient-frontier/decentralized-finance-liquidations-a-business-opportunity-assessment-c0eea7bdacec || https://arxiv.org/pdf/1904.05234.pdf (FlashBoyz)

1) Create API that will manage bot
Create a RESTful API for Trader Joe Smart Contracts(
	https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
	https://tutorialedge.net/golang/creating-restful-api-with-golang/
	https://tutorialedge.net/golang/consuming-restful-api-with-go/
	https://tutorialedge.net/golang/authenticating-golang-rest-api-with-jwts/

	https://www.quicknode.com/guides/defi/creating-a-restful-api-for-compound-finance-smart-contracts || https://github.com/traderjoe-xyz/joe-api
 )


2) Monitor Underwater Accounts
Create Subgraph w/ Go and TheGraph
	+ Get list of accounts that have health > 0 and < 1 using subgraph query: (

		https://tutorialedge.net/golang/go-graphql-beginners-tutorial/
		https://tutorialedge.net/golang/go-graphql-beginners-tutorial-part-2/
		https://thegraph.com/hosted-service/subgraph/traderjoe-xyz/lending?query=underwater%20accounts || https://github.com/traderjoe-xyz/lending-subgraph
	)

3) Build a Liquidation Bot
( https://blog.baowebdev.com/2019/11/how-to-build-a-compound-liquidation-bot/ )
	- Full App example
		(https://github.com/l3a0/boron)
		(https://github.com/ckhatri/compound-liquidator || https://chiragkhatri.me/compound-liquidator/)

	# When choosing which borrow position to liquidate and which collateral to seize as per above, we don’t need to make the optimal liquidation.
    2) Iterate through `tokens` and find a borrow position to repay.
    3) Iterate through `tokens` and find a supply position to seize.

	Seizable position must satisfy:
    a) supplyBalanceUnderlying > 0
    b) enterMarket == true (otherwise it’s not posted as collateral)
    c) Must have enough supplyBalanceUnderlying to seize 50% of borrow value

    4)Perform flash loan:
        - Borrow borrowed token to repay borrow position.
        - Redeem underlying seized tokens
        - Swap underlying seized tokens for AVAX using Trader Joe router
        - Calculate profit made after gas

    5) Log Transactions =>
        console.log(Tranaction and Snowtrace)
        Telegram/Discord/Twitter

	- Peform Flashloan
	(
		https://docs.cream.finance/flash-loans-1/flash-loans-eip-3156
		https://gist.github.com/cryptofish7/f28165dea68c8aa2848926f2c027559d
		https://gist.github.com/cryptofish7/dce64a23ada8609355456c7f959ec19f
		https://gist.github.com/cryptofish7/c9dd14a323e555861fcf886e3162ecf9
		https://gist.github.com/cryptofish7/05d5d1a5adc1703642305acad0dff73f

	)

4) ** Build Accounts Service ( https://blog.baowebdev.com/2020/02/how-to-build-accounts-service-for-compound-finance/ ) **

** optional **


# Building a liquidation bot for Compound:
https://github.com/azizjalel/Compound-Watcher/blob/master/files/modules.py
https://www.quicknode.com/guides/defi/creating-a-restful-api-for-compound-finance-smart-contracts

https://blog.baowebdev.com/2019/11/how-to-build-a-compound-liquidation-bot/
https://blog.baowebdev.com/2020/02/how-to-build-accounts-service-for-compound-finance/
https://docs.google.com/document/d/1Jofloakadp3uPSS8G-qlBkQoyUIBYyGH8yrti7Wyc-g/edit#heading=h.58f5zcpzgclu

# Examples of how to liquidate Compound using AAVE flash loans:
https://gist.github.com/cryptofish7/f28165dea68c8aa2848926f2c027559d
https://gist.github.com/cryptofish7/dce64a23ada8609355456c7f959ec19f
https://gist.github.com/cryptofish7/c9dd14a323e555861fcf886e3162ecf9
https://gist.github.com/cryptofish7/05d5d1a5adc1703642305acad0dff73f

# CREAM flash loan docs:
https://docs.cream.finance/flash-loans-1/flash-loans-eip-3156


*/
