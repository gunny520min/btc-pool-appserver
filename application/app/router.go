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

// InitPublicRouter ..
func InitPublicRouter(r *gin.Engine) {
	pGroup := r.Group("/api/public")
	pGroup.GET("/home/index", controller.HomeCoinInfoList)
	pGroup.GET("/home/bannerNotice", controller.HomeBannerNotice)
	pGroup.GET("/home/latestBlock", controller.ExplorerLatestBlock)
	pGroup.GET("/home/poolrank", controller.ExplorerPoolRank)
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
	pgroup.GET("/info", controller.UpdateSubaccountPayAddress)
	pgroup.GET("/payset", controller.GetSubaccountPayset)
	pgroup.POST("/address/update", controller.UpdateSubaccountPayAddress)
	pgroup.POST("/paylimit/update", controller.UpdateSubaccountPayAddress)
	pgroup.GET("/minerConfig", controller.UpdateSubaccountPayAddress)
	pgroup.GET("/earnStats", controller.UpdateSubaccountPayAddress)
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
