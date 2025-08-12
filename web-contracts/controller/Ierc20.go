package controller

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"web-contracts/config"
	erc20 "web-contracts/contract"
	"web-contracts/schema"
)

// erc授权
func HandleApprove(c *gin.Context) {
	var p schema.ApproveRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 连接以太坊客户端（RPC节点）
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "连接 RPC 失败: " + err.Error()})
		return
	}
	// 5. 解析参数
	privateKey, err := crypto.HexToECDSA(config.PRIVATE_KEY_HEX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析私钥失败: " + err.Error()})
		return
	}
	chainID := big.NewInt(config.CHAIN_ID)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Value = big.NewInt(0)
	auth.GasLimit = 100000
	auth.GasPrice, _ = client.SuggestGasPrice(context.Background())

	amount := new(big.Int)
	if _, ok := amount.SetString(p.Value, 10); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "授权数量格式错误"})
		return
	}
	//获取erc20
	erc20Instance, err := erc20.NewErc20(common.HexToAddress(config.ERC20_ADDRESS), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERC20 实例化失败: " + err.Error()})
		return
	}
	//授权
	tx, err := erc20Instance.Approve(auth, common.HexToAddress(p.Spender), amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Approve 失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"txHash":  tx.Hash().Hex(),
		"message": "授权交易已发送，请等待链上确认",
	})

}

// 查询授权额度
func HandleAllowance(c *gin.Context) {
	var p schema.AllowanceRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "连接 RPC 失败: " + err.Error()})
		return
	}
	erc20Instance, err := erc20.NewErc20(common.HexToAddress(config.ERC20_ADDRESS), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERC20 实例化失败: " + err.Error()})
		return
	}
	allowance, err := erc20Instance.Allowance(nil, common.HexToAddress(p.Owner), common.HexToAddress(p.Spender))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询授权额度失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"owner":     p.Owner,
		"spender":   p.Spender,
		"allowance": allowance.String(), // 返回字符串，避免大数精度丢失
	})
}

// 查询余额
func HandleBalance(c *gin.Context) {
	var p schema.BalanceRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "连接 RPC 失败: " + err.Error()})
		return
	}
	erc20Instance, err := erc20.NewErc20(common.HexToAddress(config.ERC20_ADDRESS), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERC20 实例化失败: " + err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERC20 实例化失败: " + err.Error()})
		return
	}
	balance, err := erc20Instance.BalanceOf(nil, common.HexToAddress(p.BalanceAdress))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询授权额度失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"BalanceAdress": p.BalanceAdress,
		"allowance":     balance.String(), // 返回字符串，避免大数精度丢失
	})

}
