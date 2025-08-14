package main

import (
	"log"
	"web-contracts/mysqlbase"
	"web-contracts/router"
	"web-contracts/service"
)

func main() {

	// 初始化数据库
	mysqlbase.InitDb()
	// 2. 程序退出时自动关闭数据库连接（确保资源释放）
	defer func() {
		if err := mysqlbase.CloseConn(); err != nil {
			log.Fatal("关闭数据库连接失败: ", err)
		}
	}()

	//初始化客户端
	if err := service.InitClient(); err != nil {
		log.Fatalf("初始化客户端失败，程序退出: %v", err) // 初始化失败时终止程
	}
	//合约调用
	router.InitRouter()
	// 启动事件监听协程
	go service.StartEventMonitor()
	//

}
