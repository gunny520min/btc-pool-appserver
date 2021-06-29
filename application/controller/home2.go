package controller

import (
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"github.com/gin-gonic/gin"
)

func HomeInfo(c *gin.Context) {
	if info,err := service.HomeService.GetHomeInfo(c); err != nil {
		output.ShowErr(c, err)
	} else  {
		output.Succ(c, info)
	}
}
