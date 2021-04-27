package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"

	"github.com/gin-gonic/gin"
)

func GetMergeEarnstats(c *gin.Context) {
	var params struct {
		AccountParams
		MergeType string `json:"mergeType" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetMergeEarnstats(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := make(map[string]model.MergeEarnstats)
		stats := d["earnstats"]
		res["earnstats"] = model.MergeEarnstats{
			EarningsYesterday: stats.EarningsYesterday,
			EarningsToday:     stats.EarningsToday,
			Unpaid:            stats.Unpaid,
			Paid:              stats.Paid,
			EarnUnit:          params.MergeType,
		}
		output.Succ(c, res)
	}
}

func GetMergeEarnHistory(c *gin.Context) {
	var params struct {
		AccountParams
		PageParams
		MergeType string `json:"mergeType" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetMergeEarnHistory(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string][]btcpoolclient.EarnHistory)
		// res["list"] = d["list"]
		output.Succ(c, d)
	}
}

func GetEarnstats(c *gin.Context) {
	var params struct {
		AccountParams
		MergeType string `json:"mergeType" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetEarnstats(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := make(map[string]model.Earnstats)
		stats := d["earnstats"]
		res["earnstats"] = model.Earnstats{
			EarningsYesterday: stats.EarningsYesterday,
			EarningsToday:     stats.EarningsToday,
			Unpaid:            stats.Unpaid,
			Paid:              stats.Paid,
			EarnUnit:          params.MergeType,
		}
		output.Succ(c, res)
	}
}

func GetEarnHistory(c *gin.Context) {
	var params struct {
		AccountParams
		PageParams
		IsDecimal string `json:"is_decimal" binding:"-"`
		AccessKey string `json:"access_key" binding:"-"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	params.IsDecimal = "1"
	if d, err := btcpoolclient.GetEarnHistory(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetDashboardHome(c *gin.Context) {

}
