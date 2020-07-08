pragma solidity ^0.6.0;

import "./installed_contracts/zeppelin/contracts/access/Ownable.sol";
import "./EscrowContractInterface.sol";

contract AuthenticityOracle is Ownable {
    uint private randNonce = 0;
    uint private modulus = 1000;
    mapping(uint256=>bool) pendingRequests;
    event GetAuthenticityCheckEvent(address callerAddress, uint id);
    event SetAuthenticityCheckEvent(bool authenticityCheck, address callerAddress);
    function GetAuthenticityCheck() public returns (uint256) {
        randNonce++;
        uint id = uint(keccak256(abi.encodePacked(now, msg.sender, randNonce))) % modulus;
        pendingRequests[id] = true;
        emit GetAuthenticityCheckEvent(msg.sender, id);
        return id;
    }
    function setAuthenticityCheck(bool _authenticityCheck, address _callerAddress,   uint256 _id) public onlyOwner {
        require(pendingRequests[_id], "Request not in pending list.");
        delete pendingRequests[_id];
        EscrowContractInterface escrowContractInstance;
        escrowContractInstance = EscrowContractInterface(_callerAddress);
        escrowContractInstance.callback(_authenticityCheck, _id);
        emit SetAuthenticityCheckEvent(_authenticityCheck, _callerAddress);
  }
}