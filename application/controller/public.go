package controller

import (
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"

	"github.com/gin-gonic/gin"
)

// HomeIndex 首页数据
func HomeIndex(c *gin.Context) {

	cpsParams := make(map[string]interface{})
	lang := GetLang(c)
	// cpsParams["platform"] = 1
	if GetLang(c) == "zh_CN" {

	}
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

	res := make(map[string]interface{})
	res["coinList"] = service.PublicService.FormatHomeCoinList(multiCoin, income)
	output.Succ(c, res)
}
