package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/bonedaddy/go-defi"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Connect to the Blockchain

var usdcAmount, avaxAmount = 100, 1

// Arrays List [Exchanges, Token]

// Exchange Method
// ABIs For Each Contract

// Token Method
// ABIs For Each Contract

type exchangeInfo struct {
	RouterAddress  string
	FactoryAddress string
}

type tokenInfo struct {
	TokenAddress string
}

type swapInfo struct {
	swapAmount      int
	flashloanAmount int
}

// Create the Input Funtions that will select the exchange info based on the exchange address
// Create the Input Funtions that will select the token info based on the token address
func main() {
	println("if new to this arb or using differnt tokens/exchanges...front load bot to approve tokens. Bot will fail if not")
	// InputInfo || Output to functions => https://dev.to/shellrean/go-fundamental-input-output-lb8

	blockchain()
	MonitorLatestPrices()

	var flashDeet swapInfo
	// the amount borrowed in AAVE or Joe
	flashDeet.flashloanAmount, flashDeet.swapAmount = usdcAmount, avaxAmount

	// the amount to swap in AVAX on Joe or other DEX

	executeTrade(flashDeet.flashloanAmount, flashDeet.swapAmount)

}

func blockchain() {
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block Header:", header.Number.String()) // 5671744
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block Number:", block.Number().Uint64())         // 5671744
	fmt.Println("Block Time:", block.Time())                      // 1527211625
	fmt.Println("Block Difficulty:", block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println("Block Hash:", block.Hash().Hex())                // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("Block Transactions:", len(block.Transactions())) // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count) // 144
}

// Step 1 => Int to String func
func MonitorLatestPrices() {

	/*
		Input | Output Actions
			- User Will Select The 2 Exchanges to Monitor
			- Next user can select Tokens in list or Paste Token Address
			- TokenAmount to Arb in USDC => convert to Native
			- Bot Will Go Approve tokens on both exchanges and start the Monitoring Prices
				call _approveStuff()

			- Create Range to select token and exchange
	*/
	var exchangeA exchangeInfo
	exchangeA.RouterAddress = ""
	exchangeA.FactoryAddress = ""

	var exchangeB exchangeInfo
	exchangeB.RouterAddress = ""
	exchangeB.FactoryAddress = ""

	// token we are arbing for or borrowing
	var swapToken tokenInfo
	swapToken.TokenAddress = ""

	// Usually WETH/WAVAX/WMATIC ==> Wrapped Native Token
	var exchangeToken tokenInfo
	exchangeToken.TokenAddress = ""

	_approveStuff()
	/*
		- Server that will Run Constantly in background Monitoring Blockchain
			--> Connect to Subqury and OnChain > will compare and contrast. Select the best Arb Senerio
			--> Store In DB for Backtesting and Sell Data in Chainlink Oracle

		Successful Arbs Checklist

			Reading Log Events * =>
				- Connect to APIs ** Oracles to compare on-Chain and off-Chain
				- LiquidityPool Added and Subtracted in Volume
				- Pair/Pool Contract Swap Event
				- avg SlipageCost per swap
				- Token Price Change % compared to each other
				- Estimate the Difference
			--> Whoever's token price has a higher price where we Sell. The Lower is where we Buy
			Send Info to estimateProfitAfterTradingFees()

				ExchangeAPrice = Buy Price
			 	ExchangeBPrice = Sell Price

				Determine if a trade is profitable or not?
					use funnyMath() to calulations
				return possibleTrade
	*/

	funnyMath()
}

// Step 2
func estimateProfitAfterTradingFees(ExchangeAPrice, ExchangeBPrice int) int {
	MonitorLatestPrices()

	if ExchangeAPrice < ExchangeBPrice {
		return ExchangeAPrice
	} else {
		return ExchangeBPrice
	}

}

// Step 3
func executeTrade(ExchangeAPrice, ExchangeBPrice int) int {
	// loop
	var profitableTrade = estimateProfitAfterTradingFees(1, 3) // =< profit should pay gas so it should be equal or greater than the gas being paid
	println(`Possible Profit: `, profitableTrade)

	var flashDeet swapInfo
	// the amount borrowed in AAVE or Joe
	flashDeet.flashloanAmount = usdcAmount

	// the amount to swap in AVAX on Joe or other DEX
	flashDeet.swapAmount = avaxAmount

	if profitableTrade <= 2 {
		_executeTrade(flashDeet.flashloanAmount, flashDeet.swapAmount)
		fmt.Println(flashDeet.flashloanAmount, flashDeet.swapAmount)
	}

	return profitableTrade

}

// Internal = approve tokens on exchanges if not already
func _approveStuff() {
	// Approve Exchanges
	// Approve Tokens to arb before app works
}

// Internal = Does Trade LogicBased
func _executeTrade(int, int) {

	/*
		Input Output Action
			- User Will Select The 2 Exchanges to Monutio
			- Next user can select Tokens in list or Paste Token Address
			- If Flashloan...Select provider => AAVE or TraderJoe
			- Bot Will Go Approve tokens on both exchanges and start the Arb
	*/

	// check if tokens are approved if not already
	_approveStuff()

	// Create Slice that will change order based on arbitarge order
	// Log Each Transaction : Send Notifications

}

func funnyMath() { // Recursion
	// Calculate Slippage
	/*
		if != Platypus
			1a) the ExchangeA and ExchangeB trade fee (0.3% on each) => the price difference after a trade fee:
		else
			1b) the ExchangeA and Platypus trade fee (0.3% on amm) and (0.01% on ptp) => the price difference after a trade fee:

	// Calculate Gas
		2) the Avalanche transaction fee (cheap on avax).
		3) Slippage + Gas
	*/
}

/*
// High-Level Overview

subscribeTo(Exchange, "usdc", "avax", (latestTrade) =>
	{
		latestExchangePrice = latestTrade.price

		if (profitableTrade(latestExchangePrice, latestPlatapusPrice)) {
			executeTRADE("usdc", "avax")
		}
	}

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

# (1) Monitor latest prices on Exchange and Platypus #

*create structs *
	Token:
	=> TokenAddress
	=> ERC20 ABI

	Exchange:
	=> Exchange address
	=> Exchange ABI

	Platypus Pool
	https://snowtrace.io/address/0xe2d3eb21db91f6a8dde5dac9fde7f546aa7590ba#code

	JoeRouter
	https://snowtrace.io/address/0x60ae616a2155ee3d9a68541ba4544862310933d4#code

	PNGRouter
	https://snowtrace.io/address/0xefa94de7a4656d787667c749f7e1223d71e9fd88#code

	pairAddress:
	=> { USDC/WAVAX_Exchange || USDT/WAVAX_Exchange || DAI/WAVAX_Exchange }
	=> { JoePair || PNGPair || SushiPair }

	Platapus:
	=> Exchange address
	=> Exchange ABI


* Reading Log Events * => https://goethereumbook.org/event-read/ || https://goethereumbook.org/event-read-erc20/ || "emit" events during execution

	Swap Events => Pair Contract per AMM ... [ event Swap(address indexed sender, uint amount0In, uint amount1In, uint amount0Out, uint amount1Out, address indexed to) ]
	Swap Event => Platypus ...[ event Swap(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address indexed to) ]
	ERC20 Transfer Events => Per Token and Amount ... [event Transfer(address indexed from, address indexed to, uint value)  ]


(Example Code)
contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
query := ethereum.FilterQuery{
  Addresses: []common.Address{contractAddress},
}

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

# (2) Decide whether to trade #

- Based on the above function we will
- Determine if a trade is profitable or not?

There are 3 (math) factors:
1a) the ExchangeA and ExchangeB trade fee (0.3% on each) => the price difference after a trade fee:
1b) the ExchangeA and Platypus trade fee (0.3% on amm) and (0.01% on ptp) => the price difference after a trade fee:
2) the Avalanche transaction fee (cheap on avax).
3) Slippage on the Uniswap market and slippage on the SushiSwap market

(Example Code)
function estimateProfitAfterTradingFees(ExchangeAPrice, ExchangeBPrice) {

	// Finds the price profit in trade
	const diff = Math.abs(ExchangeBPrice - ExchangeAPrice);
	const diffRatio = diff / Math.max(ExchangeBPrice, ExchangeAPrice);

	// multiply by 2 because we trade 2 times  (once on ExchangeA and once on ExchangeB)
	const fees = Exchange_FEE * 2;
	return diffRatio - fees;
}

(Example Code)
function estimateProfitAfterTradingFees(ExchangeAPrice, PlatypusPrice) {

	// Finds the price profit in trade

	const diff = Math.abs(PlatypusPrice - ExchangeAPrice);
	const diffRatio = diff / Math.max(PlatypusPrice, ExchangeAPrice);

	// multiply by 2 because we trade 2 times  (once on ExchangeA and once on ExchangeB)
	const ptpfees = Platypus_FEE;
	const fees = Exchange_FEE ;
	return diffRatio - (fees + ptpfees)
	}
}

If the profit after trading fees is greater than $0.01 USD equivalent, should we do the trade? No, because the Ethereum transaction fee (gas) will probably cost more.

(Example Code)
	if (ExchangeA <= $1.00 && estimateProfitAfterTradingFees(ExchangeAPrice, PlatypusPrice)){
		Println("Is Profitable")
	}
OK, what if the profit was $4.01, should we do it then? Yes, if the amount we’re buying doesn’t move the price.
OK, how do I know if it moves the price? You calculate the slippage, which can be derived from the size of both “reserves” (liquidity).

PairContract: { token0Reserves: 400, token1Reserves: 1 } = Those numbers represent the number of tokens in this smart contract,

liquidity = {
	token0Reserve = USDC
	token1Reserve = WAVAX
}

use the 3 methods: depositLiquidity, withdrawLiquidity, swap.

- To get the price of token1 simply find the ratio: 400/1 or 400. To get the price of token0, take the inverse of the ratio: 1/400 or 0.0025.
- These AMMs are 2-way: a user can buy token0 selling token1, or buy token1 selling token0.

Back to the point, how do we calculate the slippage?
We’ll use the relationship between the constant product of 400 and reserve sizes to see prices at various percentages of the supply of token1 reserves.

- calculate price of token1 after buying 50% of token1’s supply,
- we’ll solve for how many units of token0 would need to exist to maintain the constant product of 400
- if only 0.5 units (50% of the original quantity of 1) of token1 exist.

*Understand Liquidty Pool and theGraph it*

(constant product) = token0 reserves * token1 reserves;
	400 = token0Reserves * (1*0.5)
	Solve for token0Reserves: 400 = 0.5 * token0Reserves
	400/0.5 = 800

After buying 50% of token1:
800 units of token0 and 0.5 units of token1

(Example Code)
const profitRate = estimateProfitAfterTradingFees(ExchangeAPrice, ExchangeBPrice); // Calls above Function
const maxBet = findMaxBet(profitRate, exchangeAReserves, exchangeAReserves);
const expectedProfit = maxBet * profitRate;

if (expectedProfit > 0) {
	executeTrade(maxBet);
}


BUT this trade will fail to complete 100% of the time Because
	- either a competing arb bot will make the trade for less profit
	- generic front running bot will clone your transaction with a higher gas price.

Solutions:
1) allocate 100% of estimatedProfit to gas and then reduce it until transactions start failing (competing bots).
2) Open Research to improve...Possible Machine Learning

(Example Code)-(Check go-eth docs)
function getGasPrice(n = 1) {
	const fixedFee = 0.0001 * n;
	const gasPrice = (expectedProfit - fixedFee) / ESTIMATED_GAS_USAGE;
}

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

# (3) Execute trade #
Before we can execute a "swap" on Uniswap or SushiSwap we need to call the "approve" method on each individual ERC20 you want to trade with each exchange you want to trade.
For our scenario, we'll need 4 approvals: [ Make Sure we approve tokens, Create case loop the will approve tokens if not Approved ]

(Example Code)
const exchangeARouterAddress = "0x7a250d5630b4cf539739df2c5dacb4c659f2488d";
const exchangeARouterAdress = "0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f";
const usdcErc20Address = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48";
const wavaxErc20Address = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2";

// allow ExchangeA and ExchangeA to move up to 1000.0 of my units of USDC
approveExchangeA(usdcErc20Address, 1000.0);
approveExchangeB(usdcErc20Address, 1000.0);

// allow ExchangeA and ExchangeA to move up to 5 of my units of ETH
approveExchangeA(wavaxErc20Address, 5.0);
approveExchangeB(wavaxErc20Address, 5.0);const gasPriceGwei = "100";

const gasPriceWei = ethers.utils.parseUnits(gasPriceGwei, "gwei"); // in GWEI
const wallet = new ethers.Wallet( Buffer.from( "", // paste your private key in "hex" ));
const signer = wallet.connect(provider);

function approveExchangeA( erc20Address, amountToApproveInNativeUnitFloat) {
	const erc20Contract = new ethers.Contract(erc20Address, erc20Abi, signer);

	return erc20Contract.decimals().then((decimals) => {
		return erc20Contract.approve( ExchangeARouterAddress, ethers.utils.parseUnits(`${amountToApproveInNativeUnitFloat}`, decimals ),

		{ gasLimit: 100000, gasPrice: gasPriceWei } ); 		// manually set gas price since ethers.js can't estimate
		});
	}


With the approvals done, we can finally execute a trade:
(Example Code)
const uniswapRouterAbi = [ "function swapExactTokensForTokens(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline) external returns (uint[] memory amounts)",];

//  BuyEthForUsdc

function buyEthWithUsdc(amountUsdcFloat) {
	const exchangeContract = new ethers.Contract(uniswapRouterAddress, uniswapRouterAbi, signer) // usdc uses 6 decimals
	return exchangeContract.swapExactTokensForTokens(
		ethers.utils.parseUnits(`${amountUsdcFloat}`, 6),
		ethers.utils.parseUnits(`${amountUsdcFloat}`, 6, // this is the expected minimum output
		[usdcErc20Address, wethErc20Address], // notice the ordering of this array, give usdc, get weth wallet.address,
		createDeadline(),
		Math.floor(Date.now() / 1000) + 20 createGasOverrides(){ gasLimit: ethers.utils.hexlify(300000), gasPrice: gasPriceWei });

	}

// SellEthForUsdc

	function buyUsdcWithEth(amountEthFloat) {
		const exchangeContract = new ethers.Contract(uniswapRouterAddress, uniswapRouterAbi, signer) // eth uses 18 decimals
		return exchangeContract.swapExactTokensForTokens(
			ethers.utils.parseUnits(`${amountEthFloat}`, 18), 0,
			[wethErc20Address, usdcErc20Address], // notice the ordering of this array: give weth, get usdc wallet.address,
			createDeadline(), //
			Math.floor(Date.now() / 1000) + 20 createGasOverrides(){ gasLimit: ethers.utils.hexlify(300000), gasPrice: gasPriceWei } );
	}

	// Cut flashloan fee
	// Send Profit to Wallet.
	// We’d also need to hide transactions to prevent generic front running bots

https://docs.platypus.finance/platypus-finance-docs/developers/contracts
https://uniswap.org/developers
https://docs.google.com/document/d/15h9Eu5jHl92qZ_Mcu4Bpy2W4lxQjHUF-D-HSPlEIqhs/edit
https://messari.io/article/arbitraging-uniswap-and-sushiswap-in-node-js || https://gist.github.com/jotto/81ef912e3db07b60ac643b778714c38f

*/
