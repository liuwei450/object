package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type StakeRecord struct {
	ID        uint   `gorm:"primaryKey"`
	User      string `gorm:"index"`
	StakeID   string
	Amount    string
	Period    uint8
	Timestamp uint64
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(entity.MYSQL_DSN), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	DB.AutoMigrate(&StakeRecord{})
}
