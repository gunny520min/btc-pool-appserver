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

	// get notice

	// 获取数据， 判断是否发生了错误
	cps := <-cpsChan

	// TODO:

	// res
	res := make(map[string]interface{})
	res["banner_list"] = service.PublicService.FormatBannerList(cps, lang)

	output.Succ(c, res)
}
