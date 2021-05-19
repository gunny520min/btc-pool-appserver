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
	//ExplorerParams
	Coins string `json:"coins" form:"coins" binding:"required"`
}
type PoolRankParameters struct {
	//ExplorerParams
	From  string `json:"from" form:"from" binding:"required"`
	Coins string `json:"coins" form:"coins" binding:"required"`
}

func ExplorerLatestBlock(c *gin.Context) {
	var blockParams LatestBlocksParameters
	if err := c.ShouldBindQuery(&blockParams); err != nil {
		fmt.Printf(">> explorer params1 = %v %v\n", blockParams, err)
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	p := make(map[string]string)
	p["coins"] = blockParams.Coins
	p["show_unconfirm_info"] = "true"
	if res, err := service.PublicService.GetLatestBlocks(c, p); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, res)
	}
}

func ExplorerPoolRank(c *gin.Context) {

	var poolRankParams PoolRankParameters
	if err := c.ShouldBindQuery(&poolRankParams); err != nil {
		fmt.Printf(">> explorer params2 = %v %v\n", poolRankParams, err)
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	p := make(map[string]string)
	p["coins"] = poolRankParams.Coins
	p["from"] = poolRankParams.From
	if res, err := service.PublicService.GetPoolRank(c, p); err!=nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, res)
	}
}
