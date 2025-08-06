// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Callee{
    uint public x;
    function setX(uint _x)public returns (uint){
        x=_x;
        return x;
    }

}