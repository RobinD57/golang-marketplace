pragma solidity ^0.6.0;

contract EscrowContractInterface {
  function callback(bool _authenticityCheck, uint256 id) public;
}
