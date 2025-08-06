// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

contract BalanceManager {
    mapping(address=>uint256) public balanceOf;

    constructor(uint totalBalance){
        balanceOf[msg.sender]=totalBalance;
    }

    function transfer( address _to, uint256 _amount) public {
        address _from=msg.sender;
        uint256 fb=balanceOf[_from];
        uint256 tb=balanceOf[_to];
        require(_amount<=fb, "Insufficient balance");
        fb-=_amount;
        tb+=_amount;
        balanceOf[_from]=fb;
        balanceOf[_to]=tb;
    }
    
}