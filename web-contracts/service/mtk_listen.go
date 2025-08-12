package service

import (
	"context"
	"fmt"
	"web-contracts/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"web-contracts/contract"
	"web-contracts/models"
)

func ListenStakedEvents() {
	client, err := ethclient.Dial(config.RPC_URL)
	if err != nil {
		log.Fatal("连接 RPC 失败:", err)
	}

	contractAddr := common.HexToAddress(config.CONTRACT_ADRESS)
	instance, err := contract.NewStakeContract(contractAddr, client)
	if err != nil {
		log.Fatal("加载合约失败:", err)
	}

	query := &bind.WatchOpts{
		Context: context.Background(),
		Start:   nil, // 实时监听
	}

	ch := make(chan *contract.StakeContractStaked)
	sub, err := instance.WatchStaked(query, ch, nil)
	if err != nil {
		log.Fatal("订阅事件失败:", err)
	}
	fmt.Println("正在监听 Staked 事件...")

	for {
		select {
		case err := <-sub.Err():
			log.Println("监听出错:", err)
		case e := <-ch:
			fmt.Println("监听到 Staked 事件:", e.User.Hex(), e.StakeId, e.Amount, e.Period, e.Timestamp)

			models.DB.Create(&models.StakeRecord{
				User:      e.User.Hex(),
				StakeID:   e.StakeId.String(),
				Amount:    e.Amount.String(),
				Period:    e.Period,
				Timestamp: e.Timestamp.Uint64(),
			})
		}
	}
}
