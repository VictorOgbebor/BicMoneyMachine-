
# Liquidation Bot using Flash Loans Spec(Details) = A design spec for a bot to liquidate underwater borrow positions using flash loans from our own JToken markets.
# https://docs.google.com/document/d/1k8GusDAk-dLO8heNG-d4YJkmx8Z8vVMsIfS1R6QeMUE/edit# 
 
# How Liquidations Work In a Nutshell
    Bob supplies 100 AVAX and 1,000 USDT.

    He borrows 4,000 USDC and 2,500 DAI. Imagine now that AVAX drops in price and his account is incurring a shortfall (i.e. it’s underwater).

    The close factor is 50% - this means you can liquidate by repaying 50% of a borrowed position.

    As a liquidator, you have several options:
        - Do I liquidate 4,000 USDC or 2,500 DAI?
        - Which of Bob’s collateral do I seize: AVAX or DAI?

    Let’s say you liquidate his 4,000 USDC. With a close factor of 50%, you can repay up to 2,000 USDC and seize an equivalent amount (i.e. $2,000) worth of AVAX or USDT collateral. 
        Which do you seize?
        - You can’t seize $2,000 worth of USDT collateral since Bob only has 1,000 USDT collateral.
        - So bot should seize $2,000 worth of AVAX.
        - If there are multiple options to seize collateral, then this becomes a fractional knapsack optimisation problem.

***************************************************************************************************************************************************************************
# Functions(Code)
    Steps
    1) Algorithm => 

        Get list of accounts that have health > 0 and < 1 using subgraph query: 
        (https://thegraph.com/hosted-service/subgraph/traderjoe-xyz/lending?query=underwater%20accounts )
        {
        accounts(where: {health_gt: 0, health_lt: 1, totalBorrowValueInUSD_gt: 0}) {
            id
            health
            totalBorrowValueInUSD
            totalCollateralValueInUSD
        }
    }

# When choosing which borrow position to liquidate and which collateral to seize as per above, we don’t need to make the optimal liquidation.
    2) Iterate through `tokens` and find a borrow position to repay.
    3) Iterate through `tokens` and find a supply position to seize. Seizable position must satisfy:
    4) supplyBalanceUnderlying > 0
    5) enterMarket == true (otherwise it’s not posted as collateral)
    6) Must have enough supplyBalanceUnderlying to seize 50% of borrow value
    7)Perform flash loan:
        - Borrow borrowed token to repay borrow position.
        - Redeem underlying seized tokens
        - Swap underlying seized tokens for AVAX using Trader Joe router
        - Calculate profit made after gas
    8) Log Transactions => 
        console.log(Tranaction and Snowtrace)
        Telegram/Discord/Twitter

# ** Flashloan Details **
    TJ use CREAM contracts for flash loans.
    Flash loan is done from the JWrappedNative (AVAX) or JCollateralCapErc20 (non-AVAX) contracts. => (Contracts found here: https://github.com/traderjoe-xyz/joe-lending )

    E.g. to make a flash loan of AVAX, you create a `FlashLoan` contract.(https://github.com/CreamFi/flashloan-playground/blob/main/contracts/FlashloanBorrower.sol)
        1) In one of the functions you make the flash loan of AVAX via JWrappedNative.flashLoan().
        2) Once the `FlashLoan` contract receives the AVAX, the onFlashLoan() method from the same contract is called 
            - this is where you put in logic to liquidate an AVAX borrow position, 
            - redeem underlying seize tokens, swap underlying seize tokens for AVAX 
            - return the flash loaned AVAX back to JWrappedNative.

    Flash loaning from JWrappedNative gives you WAVAX, not AVAX
    You cannot flash loan and liquidateBorrow the same token. Instead you will need to:
        ** Flash loan different token
        ** Swap token to the one you want to repay debt for
        ** Call liquidateBorrow
       
***************************************************************************************************************************************************************************
# Goal is to be the last resort liquidator, so most important is that the code is correct; we don’t need to optimise for speed.
    As for the bot itself, recommend using python or node.js and web3/ethers library to make the relevant smart contract calls.
        Well-written readme
            Transaction proof that it works (testnet or mainnnet)
            Unit tests written
            It always makes the optimal liquidation.
                Beat out all other liquidators by listening to last mined block and opportunistically liquidating accounts.

# Examples of how to liquidate Compound using AAVE flash loans:
https://gist.github.com/cryptofish7/f28165dea68c8aa2848926f2c027559d
https://gist.github.com/cryptofish7/dce64a23ada8609355456c7f959ec19f
https://gist.github.com/cryptofish7/c9dd14a323e555861fcf886e3162ecf9
https://gist.github.com/cryptofish7/05d5d1a5adc1703642305acad0dff73f

# CREAM flash loan docs:
https://docs.cream.finance/flash-loans-1/flash-loans-eip-3156

# Building a liquidation bot for Compound:
https://github.com/azizjalel/Compound-Watcher/blob/master/files/modules.py
https://blog.baowebdev.com/2019/11/how-to-build-a-compound-liquidation-bot/
https://www.quicknode.com/guides/defi/creating-a-restful-api-for-compound-finance-smart-contracts
