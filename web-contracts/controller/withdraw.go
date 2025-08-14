package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-contracts/schema"
)

func HandleWithdraw(c *gin.Context) {
	var req schema.WithdrawRequest
	// 1. 绑定并校验请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, schema.WithdrawResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

}

