package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"

	"github.com/gin-gonic/gin"
)

func GetAccountInfo(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetAccountInfo(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func GetSubaccountPayset(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.GetSubaccountPayset(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := make(map[string]interface{})
		res["payset"] = d["payset"]
		output.Succ(c, res)
	}
}

type UpdateAddressParams struct {
	Puid       AccountParams `json:"puid" binding:"required"`
	NewAddress string        `json:"newAddress" binding:"required"`
	VerifyMode string        `json:"verifyMode" binding:"required"`
	VerifyId   string        `json:"verifyId" binding:"required"`
	VerifyCode string        `json:"verifyCode" binding:"required"`
}

func UpdateSubaccountPayAddress(c *gin.Context) {
	var params UpdateAddressParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.GetSubaccountPayset(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := make(map[string]interface{})
		res["payset"] = d["payset"]
		output.Succ(c, res)
	}
}

type PayLimitParams struct {
	AccountParams
	Amount   string `json:"amount" binding:"required"`
	CoinType string `json:"coinType" binding:"required"`
}

func UpdateSubaccountPayLimit(c *gin.Context) {
	var params PayLimitParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.UpdateSubaccountPayLimit(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		// res["payset"] = d["payset"]
		output.Succ(c, d)
	}
}

func GetAccountMinerConfig(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetAccountMinerConfig(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		// res["payset"] = d["payset"]
		output.Succ(c, d)
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
		res := make(map[string][]btcpoolclient.MergeEarnHistory)
		res["list"] = d["list"]
		output.Succ(c, res)
	}
}

func SetSubacountHiiden(c *gin.Context) {
	var params struct {
		AccountParams
		HiddenPuid string `json:"hidden_puid" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.SubacountHiiden(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func CancelSubacountHiiden(c *gin.Context) {
	var params struct {
		AccountParams
		CancleHiddenPuid string `json:"cancle_hidden_puid" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.SubacountHiidenCancel(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}
