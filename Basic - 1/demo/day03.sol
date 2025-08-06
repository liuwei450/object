// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract  SignatureVerifier {
    function getMessageHash(string memory message) public pure returns (bytes32) {
        // 把字符串编码为 bytes 后进行 keccak256 哈希
        return keccak256(abi.encodePacked(message));
    }
    function getEthSignedMessageHash(bytes32 messageHash) public pure returns (bytes32) {
         /*
         Ethereum 签名是对带前缀的消息签名的：
         "\x19Ethereum Signed Message:\n32" + messageHash
         这是为了防止签名被重用为交易、授权等内容。
        */
         return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
    }
    // @notice 验证：传入的签名是否由 signer 对 message 签发
    function verify(
        address signer, // 预期签名者地址
        string memory message,// 原始消息
         bytes memory signature  // 用户签名（65字节：r[32] + s[32] + v[1]）
         ) public pure returns (bool) {
            // 1. 对消息进行第一次哈希
        bytes32 messageHash = getMessageHash(message);
        // 2. 加上 Ethereum 前缀，再哈希一次（确保与 signMessage() 结果一致）
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);
        // 3. 恢复出签名者地址，看是否等于期望地址
        return recoverSigner(ethSignedMessageHash, signature) == signer;
    }
    // @notice 根据签名恢复出地址
    function recoverSigner(bytes32 _ethSignedMessageHash, bytes memory _signature) public pure returns (address) {
        require(_signature.length == 65, "Invalid signature length");

        bytes32 r;
        bytes32 s;
        uint8 v;

        // 使用内联汇编解析 r, s, v
        assembly {
            r := mload(add(_signature, 32)) // 取前32字节为 r
            s := mload(add(_signature, 64))  // 再取32字节为 s
            v := byte(0, mload(add(_signature, 96))) // 最后1字节为 v（恢复 ID）
        }
        // v 有可能是 0/1，也可能是 27/28，我们统一转为 27/28（ecrecover 要求）
        if (v < 27) {
            v += 27;
        }
        
        // 使用 ecrecover 内置函数：恢复出签名地址
        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

}