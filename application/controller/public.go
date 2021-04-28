package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// HomeIndex 首页数据
func HomeBannerNotice(c *gin.Context) {
	// cpsParams := make(map[string]interface{})
	// get banner
	// cpsChan := service.PublicService.AsyncGetBanner(c, cpsParams)
	// bannerRes := <-cpsChan

	// // get notice
	// noticeChan := service.PublicService.AsyncGetNotice(c, cpsParams)
	// noticeRes := <-noticeChan

	// 获取数据， 判断是否发生了错误
	// TODO:

	lang := GetLang(c)
	// res
	res := make(map[string]interface{})
	if d, err := service.PublicService.GetBannerAndNotice(c, ""); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res["banner"] = service.PublicService.FormatBannerList(d.Banner, lang)
		res["notice"] = service.PublicService.FormatNoticeList(d.Notice, lang)
	}

	output.Succ(c, res)
}

func HomeCoinInfoList(c *gin.Context) {
	//get mul coin stat
	multiCoinChan := service.PublicService.AsnycGetMultiCoinStats(c)
	incomeChan := service.PublicService.AsyncGetAllCoinIncome(c)
	multiCoin := <-multiCoinChan
	income := <-incomeChan
	// coin := GetParam(c, "coin")
	res := make(map[string]interface{})
	// if strings.Contains("btc,bch", strings.ToLower(coin)) {
	// 	// get pool rank
	// 	// get latest blocks

	// 	res["blocks"] = service.PublicService.FormatLatestBlockList(latestBlock)
	// 	res["poolRank"] = service.PublicService.FormatPoolRankList(poolrank)
	// } else {
	// 	res["blocks"] = make([]model.LatestBlock, 0)
	// 	res["poolRank"] = make([]model.PoolRank, 0)
	// }
	res["coinList"] = service.PublicService.FormatHomeCoinList(multiCoin, income)
	output.Succ(c, res)
}

func GetHomeHashrateHistory(c *gin.Context) {
	// params := make(map[string]interface{})
	res := make(map[string]interface{})
	// params["dimension"] = "1h"
	// params["count"] = 72
	// params["real_point"] = "1"

	coin := c.Query("coin")
	if len(coin) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
	} else {
		fmt.Println("get hashrate")
		// params["coin_type"] = coin
		var p struct {
			Dimension string `json:"dimension"`
			Count     int    `json:"count"`
			RealPoint string `json:"real_point"`
			CoinType  string `json:"coin_type"`
		}
		p.CoinType = coin
		p.Count = 72
		p.Dimension = "1h"
		p.RealPoint = "1"

		// if p, err := urlEncoded(params); err != nil {
		// 	res["histories"] = make(map[string]interface{})
		// 	res["unit"] = ""
		// 	output.Succ(c, res)
		// } else {
		shareData := service.PoolService.GetShareHashrate(c, p)
		res["histories"] = service.PoolService.FormatHashrateChartData(shareData)
		res["unit"] = service.PoolService.FormatHashrateChartUnit(shareData)
		output.Succ(c, res)
		// }
	}
}

func GetCaptcha(c *gin.Context) {
	var params struct {
		AccountParams
		Type string `json:"type:" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	h := c.Request.Header
	l := h.Get("accept-language")
	if len(l) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	var p = make(map[string]string)
	p["puid"] = params.AccountParams.Puid
	p["lang"] = l
	if d, err := btcpoolclient.GetCpatcha(c, p, params.Type); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}
