package main

import (
	"github.com/gin-gonic/gin"
	"web-contracts/models"
	"web-contracts/service"
)

func main() {

	// 初始化数据库
	models.InitDB()
	// 启动事件监听协程
	go service.ListenStakedEvents()
	// 启动 Gin API
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":8080")
}
