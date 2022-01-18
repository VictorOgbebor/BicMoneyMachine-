pragma solidity ^0.8.5;
/* 
*** A multisig wallet that will stores profits from trades 
– Anyone can deposit funds into the smart contract

– The contract creator should be able to input 
(1): the addresses of the owners and 
    - Ownable Contract
(2):  the numbers of approvals required for a transfer, in the constructor
 
For example, input 3 addresses and set the approval limit to 2. 
    - create Signer Contract that will sign off on transactions
        - Like Ownable 

– Anyone of the owners should be able to create a transfer demands. The creator of the transfer  will specify what amount and to what address the transfer will be made.

– Owners should be able to approve transfer requests.

– When a transfer request has the required approvals, the transfer should be sent. 
-----------------------------------------------------------------------------------
/* Functions
    // deposit public
        - Deposit-> Put funds in wallet
        - depositApproval = Multisig Requirements and Confirmations
        - updateBalance

    // assignSigner internal (chainlink VRF intergration)
        // Rotate address
        // require unique nft to identify  user

    // getSignersHistory public
        Function will display each cosigner and their approves

    // send public 
        Send -> Send Funds from account(Max is balance of contract owner)
        sendApproval = Multisig Requirements and Confirmations
        // require signatures = Require the Signature Modifier
        // UpdateBalance

    // withdrawFunds public 
        //require signatures
        Withdraw -> Pull funds from wallet
        withdrawApproval = Multisig Requirements and Confirmations

    // balanceOfwallet public
        balanceOfwallet -> shows how much a user has in there wallet
            Update with wallet activity
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
cosigner 1 = address 1
cosigner 2 = address 2
cosigner 3 = address 3
cosigner 4 = address 4
cosigner 4 = address 5
*/
// import "./TraderAccount.sol";

contract BankTreasury {
    address[] public traders; // mapping to owners
    uint limits;

// Mappings
    mapping(address => mapping(uint => bool)) cosigns; // the address is a cosigner who approves transfers
    mapping(address => uint256) bankBalance;

    BankTransfer[] transferDemands;

    struct BankTransfer {
        uint amount;
        address payable reciever;
        uint cosigns;
        bool transferSent;
        uint id;
    }
// Events 
    event TransferDemandActivated(uint indexed _id, uint _amount, address indexed _transferSender, address indexed _reciever);
    event CosignerApproved(uint _id, uint indexed _cosigns, address indexed _cosigner);
    event TransferConfirmed(uint indexed _id);

    modifier onlyTraders() { // like onlyOwner
        bool trader = false;
        for(uint i = 0; i < traders.length; i++){
            if(traders[i] == msg.sender) {
                trader = true;
        }
    }
    require(trader = true);
        _; // runs the lines
    }

    constructor(address[] memory _traders, uint _limits) {
        traders = _traders;
        limits = _limits;
    }

    function deposit() public payable onlyTraders {
    // Will deposit funds into the contract
    }

    
    function createTransfer(uint _amount, address payable _reciever) public {
    // Will fill out the BankTransfer struct w/ the info. Will add to the tradeDemands list
        emit TransferDemandActivated(transferDemands.length, _amount, msg.sender, _reciever);
        transferDemands.push(BankTransfer(_amount, _reciever, 0, false, transferDemands.length));

    }

    function approve(uint _id) public onlyTraders {
    // set approvals for transfer demands.
    // updates the mappings and approval. 
    // amount of approvals limit reached...funds will be released
    // No Double Vote for Admins and Demands confirmations
    // Once sent, deand is void to be voted on again
    require(cosigns[msg.sender][_id] == false); // Start false
    require(transferDemands[_id].transferSent == false);

    cosigns[msg.sender][_id] = true;
    transferDemands[_id].cosigns++;
    emit CosignerApproved(_id, transferDemands[_id].cosigns, msg.sender);
    if(transferDemands[_id].cosigns >= limits) {
        transferDemands[_id].transferSent = true;
        transferDemands[_id].reciever.transfer(transferDemands[_id].amount); 
        }
        emit TransferConfirmed(_id);

    }

    function getBalance() public view returns (uint256){
        return bankBalance[msg.sender];
    }

    // function getTransferDemands() public view returns (BankTransfer[] memory)
    // function getBalance() public view returns (BankTransfer[] memory)


}