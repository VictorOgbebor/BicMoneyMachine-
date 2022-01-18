pragma solidity ^0.8.5;
/* Contract Details:

    - Flashloan Trader Contract
    - deposit tradeAmount to trade (Play Money or Gas Funds) 
    - wont trigger unless trade is profitable trades
    - will deposit profits into bankContract(Earnings)
    - will log profitable trades 
*/

/*
AAVE lending Pool = 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
BankerJoe = 0xdc13687554205E5b89Ac783db14bb5bba4A1eDaC
BenqiTroller = 0x486Af39519B4Dc9a7fCcd318217352830E8AD9b4
=============================================================
usdc = 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e
usdt = 0xc7198437980c041c805a1edcba50c1ce5db95118
dai = 0xd586e7f844cea2f87f50152665bcbc2c279d8d70
wavax = 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7
link = 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7
mim = 0x130966628846bfd36ff31a822705796e8cb8c18d
frax = 0xd24c2ad096400b6fbcd2ad8b24e7acbc21a1da64
*/
import "./traderAccount.sol";

interface FrontOfficeInterface {
    function addTransactions(address _from, address _to, uint256 _amount) external payable;

}

contract traderContract is TraderAccount { 
FrontOfficeInterface FrontOffice = FrontOfficeInterface(0xD7ACd2a9FD159E69Bb102A1ca21C9a3e3A5F771B); // the address will change when we deploy mainnet
// Variables 
    // uint256 tradeTotal;
    Loan[] loan;
    Swap[] swap;

// Mappings
    mapping(address => uint256) traderBalance; // traderBalance // This will be the wallet balance
    mapping (address => Trader) public trader; 

    struct Trader {
        address id;
        uint256 balance;
    }
    struct Loan { // Asset A
        address lender;
        address assetA;
        uint256 loanAmount;
    }
    struct Swap { // Asset B
        address router;
        address assetB;
    }

// events
    event balanceUpdated(address indexed depositer, uint256 indexed amount);
    event sent(address indexed sender, uint256 indexed amount,  address indexed reciever);
    event staked(address indexed sender, uint256 indexed amount,  address indexed reciever);
    event deposited(address indexed sender, uint256 indexed amount);
    event loanSet(address indexed lender, uint256 indexed amount,  address indexed asset);
    event swapSet(address indexed sender, address indexed asset);

// Modifiers

    modifier profitTax(uint256 tax) { // profitTax
        require(msg.value >= tax);
        _; // runs the lines
    }

// If Else Statement...
    function whoIsTrading() public onlyTrader view returns(string memory){
      string memory greetTrader = "Hello Trader, Lets make some money";
        if(msg.sender == traderAccount){
            return greetTrader;
        } 
        else {
            return " Who are You?";
        }
    }

// For Loop 
    function countTrades(uint tradeId) public pure returns(uint) {
        for (uint i = 0; i < 10; i++){ 
            tradeId++;
        }
        return tradeId;
    }

// Setter (Like a Constructor)
    function setLoan(address _lender, address _assetA, uint256 _amountA) public {
        require(traderAccount == msg.sender, "Not allowed to access");

        Loan memory grabLoan = Loan(_lender, _assetA, _amountA);
        loan.push(grabLoan);

        emit loanSet(_lender, _amountA, _assetA);
    }

    function setSwap(address _router, address _assetB) public {
        require(traderAccount == msg.sender, "Not allowed to send");

        Swap memory sendSwap = Swap(_router, _assetB);
        swap.push(sendSwap);

        emit swapSet(_router, _assetB);
    }


// Getter info like API

    function getTrader() public view returns (address) {
        return traderAccount;
    }
    function getLoanInfo(uint256 _loanIndex) public view returns(address, address, uint256) {
         Loan memory loanDetails = loan[_loanIndex];
         return(loanDetails.lender, loanDetails.assetA, loanDetails.loanAmount);
    }

    function getSwapInfo(uint256 _swapIndex)public view returns(address, address){
         Swap memory swapDetails = swap[_swapIndex];
         return(swapDetails.router, swapDetails.assetB);
    }

    function getBalance() public view returns (uint256){
        return traderBalance[msg.sender];
    }



   // =========Actions=========
    function addTrader(address id, uint256 balance) public {
        trader[id] = Trader(id, balance);
    }

    function depositFunds() public payable onlyTrader returns (uint256){ // addBalaace
        traderBalance[msg.sender] += msg.value;
        emit deposited(msg.sender, msg.value);
        return traderBalance[msg.sender];
    }

    function withdrawFunds(uint256 amount) public onlyTrader returns (uint256){
        require(traderBalance[msg.sender] >= amount, "Not allowed to pull this much");
        traderBalance[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);        
        return traderBalance[msg.sender];
    }

    function updateBalance(uint256 amount) private onlyTrader returns (uint256){
        depositFunds();
        withdrawFunds(amount);
        emit balanceUpdated(msg.sender, amount);
        return traderBalance[msg.sender];
    }

        // Send to a other wallet
   function send(address reciever, uint256 amount) onlyTrader public { // this is calling the private call that is reusable
       require(traderBalance[msg.sender] >= amount, "insufficent funds!"); // Check msg.sender balance , 'insufficent funds'
       require(msg.sender != reciever, "Double sending dude!");

       uint256 prevBalance = traderBalance[msg.sender];
       _send(msg.sender, reciever, amount);

       assert(traderBalance[msg.sender] == prevBalance - amount);

       // Log Event and Additonal checks
       FrontOffice.addTransactions{value: 1 ether}(msg.sender, reciever, amount);
        emit sent(msg.sender, amount, reciever);
   }

    // Will send to Bank contract to earn Yields
      function stakeToBank(address bankContract, uint256 amount) onlyTrader public { // this is calling the private call that is reusable
       require(traderBalance[msg.sender] >= amount, "insufficent funds!"); // Check msg.sender balance , 'insufficent funds'
       require(msg.sender != bankContract, "Double sending dude!");

       uint256 prevBalance = traderBalance[msg.sender];
       _send(msg.sender, bankContract, amount);
       
       assert(traderBalance[msg.sender] == prevBalance - amount);
       // Log Event and Additonal checks
        emit staked(msg.sender, amount, bankContract);
   }

    // private function that gets inherited
   function _send(address from, address to, uint256 amount) onlyTrader private { // (Like a sub function) break logic to smaller code
       traderBalance[from] -= amount;
       traderBalance[to] += amount;
   }


// function payFees(){} // Trading Bot profit fee


}
