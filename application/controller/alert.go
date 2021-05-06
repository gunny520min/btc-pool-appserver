package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"

	"github.com/gin-gonic/gin"
)

func GetAlertSetting(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := service.AlertService.GetAlertSettings(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

// func GetAlertContacts(c *gin.Context) {
// 	var params AccountParams
// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		output.ShowErr(c, errs.ApiErrParams)
// 		return
// 	}
// 	if d, err := btcpoolclient.AlertContacts(c, params); err != nil {
// 		output.ShowErr(c, err)
// 		return
// 	} else {
// 		output.Succ(c, d)
// 	}
// }

func UpdateAlertHashrate(c *gin.Context) {
	var params struct {
		AccountParams
		Hashrate string `form:"hashrate" binding:"required"`
		Unit     string `form:"unit" binding:"required"`
		Enabled  string `form:"enabled" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateAlertHashrate(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func UpdateAlertMiners(c *gin.Context) {
	var params struct {
		AccountParams
		Miners  string `form:"miners" binding:"required"`
		Enabled string `form:"enabled" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateAlertMiners(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func UpdateAlertInterval(c *gin.Context) {
	var params struct {
		AccountParams
		Interval string `form:"interval" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateAlertInterval(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func DeleteAlertContact(c *gin.Context) {
	var params struct {
		AccountParams
		Id string `form:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.DeleteAlertContact(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func CreateAlertContact(c *gin.Context) {
	var params struct {
		AccountParams
		Note        string `form:"note" binding:"required"`
		Email       string `form:"email" binding:"required"`
		Region_code string `form:"region_code" binding:"required"`
		Phone       string `form:"phone" binding:"required"`
		Country     string `form:"country" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.CreateAlertContact(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func UpdateAlertContact(c *gin.Context) {
	var params struct {
		AccountParams
		Note        string `form:"note" binding:"required"`
		Email       string `form:"email" binding:"required"`
		Region_code string `form:"region_code" binding:"required"`
		Phone       string `form:"phone" binding:"required"`
		Id          string `form:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateAlertContact(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

/// 报警列表
func GetAlertList(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetAlerMerge(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func AlertRead(c *gin.Context) {
	var params struct {
		AccountParams
		LogId string `form:"log_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.AlertRead(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}
