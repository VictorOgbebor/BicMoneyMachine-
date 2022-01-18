pragma solidity ^0.8.5;

contract TraderAccount {
    address internal traderAccount; // Account or address that trigger contracts

    modifier onlyTrader { // like onlyOwner
        require(msg.sender == traderAccount);
        _; // runs the lines
    }
// Constructor
    constructor() {
        traderAccount = msg.sender;
    }
}