package schema

type Stake struct {
	Amount int64 `form:"Amount" json:"amount" ` // 质押数量
	Period uint8 `form:"Period" json:"period"`  // 枚举索引：0,1,2,3
}
