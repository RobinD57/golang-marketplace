pragma solidity ^0.6.0;

import "./AuthenticityOracleInterface.sol";
import "./installed_contracts/zeppelin/contracts/payment/Escrow.sol";

contract MadEscrow is Escrow {
    string public name = "Solidity MAD escrow contract";

    // Oracle parameters
    AuthenticityOracleInterface private oracleInstance;
    address private oracleAddress;
    bool private authenticityCheck;
    mapping(uint256=>bool) myRequests;

    // Oracle events
    event newOracleAddressEvent(address oracleAddress);
    event ReceivedNewRequestIdEvent(uint256 id);
    event AuthenticityCheckCompletedEvent(bool authenticityCheck, uint256 id);

    // Escrow parameters
    address public buyer;
    address public seller;
    uint public purchaseAmount = 0; //Total paid by BUYER

    // Buyer required escrow funding (2x purchaseAmount)
    uint public buyerRequiredEscrow = 0;
    // Seller required escrow funding (1x purchaseAmount)
    uint public sellerRequiredEscrow = 0;
    // TOTAL required escrow funding (3x purchaseAmount)
    uint public totalEscrowAmount = 0;

    // Buyer escrow funding received
    uint public buyerEscrowedFunds = 0;
    // Seller escrow funding received
    uint public sellerEscrowedFunds = 0;

    // This will be true when the buyer has sent 2x the purchaseAmount
    // and when buyerEscrowedFunds >= buyerRequiredEscrow
    bool public buyerHasFunded = false;
    // This will be true when the seller has sent 1x the purchaseAmount
    // True when sellerEscrowedFunds >= sellerRequiredEscrow
    bool public sellerHasFunded = false;

    // True when both buyerHasFunded = true AND sellerHasFunded = true
    bool public escrowFullyFunded = false;


    // Once FullyFunded, both parties must finalize the trade/sale
    bool public buyerfinalized = false;
    bool public sellerfinalized = false;

    // escrow events
    event Terms(uint purchaseAmount, uint totalEscrowAmount, uint BuyerRequiredFunding, uint SellerRequiredFunding);
    event Initiated(address initiator, address buyer, address seller, uint tradeFundsforXfer);
    event FundsReceived(address sender, uint amount);
    event PartyFinalized(address party);
    event EscrowComplete(bool completed, uint buyerReceives, uint sellerReceives);

    function setOracleAddress(address _oracleAddress) public onlyOwner {
            oracleAddress = _oracleAddress;
    }

    function setOracleInstanceAddress (address _oracleInstanceAddress) public onlyOwner {
        oracleAddress = _oracleInstanceAddress;
        oracleInstance = AuthenticityOracleInterface(oracleAddress);
        emit newOracleAddressEvent(oracleAddress);
    }

    function updateAuthenticityCheck() public {
      uint256 id = oracleInstance.authenticityCheck();
      myRequests[id] = true;
      emit ReceivedNewRequestIdEvent(id);
    }

    function callback(bool _authenticityCheck, uint256 _id) public onlyOracle {
      require(myRequests[_id], "Request not in pending list.");
      authenticityCheck = _authenticityCheck;
      delete myRequests[_id];
      emit AuthenticityCheckCompletedEvent(_authenticityCheck, _id);
    }

    modifier onlyOracle() {
     require(msg.sender == oracleAddress, "You are not authorized to call this function.");
      _;
    }


    function initEscrow(address _buyer, address _seller, uint _purchaseAmount) public {
        require(_seller != address(0));
        require(_buyer != address(0));
        require(_buyer != _seller);

        // require the contract creator to be either _buyer or _seller - Check whether that's the best approach!
        require(msg.sender == _seller || msg.sender == _buyer);
        require(_purchaseAmount > 0);

        escrowFullyFunded = false;
        buyer = _buyer;
        seller = _seller;
        initiated = true;

        buyerRequiredEscrow = _purchaseAmount * 2;
        sellerRequiredEscrow = _purchaseAmount;
        totalEscrowAmount = _purchaseAmount * 3;
        purchaseAmount = _purchaseAmount;

        // require that the required total is a sum of the parts
        require(totalEscrowAmount == (buyerRequiredEscrow + sellerRequiredEscrow));

        // event emission
        emit Initiated(msg.sender, buyer, seller, _purchaseAmount);
        emit Terms(purchaseAmount, totalEscrowAmount, buyerRequiredEscrow, sellerRequiredEscrow);
    }

    function() external payable {
        address _funder = msg.sender;
        uint _amount = msg.value;
        fundContract(_funder, _amount);
        emit FundsReceived(_funder, _amount);
    }

    function fundContract(address _funder, uint _amount) public {
        if (_funder == buyer && buyerEscrowedFunds <= buyerRequiredEscrow) {
          uint buyerCurrentFunds = buyerEscrowedFunds;
          buyerEscrowedFunds = buyerCurrentFunds + _amount;
        }
        if (_funder == seller && sellerEscrowedFunds <= sellerRequiredEscrow) {
          uint sellerCurrentFunds = sellerEscrowedFunds;
          sellerEscrowedFunds = sellerCurrentFunds + _amount;
        }
        checkFundingStatus();
    }

    function checkFundingStatus() internal {
        if (buyerEscrowedFunds >= buyerRequiredEscrow) {
          buyerHasFunded = true;
        }
        if (sellerEscrowedFunds >= sellerRequiredEscrow) {
          sellerHasFunded = true;
        }
        if (sellerHasFunded == true && buyerHasFunded == true) {
          escrowFullyFunded = true;
        } else {
          escrowFullyFunded = false;
        }
    }

    function buyerFinalize() public {
        require(msg.sender == buyer);
        buyerfinalized = true;
        emit PartyFinalized(msg.sender);

        if (sellerfinalized == true) {
          return finalizeTrade();
        }
    }

    function sellerFinalize() public {
        require(msg.sender == seller);
        sellerfinalized = true;
        emit PartyFinalized(msg.sender);

        if (buyerfinalized == true) {
          finalizeTrade();
        }
    }

    function finalizeTrade() internal {
        require(sellerHasFunded == true);
        require(buyerHasFunded == true);
        require(buyerfinalized == true);
        require(sellerfinalized == true);

        require(finalized == false);

        // Seller receives the purchaseAmount and their escrowed funds
        uint256 forSeller = purchaseAmount + sellerEscrowedFunds;

        // Buyer receives their escrowed funds minus the purchaseAmount
        uint256 forBuyer = buyerEscrowedFunds - purchaseAmount;

        emit EscrowComplete(finalized, forBuyer, forSeller);

        seller.transfer(forSeller);
        buyer.transfer(forBuyer);
        finalized = true;

        // destroy escrow contract to preserve anonymity and reduce chainwaste
        selfdestruct(seller);
    }

//    function withdrawalAllowed(address payee) public view virtual returns (bool);
//
//    function withdraw(address payable payee) public virtual override {
//        require(withdrawalAllowed(payee), "ConditionalEscrow: payee is not allowed to withdraw");
//        super.withdraw(payee);
//    }
}