pragma solidity ^0.6.0;

interface EscrowContractInterface {
  function callback(bool _authenticityCheck, uint256 id) external;
}
