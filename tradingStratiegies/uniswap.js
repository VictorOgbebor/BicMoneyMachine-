
require('dotenv').config;
const ethers = require('ethers');


const addresses = {
    WETH: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', // or WAVAX || WMATIC || WCRO
    factory: '0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f', 
    router: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
    recipient: '0xf11b2fc4f28150517af11c2c456cbe75e976f663',
    CreatePair: ['event PairCreated(address indexed token0, address indexed token1, address pair, uint)']
}

const mnemonic = process.env.MNEMONIC
const provider = new ethers.providers.WebSocketProvider(process.env.ETH);
const wallet = ethers.Wallet.fromMnemonic(mnemonic);
const account = wallet.connect(provider);
const factory = new ethers.Contract(addresses.factory, addresses.CreatePair, account);

const router = new ethers.Contract(addresses.router,
    [
        'function getAmountsOut(uint amountIn, address[] memory path) public view returns (uint[] memory amounts)',
        'function swapExactTokensForTokens(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline) external returns (uint[] memory amounts)'
    ],
    account
);

factory.on('PairCreated', async (token0, token1, pairAddress) =>
{
    console.log(`
    New Pair Detected
    =================
    token0: ${token0}
    token1: ${token1}
    pairAddress: ${pairAddress}
    =================
    `);

    let tokenIn, tokenOut;
    if (token0 === WETH)
    {
        tokenIn = token0;
        tokenOut = token1;
    }
    if (token1 === WETH)
    {
        tokenIn = token1;
        tokenOut = token0;
    }

    if (typeof tokenIn === 'undefined')
    {
        return;
    }

    const amountIn = ethers.utils.parseUnits('0.1', 'ether');
    const amounts = await router.getAmountOut(amountIn, [tokenIn, tokenOut]);

    const amountOutMin = amounts[1].sub(amounts[1].div(10));

    console.log(`
    Buying new token 
    ++++++++++++++++
    tokenIn: ${amountIn.toString()} ${tokenIn} (WETH)
    tokenOut: ${amountOutMin.toString()} ${tokenOut}
    `);

    const tx = await router.swapExactTokensForTokens(amountIn, amountOutMin[tokenIn, tokenOut], addresses.recipient, Date.now() + 1000 * 60 * 10);

    const receipt = await tx.wait();
    console.log(`transaction Receipt: ${receipt}`)
})