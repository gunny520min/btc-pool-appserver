package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/lang"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"btc-pool-appserver/application/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMergeCoinInfo(c *gin.Context) {
	var params struct {
		AccountParams
		CoinType string `form:"coinType" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	langC := GetLang(c)
	if data, err := service.MergeService.GetMergeCoinData(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		infos := service.MergeService.FormatMergeCoinInfos(data, langC)
		coinType := params.CoinType
		btcBchMerge := strings.ToLower(coinType) == "btc" || strings.ToLower(coinType) == "bch"
		ltcMerge := strings.ToLower(coinType) == "ltc"
		res := make(map[string]interface{})
		if btcBchMerge {
			res["desc"] = lang.Trans("merge_coin_desc1", langC, "")
		} else if ltcMerge {
			res["desc"] = lang.Trans("merge_coin_desc2", langC, "")
		}
		list := make([]model.MergeCoin, 0)

		var filterCoins []string
		if btcBchMerge {
			filterCoins = []string{"nmc", "ela", "vcash"}
		} else if ltcMerge {
			filterCoins = []string{"doge"}
		}

		for _, filterCoin := range filterCoins {
			if i := infos[filterCoin]; i != nil {
				list = append(list, *i)
			}
		}
		res["list"] = list

		output.Succ(c, res)
	}
}

func UpdateMergeCoinAddress(c *gin.Context) {
	var params struct {
		AccountParams
		CoinType   string `form:"coinType" binding:"required"`
		NewAddress string `form:"newAddress" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateMergeCoinAddress(c, params); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}
