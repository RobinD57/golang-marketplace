pragma solidity ^0.6.0;

import "./AuthenticityOracleInterface.sol";
import "./installed_contracts/zeppelin/contracts/payment/Escrow.sol";

contract MadEscrow is Escrow {
    string public name = "Solidity MAD escrow contract";
    AuthenticityOracleInterface private oracleInstance;
    address private oracleAddress;
    bool private authenticityCheck;
    mapping(uint256=>bool) myRequests;
    event newOracleAddressEvent(address oracleAddress);
    event ReceivedNewRequestIdEvent(uint256 id);
    event AuthenticityCheckCompletedEvent(bool authenticityCheck, uint256 id);

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

//    function withdrawalAllowed(address payee) public view virtual returns (bool);
//
//    function withdraw(address payable payee) public virtual override {
//        require(withdrawalAllowed(payee), "ConditionalEscrow: payee is not allowed to withdraw");
//        super.withdraw(payee);
//    }
}