// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Callee{
    uint public x;
    address public caller;
    address public eoaAddress;
    function setX(uint _x)public   {
        caller=msg.sender;
        eoaAddress=tx.origin;
         x=_x;
       
    }

}

contract Caller{
    address public caller;
    address calleeAdress;
    address public eoaAddress;
    
    constructor(address _calleeAdress){
         calleeAdress=_calleeAdress;
    }

    function setCalleeX(uint _x)public {
      caller=msg.sender;
      eoaAddress=tx.origin;
      Callee call=Callee(calleeAdress);
      call.setX(_x);
    }

}