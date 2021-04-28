package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"btc-pool-appserver/application/service"
	"sync"

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
			Paid:              stats.TotalPaid,
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
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if stats, err := service.PoolService.GetDashboardIncome(c, params.Puid); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, stats)
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
	var params struct {
		Puid string `json:"puid" binding:"-"`
	}
	c.ShouldBindJSON(&params)
	var res model.Dashboard

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//notice
		defer wg.Done()
		if _, currentCoinAccount, err := service.PoolService.GetDashboardSubaccounts(c, params.Puid); err != nil {
			output.ShowErr(c, err)
			return
		} else {
			res.IsSmart = currentCoinAccount.IsSmart()
			res.Puid = currentCoinAccount.Puid
			var mCoinType string
			var smartStr string
			if currentCoinAccount.IsSmart() {
				if GetLang(c) == "zh_cn" {
					smartStr = "机枪"
				} else {
					smartStr = "Smart Pool"
				}
				mCoinType = smartStr
			} else {
				mCoinType = currentCoinAccount.CoinType
			}
			res.Title = currentCoinAccount.Name + mCoinType + currentCoinAccount.RegionText
		}
	}()
	wg.Wait()
	wg.Add(4)
	go func() {
		defer wg.Done()
		if income, err := service.PoolService.GetDashboardIncome(c, res.Puid); err != nil {

		} else {
			res.Income = income
		}
	}()

	wg.Wait()

}
