package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"

	"github.com/gin-gonic/gin"
)

func GetWatcherList(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetWatcherList(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func CreateWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		Note        string `form:"note" binding:"required"`
		Lang        string `form:"lang" binding:"required"`
		Authorities string `form:"authorities" binding:"required"`
		GrinValue   string `form:"grin_value" binding:"-"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.CreateWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := d
		if len(params.GrinValue) > 0 {
			res["grin_value"] = params.GrinValue
		} else {
			res["grin_value"] = ""
		}
		output.Succ(c, d)
	}
}

func DeleteWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		WatcherId string `form:"watcher_id" binding:"required"`
		Lang      string `form:"lang" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.DeleteWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func UpdateWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		WatcherId   string `form:"watcher_id" binding:"required"`
		Authorities string `form:"authorities" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func WatcherAuthority(c *gin.Context) {
	var params struct {
		AccountParams
		AccessKey string `form:"access_key" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.WatcherAuthority(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

// 添加查看别人的观察者链接
func AddOtherWatcher(c *gin.Context) {
	var params struct {
		Puids        string `form:"puids" binding:"required"`
		WatcherToken string `form:"watcher_token" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := service.WatcherService.AddOtherWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}
