# Arbitrage Bot using Flash Swaps(Details) = Design an arbitrage bot using flash swaps between two AMMs
https://docs.google.com/document/d/13sfGbXdJl9gLHDQ-myG3XZgEQUHLQJIEon2qhE9fCvA/edit#

# How to Arbitrage between 2 AMMs
    We take the AVAX/USDT pool on Trader Joe and Pangolin as an example. 

    1) An arbitrage opportunity occurs if the price of AVAX is different between the two DEXes. 
        2) An arbitrageur would flash borrow AVAX from the lower priced DEX and sell it for USDT on the higher priced DEX 
            3) Then pay back the borrow using USDT (this is allowed for flash swaps - more on this below).
                4) Any remaining USDT is pocketed as profit.

- How Flash Swaps work
    1) Uniswap V2 allows flash swaps and actually, under the hood every swap is a flash swap. 
        2) This simply means that pair contracts send output tokens to the recipient before enforcing that enough input tokens have been received.
            3)Going back to our example AVAX/USDT pool, a user can flash borrow AVAX from the pool and pay back in USDT instead. 
                4) This is allowed so long the product of the new reserves matches the previous k.

***************************************************************************************************************************************************************************
# Functions(Code)
    Steps
    1) Using hardhat/brownie, listen to each block on the Avalanche C-chain [X]
        -read the prices of a given pair that is present on both DEXes (e.g AVAX/USDT). []
        - Beats out other arbitrage bots by listening to broadcasted trades []
             ** executes arbitrage opportunistically before broadcasted trade has confirmed **

    2) When the prices are off by a significant margin, do the following in a single transaction:
        - Borrow AVAX from the lower priced pair using flash swaps
        - Sell AVAX on the higher priced pair for USDT
        - Pay back USDT on the lower priced pair
        - Collect profit => convert to avax or stake joe for xJoe
    3) Log Transactions => 
        console.log(Tranaction and Snowtrace)
        Telegram/Discord/Twitter

    Minimum requirements:
        Code well-written and formatted
        Has unit tests
        Transaction proof that it works (either on mainnet or testnet)


# Resources
https://docs.uniswap.org/protocol/V2/guides/smart-contract-integration/using-flash-swaps
https://blog.infura.io/build-a-flash-loan-arbitrage-bot-on-infura-part-i/
https://blog.infura.io/build-a-flash-loan-arbitrage-bot-on-infura-part-ii/
https://github.com/Nafidinara/bot-pancakeswap?ref=https://githubhelp.com
https://github.com/pedrobergamini?tab=repositories
https://www.quicknode.com/guides/defi/how-to-interact-with-uniswap-using-javascript
https://messari.io/article/arbitraging-uniswap-and-sushiswap-in-node-js
