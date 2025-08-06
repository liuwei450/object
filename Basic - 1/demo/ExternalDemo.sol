// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

contract ExternalDemo{
    address public caller;

    function first() public {
        this.second();
    }
    function second()external  {
        caller=msg.sender;
    }
}

//0x5B38Da6a701c568545dCfcB03FcB875f56beddC4
//0x5B38Da6a701c568545dCfcB03FcB875f56beddC4