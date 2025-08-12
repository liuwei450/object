package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mtk-contracts/config"
	"mtk-contracts/contract"
	"mtk-contracts/models"
)

// 全局变量
var (
	db              *gorm.DB
	client          *ethclient.Client
	stakingContract *contract.StakeContract
)

func InitClient() {
	// 初始化客户端
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the BSC network: %v", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(config.PRIVATE_KEY_HEX)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("当前钱包地址:", fromAddress.Hex())
	// 获取链ID
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}
	// 5. 创建交易授权器（备用）
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("创建交易授权器失败: %v", err)
	}
	_ = auth
	// 6. 绑定合约
	contractAddress := common.HexToAddress(config.CONTRACT_ADRESS)
	stakingContract, err = contract.NewStakeContract(contractAddress, client)
	if err != nil {
		log.Fatalf("绑定合约失败: %v", err)
	}

	fmt.Println("初始化完成，客户端和合约已绑定")
}

// WatchStakeEvents 监听 Staked 事件并保存到数据库
func WatchStakeEvents() {
	fmt.Println("开始监听 Staked 事件...")

	watchOpts := &bind.WatchOpts{Context: context.Background()}
	logsCh := make(chan *contract.StakeContractStaked)

	sub, err := stakingContract.WatchStaked(watchOpts, logsCh, nil) // nil 表示不筛选 indexed 参数
	if err != nil {
		log.Fatalf("监听 Stake 事件失败: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("事件订阅错误: %v", err)
			return
		case evt := <-logsCh:
			log.Infof("新 Stake 事件: user=%s stakeId=%d amount=%s period=%d",
				evt.User.Hex(), evt.StakeId.Uint64(), evt.Amount.String(), evt.Period)

			record := models.StakeEvent{
				User:      evt.User.Hex(),
				StakeID:   evt.StakeId.Uint64(),
				Amount:    evt.Amount.String(),
				Period:    uint64(evt.Period),
				Timestamp: evt.Timestamp.Uint64(),
			}
			db.Create(&record)
		}
	}
}
