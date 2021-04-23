package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"btc-pool-appserver/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func MultiCoinStats(c *gin.Context) {

}

func HashrateHistory(c *gin.Context) {
	params := make(map[string]interface{})
	res := make(map[string]interface{})
	params["dimension"] = "1h"
	params["count"] = 72
	params["real_point"] = "1"

	coin := c.Query("coin")
	if len(coin) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
	} else {
		fmt.Println("get hashrate")
		params["coin_type"] = coin

		if p, err := urlEncoded(params); err != nil {
			res["histories"] = make(map[string]interface{})
			res["unit"] = ""
			output.Succ(c, res)
		} else {
			shareData := service.PoolService.GetShareHashrate(c, p)
			res["histories"] = service.PoolService.FormatHashrateChartData(shareData)
			res["unit"] = service.PoolService.FormatHashrateChartUnit(shareData)
			output.Succ(c, res)
		}
	}
}

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
		res := make(map[string][]btcpoolclient.EarnHistory)
		res["list"] = d["list"]
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
	if d, err := btcpoolclient.GetMergeEarnHistory(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := make(map[string][]btcpoolclient.EarnHistory)
		res["list"] = d["list"]
		output.Succ(c, res)
	}
}
