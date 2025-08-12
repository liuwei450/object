package mtk_contracts

import (
	"mtk-contracts/service"
)

func main() {
	//// 初始化客户端
	service.InitClient()

	service.WatchStakeEvents()
	
}
