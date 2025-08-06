// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Hello{
    string public message="Hello World";
    function fn1()public view returns  (string memory){
        return message;
    }
     function fn2()public pure  returns  (string memory){
        return "Hello World";
    }
    function fn3()public pure  returns  (string memory){
        return fn2();
    }
}