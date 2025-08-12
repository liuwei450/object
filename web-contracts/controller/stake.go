package controller

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strings"
	"web-contracts/config"
	"web-contracts/schema"
)

func HandleStake(c *gin.Context) {
	var p schema.Stake
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := ethclient.Dial(config.RPC_URL)
	//连接失败
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//将十六进制私钥字符串转换为 ECDSA 私钥对象
	privateKey, err := crypto.HexToECDSA(config.PRIVATE_KEY_HEX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//从私钥派生出公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	//获取地址
	fromAdress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 获取当前账户的交易 nonce  账户已发交易数量。
	nonce, _ := client.PendingNonceAt(context.Background(), fromAdress)
	//节点建议的 gas 价格
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	//获取链 ID
	chainId, _ := client.NetworkID(context.Background())
	//创建带私钥的交易签名对象
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce)) //发送 ETH 的数量
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	//abi.JSON 解析合约 ABI 字符串为 ABI 对象
	parsedAbi, _ := abi.JSON(strings.NewReader(config.ABI_JSON))
	contract := common.HexToAddress(config.CONTRACT_ADRESS)
	inputAmount := big.NewInt(p.Amount)
	tx, err := bind.NewBoundContract(contract, parsedAbi, client, client, client).Transact(auth, "stake", inputAmount, p.Period)
	if err != nil {
		c.JSON(500, gin.H{"error": "Stake transaction failed", "detail": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Stake successful", "txHash": tx.Hash().Hex()})
}
