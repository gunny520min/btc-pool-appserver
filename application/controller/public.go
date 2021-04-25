package controller

import (
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// HomeIndex 首页数据
func HomeBannerNotice(c *gin.Context) {

	cpsParams := make(map[string]interface{})
	lang := GetLang(c)
	// cpsParams["platform"] = 1
	// if GetLang(c) == "zh_CN" {

	// }
	// get banner
	cpsChan := service.PublicService.AsyncGetBanner(c, cpsParams)
	bannerRes := <-cpsChan

	// get notice
	noticeChan := service.PublicService.AsyncGetNotice(c, cpsParams)
	noticeRes := <-noticeChan

	// 获取数据， 判断是否发生了错误

	// TODO:

	// res
	res := make(map[string]interface{})
	res["banner"] = service.PublicService.FormatBannerList(bannerRes, lang)
	res["notice"] = service.PublicService.FormatNoticeList(noticeRes, lang)

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
