// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyToken {
    string public name = "MyToken";
    string public symbol = "MTK";
    uint8 public decimals = 18;
    uint256 public totalSupply;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    //NFT部门
    uint256 public nextTokenId = 1;
    mapping(uint256 => address) private _nftOwners;
    mapping(address => uint256) public ownedNFT;
    //事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event NFTMinted(address indexed to, uint256 tokenId);
    event NFTTransferred(address indexed from, address indexed to, uint256 tokenId);

    constructor(uint256 initialSupply) {
         require(initialSupply > 0, "Initial supply must be greater than zero");
        totalSupply = initialSupply * 10 ** uint256(decimals);
        balanceOf[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }
     // ----------------
    // ERC-20 Functions
    // ----------------
    function transfer(address to, uint256 value) public returns (bool success) {
        require(balanceOf[msg.sender] >= value, "ERC20: Insufficient balance");
        balanceOf[msg.sender] -= value;
        balanceOf[to] += value;
        emit Transfer(msg.sender, to, value);
        return true;
    }

    function approve(address spender, uint256 value) public returns (bool success) {
        allowance[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool success) {
        require(balanceOf[from] >= value, "Insufficient balance");
        require(allowance[from][msg.sender] >= value, "Allowance exceeded");

        balanceOf[from] -= value;
        balanceOf[to] += value;
        allowance[from][msg.sender] -= value;
        emit Transfer(from, to, value);
        return true;
    }
    // @notice 铸造一个 NFT，每个地址只能拥有一个
     function mintNFT() public returns (uint256) {
        require(ownedNFT[msg.sender] == 0, "Already owns NFT");

        uint256 tokenId = nextTokenId++;
        _nftOwners[tokenId] = msg.sender;
        ownedNFT[msg.sender] = tokenId;

        emit NFTMinted(msg.sender, tokenId);
        return tokenId;
    }
    // @notice 转移 NFT 所有权
    function transferNFT(address to, uint256 tokenId) public {
        require(_nftOwners[tokenId] == msg.sender, "Not owner of NFT");
        require(to != address(0), "Invalid address");

        _nftOwners[tokenId] = to;
        ownedNFT[msg.sender] = 0;
        ownedNFT[to] = tokenId;

        emit NFTTransferred(msg.sender, to, tokenId);
    }
    /// @notice 获取 tokenId 的所有者
    function ownerOf(uint256 tokenId) public view returns (address) {
        address owner = _nftOwners[tokenId];
        require(owner != address(0), "Token does not exist");
        return owner;
    }
}