package models

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type StakeRecord struct {
	ID          uint      `gorm:"primary_key;auto_increment;comment:'自增主键'"`
	UserAddress string    `gorm:"size:80;index;not null;comment:'用户钱包地址（0x开头，64位哈希+0x前缀）'"`     // 用户钱包地址（0x开头）
	StakeID     string    `gorm:"size:100;not null;comment:'质押代币ID（如ERC-20的合约地址或NFT的TokenID）'"` // 质押的代币ID
	Amount      string    `gorm:"size:64;not null;comment:'质押数量（wei单位，字符串格式避免溢出）'"`             // 质押数量
	Period      uint8     `gorm:"default:0;comment:'质押周期'"`                                     //质押周期
	Timestamp   uint64    `gorm:"not null;comment:'区块链区块时间戳（单位秒）'"`                             //质押时间戳
	TxHash      string    `gorm:"size:66;uniqueIndex;comment:'质押交易哈希（唯一，防止重复记录）'"`              // 新增：交易哈希唯一索引
	EventType   string    `gorm:"size:20;not null"`                                             // 事件类型（Staked）
	CreatedAt   time.Time `gorm:"comment:'记录创建时间（数据库本地时间）'"`
}

// TableName 自定义表名并添加表注释
func (StakeRecord) TableName() string {
	return "stake_records"
}
