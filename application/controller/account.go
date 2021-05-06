package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"fmt"

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
	Puid       AccountParams `form:"puid" binding:"required"`
	NewAddress string        `form:"newAddress" binding:"required"`
	VerifyMode string        `form:"verifyMode" binding:"required"`
	VerifyId   string        `form:"verifyId" binding:"required"`
	VerifyCode string        `form:"verifyCode" binding:"required"`
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
	Amount   string `form:"amount" binding:"required"`
	CoinType string `form:"coinType" binding:"required"`
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

func GetSubacountHiidenList(c *gin.Context) {
	fmt.Println("not implemented")
	// var params struct {
	// 	AccountParams
	// }
	// if err := c.ShouldBindJSON(&params); err != nil {
	// 	output.ShowErr(c, errs.ApiErrParams)
	// 	return
	// }
	// if d, err := btcpoolclient.SubacountHiiden(c, params); err != nil {
	// 	output.ShowErr(c, err)
	// 	return
	// } else {
	// 	output.Succ(c, d)
	// }
}

func SetSubacountHiiden(c *gin.Context) {
	var params struct {
		AccountParams
		HiddenPuid string `form:"hidden_puid" binding:"required"`
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
		CancleHiddenPuid string `form:"cancle_hidden_puid" binding:"required"`
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
