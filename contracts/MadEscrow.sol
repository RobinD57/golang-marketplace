//pragma solidity ^0.4.23;
//
///*
//This escrow contract uses a similar mechanism as the MAD escrow protocol:
//https://pbs.twimg.com/media/DcsJhB7VQAA-kHx.jpg?name=orig and is inspired by a solidity MAD escrow
//for ETH <=> VEO trades: https://github.com/RyanHendricks/mad-escrow-solidity/blob/master/contracts/VeoEscrow.sol
//The difference is that the finalization does not depend
//*/
//
//contract VeoEscrow {
//
//    // Escrow parameters set at deployment
//    address public buyer; //BUYER
//    address public seller; //SELLER
//    uint public purchaseAmount = 0; //Total paid by BUYER
//
//    // escrow state
//    bool public initiated = false;
//    bool public finalized = false;
//
//    // @dev Buyer required escrow funding (2x purchaseAmount)
//    uint public buyerRequiredEscrow = 0;
//    // @dev Seller required escrow funding (1x purchaseAmount)
//    uint public sellerRequiredEscrow = 0;
//    // @dev TOTAL required escrow funding (3x purchaseAmount)
//    uint public totalEscrowAmount = 0;
//
//    // @dev Buyer escrow funding received
//    uint public buyerEscrowedFunds = 0;
//    // @dev Seller escrow funding received
//    uint public sellerEscrowedFunds = 0;
//
//    // @notice This will be true when the buyer has sent 2x the purchaseAmount
//    // @dev True when buyerEscrowedFunds >= buyerRequiredEscrow
//    bool public buyerHasFunded = false;
//    // @notice Has Seller sent required amount of ETH to escrow
//    // @dev True when sellerEscrowedFunds >= sellerRequiredEscrow
//    bool public sellerHasFunded = false;
//
//    // @notice Has enough ETH been sent to escrow to meet escrow requirement
//    // @dev True when both buyerHasFunded = true AND sellerHasFunded = true
//    bool public escrowFullyFunded = false;
//
//
//    // Once FullyFunded, both parties must finalize the trade/sale
//    bool public buyerfinalized = false;
//    bool public sellerfinalized = false;
//
//    event Terms(uint purchaseAmount, uint totalEscrowAmount, uint BuyerRequiredFunding, uint SellerRequiredFunding);
//    event Initiated(address initiator, address buyer, address seller, uint tradeFundsforXfer);
//    event FundsReceived(address sender, uint amount);
//    event PartyFinalized(address party);
//    event EscrowComplete(bool completed, uint buyerReceives, uint sellerReceives);
//
//
//    /** Constructor
//     * @param _buyer address of the buyer
//     * @param _seller address of the seller
//     * @param _purchaseAmount uint amount of ETH being paid to seller
//     */
//    function VeoEscrow(
//        address _buyer,
//        address _seller,
//        uint _purchaseAmount
//    ) public {
//        // require _seller and _buyer to be non-zero and unique.
//        require(_seller != address(0));
//        require(_buyer != address(0));
//        require(_buyer != _seller);
//
//        // require the contract creator to be either _buyer or _seller
//        require(msg.sender == _seller || msg.sender == _buyer);
//
//        // require escrow amount to be non-zero
//        require(_purchaseAmount > 0);
//
//        // ensure variable seet to false
//        escrowFullyFunded = false;
//
//        // assign trade parties
//        buyer = _buyer; //Buyer
//        seller = _seller; //Seller
//
//        // set state to initiated
//        initiated = true;
//
//        // set required funding amounts
//        buyerRequiredEscrow = _purchaseAmount * 2;
//        sellerRequiredEscrow = _purchaseAmount;
//        totalEscrowAmount = _purchaseAmount * 3;
//        purchaseAmount = _purchaseAmount;
//
//        // assert that the required total is a sum of the parts
//        assert(totalEscrowAmount == (buyerRequiredEscrow + sellerRequiredEscrow));
//
//        // event emission
//        emit Initiated(msg.sender, buyer, seller, _purchaseAmount);
//        emit Terms(purchaseAmount, totalEscrowAmount, buyerRequiredEscrow, sellerRequiredEscrow);
//    }
//
//    /**
//     * @dev fallback function to receive funds sent to this escrow contract
//     * @dev this function then calls the internal fundContract() function
//     */
//    function() external payable {
//        address _funder = msg.sender;
//        uint _amount = msg.value;
//        fundContract(_funder, _amount);
//        emit FundsReceived(_funder, _amount);
//    }
//
//    /** @dev called internally to add received funds to the respective party balance
//     *  @param _funder address that sent ETH to fund escrow
//     *  @param _amount uint of ETH that was send by the _funder
//     *  @dev if the _funder's respective balance is less than required the _amount is
//     *       added to their EscrowedFunds balance
//     *  @dev finally, the checkFundingStatus() function is called internally
//     */
//    function fundContract(address _funder, uint _amount) public {
//        if (_funder == buyer && buyerEscrowedFunds <= buyerRequiredEscrow) {
//            uint buyerCurrentFunds = buyerEscrowedFunds;
//            buyerEscrowedFunds = buyerCurrentFunds + _amount;
//        }
//        if (_funder == seller && sellerEscrowedFunds <= sellerRequiredEscrow) {
//            uint sellerCurrentFunds = sellerEscrowedFunds;
//            sellerEscrowedFunds = sellerCurrentFunds + _amount;
//        }
//        checkFundingStatus();
//    }
//    /**
//     * @notice internal function to check funding status
//     * @dev set respective party's funding status to true if
//     * amount of funds received from said party is >= required funds
//     * @dev If both buyer and seller have sent enough ETH
//     *      as per required amounts than set escrowFullyFunded to true
//     */
//    function checkFundingStatus() internal {
//        if (buyerEscrowedFunds >= buyerRequiredEscrow) {
//            buyerHasFunded = true;
//        }
//        if (sellerEscrowedFunds >= sellerRequiredEscrow) {
//            sellerHasFunded = true;
//        }
//        if (sellerHasFunded == true && buyerHasFunded == true) {
//            escrowFullyFunded = true;
//        } else {
//            escrowFullyFunded = false;
//        }
//    }
//
//    /**
//     * @notice Once fully funded, both parties must finalize the transaction
//     * @dev Once both finalizations are received the funds are released accordingly
//     * @dev The logic works such that if one party finalizes the contract checks
//     *      to see if the other party has already finalized.
//     */
//    function buyerFinalize() public {
//        require(msg.sender == buyer);
//        buyerfinalized = true;
//        emit PartyFinalized(msg.sender);
//
//        if (sellerfinalized == true) {
//            return finalizeTrade();
//        }
//    }
//
//    function sellerFinalize() public {
//        require(msg.sender == seller);
//        sellerfinalized = true;
//        emit PartyFinalized(msg.sender);
//
//        if (buyerfinalized == true) {
//            finalizeTrade();
//        }
//    }
//
//    /**
//     * @dev finalizeTrade called internally to distribute the escrow funds
//     */
//    function finalizeTrade() internal {
//        // require both parties to have funded the escrow as required
//        require(sellerHasFunded == true);
//        require(buyerHasFunded == true);
//
//        // require both parties to have finalized the trade
//        require(buyerfinalized == true);
//        require(sellerfinalized == true);
//
//        // require that the trade has not yet been finalized
//        require(finalized == false);
//
//        // Seller receives the purchaseAmount and their escrowed funds
//        uint256 forSeller = purchaseAmount + sellerEscrowedFunds;
//
//        // Buyer receives their escrowed funds minus the purchaseAmount
//        uint256 forBuyer = buyerEscrowedFunds - purchaseAmount;
//
//        // event emission
//        emit EscrowComplete(finalized, forBuyer, forSeller);
//
//        // transfer the funds
//        seller.transfer(forSeller);
//        buyer.transfer(forBuyer);
//
//        // finalize the trade
//        finalized = true;
//
//        // destroy escrow contract to preserve anonymity and reduce chainwaste
//        selfdestruct(seller);
//
//
//    }
//
//}