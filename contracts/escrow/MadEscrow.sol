pragma solidity ^0.6.0;

import "../installed_contracts/zeppelin/contracts/utils/Address.sol";

abstract contract ConditionalEscrow is Escrow {
    AuthencityOracleInterface private oracleInstance;
    address private oracleAddress;
    bool private authenticityCheck;
    mapping(uint256=>bool) myRequests;
    event AuthencityCheckCompletedEvent(bool authenticityCheck, uint256 id);
    event ReceivedNewRequestIdEvent(uint256 id);

    function setOracleAddress(address _oracleAddress) public onlyOwner {
            oracleAddress = _oracleAddress;
    }

    function updateAuthenticityCheck() public {
      uint256 id = oracleInstance.authentictyCheck();
      myRequests[id] = true;
      emit ReceivedNewRequestIdEvent(id);
    }

    function callback(bool _authenticityCheck, uint256 _id) public onlyOracle {
      require(myRequests[_id], "Request not in pending list.");
      authenticityCheck = _authenticityCheck;
      delete myRequests[_id];
      emit AuthencityCheckCompletedEvent(_authenticityCheck, _id);
    }

    modifier onlyOracle() {
     require(msg.sender == oracleAddress, "You are not authorized to call this function.");
      _;
    }

    function withdrawalAllowed(address payee) public view virtual returns (bool);

    function withdraw(address payable payee) public virtual override {
        require(withdrawalAllowed(payee), "ConditionalEscrow: payee is not allowed to withdraw");
        super.withdraw(payee);
    }
}