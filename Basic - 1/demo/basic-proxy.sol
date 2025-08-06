// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Logic{
    address public logic;
    uint256 public count;

    function inc() external {
        count+=1;
    }

}

contract Proxy{
    address public logic;
    uint256 public count;
    constructor(address _logic){
        logic=_logic;
    }
    
    fallback() external  { 
      (bool ok,bytes memory res)=  logic.delegatecall(msg.data);
      require(ok,"delegatecall failed");
    }
}