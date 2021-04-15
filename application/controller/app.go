package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"

	"github.com/gin-gonic/gin"
)

// AppVersion 获取升级包信息
func AppVersion(c *gin.Context) {
	v := c.Query("version")
	if len(v) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	params := ("version=" + v)
	if d, err := btcpoolclient.AppVersionCheck(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

// url config
func UrlConfig(c *gin.Context) {
	if d, err := btcpoolclient.UrlConfig(c, ""); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}
