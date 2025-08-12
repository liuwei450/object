package schema

// 查询授权额度
type AllowanceRequest struct {
	Owner   string `form:"owner" json:"owner" binding:"required"`     // 拥有者地址
	Spender string `form:"spender" json:"Spender" binding:"required"` // 被授权者地址
}

// 授权
type ApproveRequest struct {
	Spender string `form:"Spender" json:"spender" binding:"required"` // 被授权的地址
	Value   string `form:"Value" json:"value" binding:"required"`     // 授权数量（按 token decimals 传）
}

// 余额
type BalanceRequest struct {
	BalanceAdress string `form:"BalanceAdress" json:"BalanceAdress" binding:"required"` // 查询地址
}
