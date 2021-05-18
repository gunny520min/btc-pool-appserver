package controller

import (
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取最新出块列表
type ExplorerParams struct {
	App_a     string `json:"app_a" form:"app_a" binding:"required"`
	App_b     string `json:"app_b" form:"app_b" binding:"required"`
	Coins     string `json:"coins" form:"coins" binding:"required"`
	Nonce     string `json:"nonce" form:"nonce" binding:"required"`
	Sign      string `json:"sign" form:"sign" binding:"required"`
	Timestamp string `json:"timestamp" form:"timestamp" binding:"required"`
}
type LatestBlocksParameters struct {
	ExplorerParams
	ShowUnconfirmInfo string `json:"show_unconfirm_info" form:"show_unconfirm_info" binding:"required"`
}
type PoolRankParameters struct {
	ExplorerParams
	From string `json:"from" form:"from" binding:"required"`
}

func ExplorerLatestBlock(c *gin.Context) {

	var blockParams LatestBlocksParameters
	if err := c.ShouldBindQuery(&blockParams); err != nil {
		fmt.Printf(">> explorer params1 = %v %v\n", blockParams, err)
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	latestBlockCh := service.PublicService.AsnycGetLatestBlocks(c, "", blockParams)
	latestBlock := <-latestBlockCh

	res := service.PublicService.FormatLatestBlockList(latestBlock)
	output.Succ(c, res)
}

func ExplorerPoolRank(c *gin.Context) {

	var poolrankParams PoolRankParameters
	if err := c.ShouldBindQuery(&poolrankParams); err != nil {
		fmt.Printf(">> explorer params2 = %v %v\n", poolrankParams, err)
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	// coin := c.Query("coin")
	// if coin == "" {
	// 	output.ShowErr(c, errs.ApiErrParams)
	// 	return
	// }
	poolrankCh := service.PublicService.AsnycGetPoolRank(c, "", poolrankParams)
	poolrank := <-poolrankCh

	res := service.PublicService.FormatPoolRankList(poolrank)
	output.Succ(c, res)
}
