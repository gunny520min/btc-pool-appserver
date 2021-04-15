package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"

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
	Puid       AccountParams `json:"puid" bind:"required"`
	NewAddress string        `json:"newAddress" bind:"required"`
	VerifyMode string        `json:"verifyMode" bind:"required"`
	VerifyId   string        `json:"verifyId" bind:"required"`
	VerifyCode string        `json:"verifyCode" bind:"required"`
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
	Amount   string `json:"amount" bind:"required"`
	CoinType string `json:"coinType" bind:"required"`
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

// func GetAccountEarnstats(c *gin.Context) {
// 	var params AccountParams
// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		output.ShowErr(c, errs.ApiErrParams)
// 		return
// 	}
// 	if d, err := btcpoolclient.GetSubaccountPayset(c, params); err != nil {
// 		output.ShowErr(c, err)
// 		return
// 	} else {
// 		res := make(map[string]interface{})
// 		res["payset"] = d["payset"]
// 		output.Succ(c, res)
// 	}
// }
