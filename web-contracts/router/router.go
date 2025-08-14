package router

import (
	"github.com/gin-gonic/gin"
	"web-contracts/controller"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 1. 全局中间件（可选，如日志、recovery）
	router.Use(gin.Logger(), gin.Recovery())

	// 2. 注册路由组（按业务模块划分，如 /api/v1）
	api := router.Group("/api/v1")
	{
		//查询合约授权
		api.POST("", controller.HandleApprove)
		// 请求 /api/v1/allowance，质押合约授权
		api.POST("/allowance", controller.HandleAllowance)
		//查余额
		api.POST("/balance", controller.HandleBalance)
		//质押
		api.POST("/stake", controller.HandleStake)
		//提现
		api.POST("/withdraw", controller.HandleWithdraw)
	}
	return router
}
