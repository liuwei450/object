// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;



contract Caller01{
    address callerAdress;
    uint public xx;
    constructor(address _callerAdress){
        callerAdress=_callerAdress;
    }

    function setCallerX() public {
        bytes memory cd=abi.encodeWithSignature("setX(uint256)", 4);
        (bool success,bytes memory rst ) = callerAdress.call(cd);
        if(!success){
            revert("call failed");
        }
        //解码
       (uint x)=abi.decode(rst, (uint));
        xx=x;
        
    }
}