package schema

// 请求体结构（JSON格式）
type WithdrawRequest struct {
	UserAddress string `json:"user_address" binding:"required,len=66"` // 用户钱包地址（0x+64位哈希，共66字符）
	StakeID     string `json:"stake_id" binding:"required"`            // 质押ID（需提现的特定质押记录ID）
	Signature   string `json:"signature" binding:"required"`           // 钱包签名（验证用户身份，防止他人盗用地址提现）
}

// 响应结构
type WithdrawResponse struct {
	Code    int    `json:"code"`    // 状态码：0成功，非0失败
	Message string `json:"message"` // 提示信息
	TxHash  string `json:"tx_hash"` // 提现交易哈希（成功时返回）
}
