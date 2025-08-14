package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"sync"
	"time"
	"web-contracts/contract"
	"web-contracts/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// ------------------------------
// 1. 事件类型常量（与合约事件对应）
// ------------------------------
const (
	EventTypeStaked    = "Staked"    // 质押事件
	EventTypeWithdrawn = "Withdrawn" // 提现事件
)

// ------------------------------
// 2. 通用事件处理接口
// ------------------------------
type EventProcessor interface {
	Validate(event interface{}) error               // 校验事件数据合法性
	Convert(event interface{}) (interface{}, error) // 转换为数据库模型
	Save(tx *gorm.DB, model interface{}) error      // 事务入库（支持去重）
}

// ------------------------------
// 3. 事件分发器（注册与路由）
// ------------------------------
var (
	eventProcessors = sync.Map{} // 事件类型 → 处理器实例（线程安全）
)

// RegisterProcessor 注册事件处理器
func RegisterProcessor(eventType string, processor EventProcessor) {
	if _, exists := eventProcessors.Load(eventType); exists {
		fmt.Println("事件处理器 已存在，将被覆盖", eventType)
	}
	eventProcessors.Store(eventType, processor)
}

// ------------------------------
// 4. 全局监听服务（单例 + 多事件订阅）
// ------------------------------
var (
	listenOnce   sync.Once
	isListening  bool
	stakedSub    ethereum.Subscription   // Staked 事件订阅实例
	withdrawnSub ethereum.Subscription   // Withdrawn 事件订阅实例
	client       *ethclient.Client       // 区块链客户端（复用）
	contractInst *contract.StakeContract // 合约实例（StakeContract 类型）
)

// StartEventMonitor 启动事件监听服务（单例）
func StartEventMonitor() {
	listenOnce.Do(func() {
		if isListening {
			fmt.Println("事件监听服务已在运行中，无需重复启动")
			return
		}
		isListening = true
		defer func() { isListening = false }()

		// 1. 初始化全局资源（客户端、合约实例）
		if err := initGlobalResources(); err != nil {
			fmt.Println("初始化全局资源失败:", err)
			return
		}

		// 2. 注册事件处理器
		RegisterProcessor(EventTypeStaked, &StakedProcessor{})
		RegisterProcessor(EventTypeWithdrawn, &WithdrawnProcessor{})

		// 3. 启动事件监听协程（质押 + 提现）
		wg := sync.WaitGroup{}
		wg.Add(2)
		go startStakedListener(&wg)    // 质押事件监听
		go startWithdrawnListener(&wg) // 提现事件监听

		wg.Wait() // 等待所有监听协程退出（理论上阻塞）
		fmt.Println("所有事件监听协程已退出，服务终止")
	})
}

// initGlobalResources 初始化客户端和合约实例
func initGlobalResources() error {
	// 从全局配置获取已初始化的资源（确保先调用 models.InitClient()）
	globalContract := models.GetInitContract()
	if globalContract.StakingContract == nil || globalContract.Client == nil {
		return errors.New("全局合约实例未初始化，请先调用 models.InitClient")
	}
	client = globalContract.Client                // 区块链客户端
	contractInst = globalContract.StakingContract // 合约实例（*contract.StakeContract）
	return nil
}

// ------------------------------
// 5. 事件监听实现（质押 + 提现）
// ------------------------------

// startStakedListener 启动质押事件监听
func startStakedListener(wg *sync.WaitGroup) {
	defer wg.Done()
	eventCh := make(chan *contract.StakeContractStaked, 100) // 带缓冲通道

	for {
		// 订阅事件（失败后自动重试）
		sub, err := subscribeStakedEvent(eventCh)
		if err != nil {
			fmt.Println("质押事件订阅失败，将重试:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		stakedSub = sub
		fmt.Println("质押事件监听已启动，等待事件触发...")

		// 处理事件和错误
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("质押事件监听错误:", err)
				sub.Unsubscribe()
				break // 退出内层循环，重试订阅
			case event := <-eventCh:
				processEvent(EventTypeStaked, event) // 分发事件处理
			}
		}
	}
}

// startWithdrawnListener 启动提现事件监听
func startWithdrawnListener(wg *sync.WaitGroup) {
	defer wg.Done()
	eventCh := make(chan *contract.StakeContractWithdrawn, 100) // 带缓冲通道

	for {
		// 订阅事件（失败后自动重试）
		sub, err := subscribeWithdrawnEvent(eventCh)
		if err != nil {
			fmt.Println("提现事件订阅失败，将重试:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		withdrawnSub = sub
		fmt.Println("提现事件监听已启动，等待事件触发...")

		// 处理事件和错误
		for {
			select {
			case err := <-sub.Err():
				fmt.Println("提现事件监听错误:", err)
				sub.Unsubscribe()
				break // 退出内层循环，重试订阅
			case event := <-eventCh:
				processEvent(EventTypeWithdrawn, event) // 分发事件处理
			}
		}
	}
}

// ------------------------------
// 6. 事件订阅底层实现
// ------------------------------

// subscribeStakedEvent 订阅 Staked 事件
func subscribeStakedEvent(eventCh chan *contract.StakeContractStaked) (ethereum.Subscription, error) {
	opts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   nil, // nil 表示从最新区块开始监听
	}
	// 调用合约绑定代码的 WatchStaked 方法（自动生成）
	return contractInst.WatchStaked(opts, eventCh, nil) // nil 表示监听所有用户
}

// subscribeWithdrawnEvent 订阅 Withdrawn 事件
func subscribeWithdrawnEvent(eventCh chan *contract.StakeContractWithdrawn) (ethereum.Subscription, error) {
	opts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   nil,
	}
	// 调用合约绑定代码的 WatchWithdrawn 方法（自动生成）
	return contractInst.WatchWithdrawn(opts, eventCh, nil) // nil 表示监听所有用户
}

// ------------------------------
// 7. 通用事件处理流程（分发 + 校验 + 入库）
// ------------------------------

// processEvent 分发事件到对应处理器
func processEvent(eventType string, event interface{}) {
	// 获取事件处理器
	processor, ok := eventProcessors.Load(eventType)
	if !ok {
		fmt.Println("未注册事件处理器", "event_type", eventType)
		return
	}

	// 类型断言：确保处理器实现 EventProcessor 接口
	ep, ok := processor.(EventProcessor)
	if !ok {
		fmt.Println("事件处理器类型错误", "event_type", eventType)
		return
	}

	// 1. 校验事件数据
	if err := ep.Validate(event); err != nil {
		fmt.Println("事件数据校验失败", "event_type", eventType, "error", err)
		return
	}

	// 2. 转换为数据库模型
	model, err := ep.Convert(event)
	if err != nil {
		fmt.Println("事件数据转换失败", "event_type", eventType, "error", err)
		return
	}

	// 3. 事务入库（确保原子性）
	if err := models.DB.Transaction(func(tx *gorm.DB) error {
		return ep.Save(tx, model)
	}); err != nil {
		fmt.Println("事件数据入库失败", "event_type", eventType, "error", err)
	} else {
		fmt.Println("事件处理成功", "event_type", eventType, "tx_hash", getTxHashFromModel(model))
	}
}

// getTxHashFromModel 从模型中提取 TxHash（日志用）
func getTxHashFromModel(model interface{}) string {
	switch m := model.(type) {
	case *models.StakeRecord:
		return m.TxHash
	case *models.WithdrawRecord:
		return m.TxHash
	default:
		return "unknown"
	}
}

// ------------------------------
// 8. 事件处理器实现（质押 + 提现）
// ------------------------------

// ------------------------------
// Staked 事件处理器
// ------------------------------
type StakedProcessor struct{}

// Validate 校验质押事件数据
func (p *StakedProcessor) Validate(event interface{}) error {
	e, ok := event.(*contract.StakeContractStaked)
	if !ok {
		return errors.New("StakedProcessor: 事件类型断言失败（应为 *contract.StakeContractStaked）")
	}
	// 基础校验
	if e.User == (common.Address{}) {
		return errors.New("用户地址为空（零地址）")
	}
	if e.Amount == nil || e.Amount.Sign() <= 0 {
		return errors.New("质押数量必须大于 0")
	}
	if e.StakeId == nil || e.StakeId.Sign() <= 0 {
		return errors.New("StakeId 无效（必须大于 0）")
	}
	if e.Timestamp == nil || e.Timestamp.Sign() <= 0 {
		return errors.New("事件时间戳无效")
	}
	if e.Raw.TxHash == (common.Hash{}) {
		return errors.New("交易哈希为空")
	}
	return nil
}

// Convert 转换为数据库模型（StakeRecord）
func (p *StakedProcessor) Convert(event interface{}) (interface{}, error) {
	e := event.(*contract.StakeContractStaked)
	return &models.StakeRecord{
		UserAddress: e.User.Hex(),         // 用户地址（0x前缀）
		StakeID:     e.StakeId.String(),   // StakeId（*big.Int 转字符串）
		Amount:      e.Amount.String(),    // 质押数量（wei 单位）
		Period:      e.Period,             // 质押周期
		Timestamp:   e.Timestamp.Uint64(), // 区块链时间戳（秒级）
		TxHash:      e.Raw.TxHash.Hex(),   // 交易哈希（唯一索引）
		EventType:   EventTypeStaked,
	}, nil
}

// Save 入库（带交易哈希去重）
func (p *StakedProcessor) Save(tx *gorm.DB, model interface{}) error {
	record := model.(*models.StakeRecord)
	// 检查交易哈希是否已存在（防重复）
	var existing models.StakeRecord
	if err := tx.Where("tx_hash = ?", record.TxHash).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("查询质押记录失败:", err)
			return err
		}
		// 新记录，执行插入
		return tx.Create(record).Error
	}
	fmt.Println("质押事件已存在（重复推送）", "tx_hash", record.TxHash)
	return nil
}

// ------------------------------
// Withdrawn 事件处理器
// ------------------------------
type WithdrawnProcessor struct{}

// Validate 校验提现事件数据
func (p *WithdrawnProcessor) Validate(event interface{}) error {
	e, ok := event.(*contract.StakeContractWithdrawn)
	if !ok {
		return errors.New("WithdrawnProcessor: 事件类型断言失败（应为 *contract.StakeContractWithdrawn）")
	}
	// 基础校验
	if e.User == (common.Address{}) {
		return errors.New("用户地址为空（零地址）")
	}
	if e.Principal == nil || e.Principal.Sign() <= 0 {
		return errors.New("提现本金必须大于 0")
	}
	if e.TotalAmount == nil || e.TotalAmount.Sign() <= e.Principal.Sign() {
		return errors.New("总金额（本金+利息）必须大于本金")
	}
	if e.StakeId == nil || e.StakeId.Sign() <= 0 {
		return errors.New("StakeId 无效（必须大于 0）")
	}
	if e.Raw.TxHash == (common.Hash{}) {
		return errors.New("交易哈希为空")
	}
	return nil
}

// Convert 转换为数据库模型（WithdrawRecord）
func (p *WithdrawnProcessor) Convert(event interface{}) (interface{}, error) {
	e := event.(*contract.StakeContractWithdrawn)
	return &models.WithdrawRecord{
		UserAddress: e.User.Hex(),              // 用户地址
		StakeID:     e.StakeId.String(),        // StakeId
		Principal:   e.Principal.String(),      // 本金（wei）
		TotalAmount: e.TotalAmount.String(),    // 总金额（本金+利息，wei）
		TxHash:      e.Raw.TxHash.Hex(),        // 交易哈希（唯一索引）
		Timestamp:   uint64(time.Now().Unix()), // 入库时间戳（秒级）
		EventType:   EventTypeWithdrawn,
	}, nil
}

// Save 入库（带交易哈希去重）
func (p *WithdrawnProcessor) Save(tx *gorm.DB, model interface{}) error {
	record := model.(*models.WithdrawRecord)
	// 检查交易哈希是否已存在
	var existing models.WithdrawRecord
	if err := tx.Where("tx_hash = ?", record.TxHash).First(&existing).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("查询提现记录失败: ", err)
			return err
		}
		// 新记录，执行插入
		return tx.Create(record).Error
	}
	fmt.Println("提现事件已存在（重复推送）", "tx_hash", record.TxHash)
	return nil
}

// ------------------------------
// 9. 初始化：注册事件处理器
// ------------------------------
func init() {
	// 提前注册处理器（可选，StartEventMonitor 中已注册）
	RegisterProcessor(EventTypeStaked, &StakedProcessor{})
	RegisterProcessor(EventTypeWithdrawn, &WithdrawnProcessor{})
}
