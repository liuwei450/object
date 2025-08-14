package models

import (
	"gorm.io/gorm"
	"time"
)

// WithdrawRecord 提现事件记录
type WithdrawRecord struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement"`
	UserAddress string         `gorm:"size:66;index:idx_user_addr;not null"`
	StakeID     string         `gorm:"size:64;index:idx_stake_id;not null"`
	Principal   string         `gorm:"size:64;not null"`                         // 本金（wei）
	TotalAmount string         `gorm:"size:64;not null"`                         // 总金额（本金+利息，wei）
	TxHash      string         `gorm:"size:66;uniqueIndex:idx_tx_hash;not null"` // 交易哈希（唯一）
	Timestamp   uint64         `gorm:"index:idx_timestamp;not null"`             // 入库时间戳（秒）
	EventType   string         `gorm:"size:20;not null"`                         // 事件类型（Withdrawn）
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
