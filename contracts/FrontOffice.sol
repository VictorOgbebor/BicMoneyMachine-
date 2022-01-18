pragma solidity ^0.8.5;

// Govenment Contract || Track Best yields and trades
contract FrontOffice {


// index Trades
    struct TradeTx {
        address from;
        address to;
        uint256 amount;
        uint txId;
    }

// Transaction[] transactionsLogs;
    TradeTx[] tradingBlocks;

    function addTransactions(address _from, address _to, uint256 _amount) external payable {
        tradingBlocks.push(TradeTx(_from, _to, _amount, tradingBlocks.length));   
    }

    function getTransactions(uint _index) public view returns(address, address, uint256){
        return (tradingBlocks[_index].from, tradingBlocks[_index].to, tradingBlocks[_index].amount);
    }

}