package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"fmt"
	"github.com/gin-gonic/gin"
)

//
//import (
//	"btc-pool-appserver/application/btcpoolclient"
//	"btc-pool-appserver/application/library/errs"
//	"btc-pool-appserver/application/library/output"
//	"btc-pool-appserver/application/model"
//	"btc-pool-appserver/application/service"
//	"fmt"
//	"sync"
//
//	"github.com/gin-gonic/gin"
//)
//
//// HomeIndex 首页数据
//func HomeBannerNotice(c *gin.Context) {
//	// cpsParams := make(map[string]interface{})
//	// get banner
//	// cpsChan := service.PublicService.AsyncGetBanner(c, cpsParams)
//	// bannerRes := <-cpsChan
//
//	// // get notice
//	// noticeChan := service.PublicService.AsyncGetNotice(c, cpsParams)
//	// noticeRes := <-noticeChan
//
//	// 获取数据， 判断是否发生了错误
//	// TODO:
//
//	lang := GetLang(c)
//	// res
//	res := make(map[string]interface{})
//	if d, err := service.PublicService.GetBannerAndNotice(c, ""); err != nil {
//		output.ShowErr(c, err)
//		return
//	} else {
//		res["banner"] = service.PublicService.FormatBannerList(d.Banner, lang)
//		res["notice"] = service.PublicService.FormatNoticeList(d.Notice, lang)
//	}
//
//	output.Succ(c, res)
//}
//
//func GetObserverBanner(c *gin.Context) {
//	lang := GetLang(c)
//	res := make(map[string]interface{})
//	para := struct {
//		WatcherBanner string `"json":"watcher_banner" "form":"watcher_banner"`
//	}{
//		"1",
//	}
//	if bannerList, err := btcpoolclient.GetBannerList(c, para); err != nil {
//		output.ShowErr(c, err)
//	} else {
//		res["banner"] = service.PublicService.FormatBannerList(bannerList, lang)
//	}
//	output.Succ(c, res)
//}
//
//func HomeCoinInfoList(c *gin.Context) {
//
//	res := make(map[string]interface{})
//	var wg sync.WaitGroup
//	wg.Add(4)
//	var coinStats map[string]btcpoolclient.CoinStat
//	var coinIncome btcpoolclient.CoinIncomList
//	var errS error = nil
//	var errI error = nil
//	//get mul coin stat
//	go func() {
//		defer wg.Done()
//		coinStats, errS = service.PublicService.GetMultiCoinStats(c)
//
//	}()
//	go func() {
//		defer wg.Done()
//		coinIncome, errI = service.PublicService.GetAllCoinIncome(c)
//
//	}()
//	go func() {
//		defer wg.Done()
//
//		p := make(map[string]string)
//		p["coins"] = "btc,bch"
//		p["show_unconfirm_info"] = "true"
//		if blocks, err := service.PublicService.GetLatestBlocks(c, p); err != nil {
//
//		} else {
//			res["blocks"] = blocks
//		}
//
//	}()
//	go func() {
//		defer wg.Done()
//		p := make(map[string]string)
//		p["coins"] = "btc,bch"
//		p["from"] = c.GetHeader("platform")
//		if poolRank, err := service.PublicService.GetPoolRank(c, p); err != nil {
//		} else {
//			res["poolRank"] = poolRank
//		}
//
//	}()
//	wg.Wait()
//	if errS == nil && errI == nil {
//		// 成功
//		res["coinList"] = service.PublicService.FormatHomeCoinList(coinStats, coinIncome)
//		output.Succ(c, res)
//	} else {
//		if errS != nil {
//			output.ShowErr(c, errS)
//		} else {
//			output.ShowErr(c, errI)
//		}
//	}
//}
//
//func GetHomeHashrateHistory(c *gin.Context) {
//	// params := make(map[string]interface{})
//	res := make(map[string]interface{})
//	// params["dimension"] = "1h"
//	// params["count"] = 72
//	// params["real_point"] = "1"
//
//	coin := c.Query("coin")
//	if len(coin) == 0 {
//		output.ShowErr(c, errs.ApiErrParams)
//	} else {
//		fmt.Println("get hashrate")
//		// params["coin_type"] = coin
//		var p struct {
//			Dimension string `json:"dimension"`
//			Count     int    `json:"count"`
//			RealPoint string `json:"real_point"`
//			CoinType  string `json:"coin_type"`
//		}
//		p.CoinType = coin
//		p.Count = 72
//		p.Dimension = "1h"
//		p.RealPoint = "1"
//
//		// if p, err := urlEncoded(params); err != nil {
//		// 	res["histories"] = make(map[string]interface{})
//		// 	res["unit"] = ""
//		// 	output.Succ(c, res)
//		// } else {
//		shareData := service.PoolService.GetShareHashrate(c, p)
//		res["histories"] = service.PoolService.FormatHashrateChartData(shareData)
//		res["unit"] = service.PoolService.FormatHashrateChartUnit(shareData) + model.GetCoinSuffixByCoinType(coin)
//		output.Succ(c, res)
//		// }
//	}
//}
//
func GetCaptcha(c *gin.Context) {
	var params struct {
		Puid string `form:"puid" binding:"required"`
		Type string `form:"type" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		fmt.Println(err)
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
	p["puid"] = params.Puid
	p["lang"] = l
	if d, err := btcpoolclient.GetCpatcha(c, p, params.Type); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, d)
	}
}
