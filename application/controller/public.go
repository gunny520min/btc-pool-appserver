package controller

import (
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"btc-pool-appserver/application/service"
	"strings"

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
	coin := GetParam(c, "coin")
	res := make(map[string]interface{})
	if strings.Contains("btc,bch", strings.ToLower(coin)) {
		// get pool rank
		// get latest blocks
		poolrankCh := service.PublicService.AsnycGetPoolRank(c, coin)
		latestBlockCh := service.PublicService.AsnycGetLatestBlocks(c, coin)
		poolrank := <-poolrankCh
		latestBlock := <-latestBlockCh

		res["blocks"] = service.PublicService.FormatLatestBlockList(latestBlock)
		res["poolRank"] = service.PublicService.FormatPoolRankList(poolrank)
	} else {
		res["blocks"] = make([]model.LatestBlock, 0)
		res["poolRank"] = make([]model.PoolRank, 0)
	}
	res["coinList"] = service.PublicService.FormatHomeCoinList(multiCoin, income)
	output.Succ(c, res)
}
