pragma solidity ^0.6.0;

interface AuthenticityOracleInterface {
  function authenticityCheck() external returns (uint256);
}
