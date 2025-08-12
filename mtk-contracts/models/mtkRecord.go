package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mtk-contracts/config"
)

var DB *gorm.DB

type StakeEvent struct {
	ID        uint `gorm:"primaryKey"`
	User      string
	StakeID   uint64
	Amount    string
	Period    uint64
	Timestamp uint64
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.MYSQL_DSN), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	DB.AutoMigrate(&StakeEvent{})
}
