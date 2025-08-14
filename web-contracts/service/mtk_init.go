package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"web-contracts/models"

	"web-contracts/config"
	"web-contracts/contract"
)

func InitClient() error {
	var err error
	// 初始化客户端
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		return fmt.Errorf("连接节点失败（%s）: %w", config.RPC_URL, err)

	}

	// 加载私钥
	//privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(config.PRIVATE_KEY_HEX)
	if err != nil {
		return fmt.Errorf("解析私钥失败: %w", err)

	}
	// 生成发送者地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("公钥类型转换失败（不是 *ecdsa.PublicKey）")

	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("当前钱包地址:", fromAddress.Hex())
	// 获取链ID
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %w", err)

	}
	// 5. 创建交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return fmt.Errorf("创建交易签名器失败: %w", err)

	}

	// 解析合约地址
	contractAddress := common.HexToAddress(config.CONTRACT_ADRESS)
	//绑定合约
	stakingContract, err := contract.NewStakeContract(contractAddress, client)
	if err != nil {
		return fmt.Errorf("绑定合约失败（地址: %s）: %w", contractAddress.Hex(), err)

	}
	// 注册到全局配置
	models.NewInitContract(models.IiitContract{
		StakingContract: stakingContract,
		Auth:            auth,
		FromAddress:     fromAddress.String(),
		Client:          client,
	})
	fmt.Println("初始化完成，客户端和合约已绑定")
	return err
}
