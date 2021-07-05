package app

import (
	"btc-pool-appserver/application/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter(r *gin.Engine) error {
	// home页面相关
	InitHomeRouter(r)
	// 用户面板相关
	//InitDashboardRouter(r)
	//// app相关
	//InitAppRouter(r)
	//// 子账户
	InitAccountSubaccount(r)
	//// 矿机
	//InitMiningWorker(r)
	//// 报警
	//InitAlert(r)
	//// 观察者链接
	//InitWatcher(r)
	//// 合并挖矿
	//InitMergeMining(r)
	return nil
}

// InitHomeRouter 首页 ..
func InitHomeRouter(r *gin.Engine) {
	pGroup := r.Group("/home")
	pGroup.GET("/info", controller.HomeInfo)
	pGroup.GET("/linkData", controller.LinkData)
	pGroup.GET("/pool/hashrates", controller.PoolHashrates)
	pGroup.GET("/pool/data", controller.GetPoolBaseInfo)
}

// 用户面板
func InitDashboardRouter(r *gin.Engine) {
	//pGroup := r.Group("/dashboard")
	//pGroup.GET("/baseInfo", controller.GetDashboardHome)
	//pGroup.GET("/workerChart", controller.GetDashboardWorkerShareHistory)
}

// InitAppRouter ...
func InitAppRouter(r *gin.Engine) {
	//pGroup := r.Group(/*"/api")
	//pGroup.GET("/app/version", controller.AppVersion)
	//pGroup.GET("/app/urlConfig", controller.UrlConfig)
	//pGroup.POST(*/"/app/getCaptcha", controller.GetCaptcha)
}

/// 子账户相关
func InitAccountSubaccount(r *gin.Engine) {
	pgroup := r.Group("/subaccount")
	//// 子账户
	pgroup.GET("/list", controller.GetSubAccountList)
	pgroup.GET("/info", controller.GetSubAccountInfo)
	pgroup.GET("/hashrate", controller.GetSubAccountHashrates)
	pgroup.GET("/changeHashrate", controller.ChangeSubAccountHashrate)
	//// 收款地址管理
	//pgroup.GET("/info", controller.GetAccountInfo)
	//pgroup.GET("/addressPayset", controller.GetSubaccountPayset)
	//pgroup.POST("/address/update", controller.UpdateSubaccountPayAddress)
	//pgroup.POST("/payLimit/update", controller.UpdateSubaccountPayLimit)
	//// 挖矿配置
	//pgroup.GET("/minerConfig", controller.GetAccountMinerConfig)
	//// 收益
	//pgroup.GET("/mining/earnStats", controller.GetEarnstats)
	//pgroup.GET("/mining/earnHistory", controller.GetEarnHistory)
	//pgroup.GET("/mining/merge/earnStats", controller.GetMergeEarnstats)
	//pgroup.GET("/mining/merge/earnHistory", controller.GetMergeEarnHistory)
	// 隐藏子账户
	pgroup.GET("/hidden", controller.SetSubaccountHidden)
	pgroup.GET("/hiddenCancel", controller.CancelSubaccountHidden)
}

/// 报警相关
func InitAlert(r *gin.Engine) {
	//pGroup := r.Group("/alert")
	//pGroup.GET("/settings", controller.GetAlertSetting)
	//pGroup.GET("/hashrate/update", controller.UpdateAlertHashrate)
	//pGroup.GET("/workerCount/update", controller.UpdateAlertMiners)
	//pGroup.GET("/interval/update", controller.UpdateAlertInterval)
	//pGroup.GET("/contact/delete", controller.DeleteAlertContact)
	//pGroup.GET("/contact/update", controller.UpdateAlertContact)
	//// 报警列表和上报已读
	//pGroup.GET("/list", controller.GetAlertList)
	//pGroup.GET("/read", controller.AlertRead)
}

/// 挖矿矿机相关
func InitMiningWorker(r *gin.Engine) {
	//pgroup := r.Group("/mining/worker/")
	//pgroup.GET("/groups", controller.GetMinerGroups)
	//pgroup.GET("/groups/delete", controller.MinerGroupDelete)
	//pgroup.POST("/groups/create", controller.MinerGroupCreate)
	//pgroup.POST("/delete", controller.MinerWorkerDelete)
	//pgroup.GET("/move", controller.MinerWorkerMove)
	//pgroup.GET("/list", controller.GetMinerWorkerList)
	//pgroup.GET("/detail", controller.GetMinerWorkerDetail)
	//pgroup.GET("/hashrate", controller.GetMinerWorkerHashrate)
}

/// 观察者链接相关
func InitWatcher(r *gin.Engine) {
	//pgroup := r.Group("/watcher")
	//pgroup.GET("/list", controller.GetWatcherList)
	//pgroup.GET("/create", controller.CreateWatcher)
	//pgroup.GET("/delete", controller.DeleteWatcher)
	//pgroup.GET("/update", controller.UpdateWatcher)
	//pgroup.GET("/authority", controller.WatcherAuthority)
	//pgroup.GET("/check", controller.AddOtherWatcher)
	//pgroup.GET("/hashrate", controller.GetWatcherHashrate)
	//pgroup.GET("/dashboard", controller.GetWatcherDashboard)
	//pgroup.GET("/dashboard/workerChart", controller.GetWatcherDashboardWorkerShareHistory)
	//pgroup.GET("/earnInfo", controller.GetWatcherEarnStats)
	//pgroup.GET("/earnHistory", controller.GetWatcherEarnHistory)
	//pgroup.GET("/observerBanner", controller.GetObserverBanner)
}

func InitMergeMining(r *gin.Engine) {
	//pgroup := r.Group("/merge")
	//pgroup.GET("/coinList", controller.GetMergeCoinInfo)
	//pgroup.GET("/updateAddress", controller.UpdateMergeCoinAddress)
}
