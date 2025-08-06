// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract HelloComrades{
    ///@notice 用来取当前合约余额
    ///@dev 一个当前合约的辅助函数
    ///@return 当前合约余额
function getBalance() external  view returns  (uint256){
    return address(this).balance;

 }
}