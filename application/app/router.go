package app

import (
	"btc-pool-appserver/application/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter(r *gin.Engine) error {
	// home页面相关
	InitPublicRouter(r)
	// app相关
	InitAppRouter(r)
	// 子账户
	InitAccountSubaccount(r)
	// 矿机
	InitMiningWorker(r)
	return nil
}

// 首页 ..
func InitPublicRouter(r *gin.Engine) {
	pGroup := r.Group("/api/public")
	pGroup.GET("/home/coinInfoList", controller.HomeCoinInfoList)
	pGroup.GET("/homepage/activities", controller.HomeBannerNotice)
	pGroup.GET("/home/latestBlock", controller.ExplorerLatestBlock)
	pGroup.GET("/home/poolrank", controller.ExplorerPoolRank)
	pGroup.GET("/home/hashrateHistory", controller.GetHomeHashrateHistory)
}

// InitAppRouter ...
func InitAppRouter(r *gin.Engine) {
	pGroup := r.Group("/api")
	pGroup.GET("/app/version", controller.AppVersion)
	pGroup.GET("/app/urlConfig", controller.UrlConfig)
}

/// 子账户相关
func InitAccountSubaccount(r *gin.Engine) {
	pgroup := r.Group("/account/subaccount/")
	// 收款地址管理
	pgroup.GET("/info", controller.GetAccountInfo)
	pgroup.GET("/subaccount/addressPayset", controller.GetSubaccountPayset)
	pgroup.POST("/subaccount/address/update", controller.UpdateSubaccountPayAddress)
	pgroup.POST("/subaccount/payLimit/update", controller.UpdateSubaccountPayLimit)
	// 挖矿配置
	pgroup.GET("/subaccount/minerConfig", controller.GetAccountMinerConfig)
	// 收益
	pgroup.GET("/subaccount/mining/earnStats", controller.GetMergeEarnstats)
	pgroup.GET("/subaccount/mining/earnHistory", controller.GetEarnHistory)
	pgroup.GET("/subaccount/mining/merge/earnStats", controller.GetMergeEarnstats)
	pgroup.GET("/subaccount/mining/merge/earnHistory", controller.GetMergeEarnHistory)
	// 隐藏子账户
	pgroup.GET("/subaccount/hidden", controller.SetSubacountHiiden)
	pgroup.GET("/subaccount/hiddenCancel", controller.CancelSubacountHiiden)
}

/// 挖矿矿机相关
func InitMiningWorker(r *gin.Engine) {
	pgroup := r.Group("/mining/worker/")
	pgroup.GET("/groups", controller.GetMinerGroups)
	pgroup.GET("/groups/delete", controller.MinerGroupDelete)
	pgroup.POST("/groups/create", controller.MinerGroupCreate)
	pgroup.POST("/delete", controller.MinerWorkerDelete)
	pgroup.GET("/move", controller.MinerWorkerMove)
	pgroup.GET("/list", controller.GetMinerWorkerList)
	pgroup.GET("/detail", controller.GetMinerWorkerDetail)
	pgroup.GET("/hashrate", controller.GetMinerWorkerHashrate)
}
