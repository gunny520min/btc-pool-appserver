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
	return nil
}

// InitPublicRouter ..
func InitPublicRouter(r *gin.Engine) {
	pGroup := r.Group("/api/public")
	pGroup.GET("/home/index", controller.HomeIndex)
	pGroup.GET("/home/multiCoinStats", controller.MultiCoinStats)
	pGroup.GET("/home/hashrateHistory", controller.HashrateHistory)
}

// InitAppRouter ...
func InitAppRouter(r *gin.Engine) {
	pGroup := r.Group("/api/operation")
	pGroup.GET("/app/version", controller.AppVersion)
}
