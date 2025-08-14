package models

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"web-contracts/contract"
)

var (
	contractInfo IiitContract // 全局合约配置实例
	dataMutex    sync.RWMutex // 读写锁（保护 contractInfo 的并发访问）
)

type IiitContract struct {
	StakingContract *contract.StakeContract // 智能合约实例（自动生成的 Go 绑定）
	Auth            *bind.TransactOpts      // 交易签名器（包含私钥、Gas 配置等）
	FromAddress     string                  // 发送者钱包地址（字符串格式）
	Client          *ethclient.Client       // 区块链客户端（连接节点的句柄）
}

// 获取合约配置
func GetInitContract() IiitContract {
	dataMutex.RLock()         // 获取读锁
	defer dataMutex.RUnlock() // 释放读锁
	return contractInfo
}

func NewInitContract(c IiitContract) {
	dataMutex.Lock()         // 获取写锁（独占）
	defer dataMutex.Unlock() // 释放写锁

	contractInfo = IiitContract{
		StakingContract: c.StakingContract,
		Auth:            c.Auth,
		FromAddress:     c.FromAddress,
		Client:          c.Client,
	}
}
