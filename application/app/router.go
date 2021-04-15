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
	return nil
}

// InitPublicRouter ..
func InitPublicRouter(r *gin.Engine) {
	pGroup := r.Group("/api/public")
	pGroup.GET("/home/bannerNotice", controller.HomeBannerNotice)
	pGroup.GET("/home/index", controller.HomeCoinInfoList)
}

// InitAppRouter ...
func InitAppRouter(r *gin.Engine) {
	pGroup := r.Group("/api/operation")
	pGroup.GET("/app/version", controller.AppVersion)
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
