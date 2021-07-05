package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"github.com/gin-gonic/gin"
)

func GetAccountInfo(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindQuery(&params); err != nil {
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
	if err := c.ShouldBindQuery(&params); err != nil {
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
	if err := c.ShouldBindQuery(&params); err != nil {
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
	if err := c.ShouldBindQuery(&params); err != nil {
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
	if err := c.ShouldBindQuery(&params); err != nil {
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

// GetSubAccountList 获取子账户列表
func GetSubAccountList(c *gin.Context) {
	var params struct {
		AccountParams
		IsHidden int `form:"isHidden"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if subaccountList,err := service.AccountService.GetSubAccountList(c, params.Puid, params.IsHidden); err!=nil {
		output.ShowErr(c, err)
	} else {
		output.SuccList(c, subaccountList)
	}
}

// GetSubAccountList 获取子账户列表
func GetSubAccountInfo(c *gin.Context) {
	var params struct {
		AccountParams
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if subaccountList,err := service.AccountService.GetCurrentSubAccount(c, params.Puid); err!=nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, subaccountList)
	}
}

// GetSubAccountHashrates 获取子账户算力
func GetSubAccountHashrates(c *gin.Context) {
	var params struct {
		Puids string `form:"puids"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if subaccountHashrates,err := service.AccountService.GetSubAccountHashrates(c, params.Puids); err!=nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, subaccountHashrates)
	}
}

// ChangeSubAccountHashrate 一键切换
func ChangeSubAccountHashrate(c *gin.Context) {
	var params struct {
		AccountParams
		Source string `form:"source"`
		Dest string `form:"dest"`
	}
	if err := c.ShouldBind(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.SubaccountChangeHashrate(c, params); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}

func SetSubaccountHidden(c *gin.Context) {
	var params struct {
		AccountParams
		HiddenPuid string `form:"hidden_puid" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.SubacountHiiden(c, params); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}

func CancelSubaccountHidden(c *gin.Context) {
	var params struct {
		AccountParams
		CancleHiddenPuid string `form:"cancle_hidden_puid" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.SubacountHiidenCancel(c, params); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}
