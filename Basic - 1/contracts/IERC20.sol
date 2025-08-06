// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface IERC20 {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8); 
    
}