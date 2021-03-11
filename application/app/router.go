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
	pGroup.POST("/home/index", controller.HomeIndex)
}

// InitAppRouter ...
func InitAppRouter(r *gin.Engine) {
	pGroup := r.Group("/api/operation")
	pGroup.POST("/app/version", controller.AppVersion)
}
