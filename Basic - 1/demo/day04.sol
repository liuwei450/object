// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

contract ArryType{
    uint8[3] data;
    uint8[] ddata;
    function testStaticArry() public   returns (uint8[3] memory){
        data[0]=22;
        return data;

    }
    function testReadDynamicsStorageArry()public view returns (uint8[] memory){
        return ddata;
    }
    function testWriteDynamicsStorageArry()public {
        ddata.push(12);
        ddata.pop();
        ddata.push(90);
    }
    function testMemoryDynaicArra(uint8 size)public pure  returns (uint8[] memory){
        uint8[]memory mdata=new uint8[](size);
        return mdata;
    }
}



