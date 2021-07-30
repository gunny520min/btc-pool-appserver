package controller

import (
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

// HomeInfo 首页数据
func HomeInfo(c *gin.Context) {
	p := service.HomeService
	langStr := c.GetHeader("Accept-Language")
	if len(langStr) == 0 {
		langStr = "en-us"
	}
	p.Lan = langStr

	res := make(map[string]interface{})
	// module  小工具
	res["module"] = p.GetModules(c)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	// banner
	go func() {
		defer wg.Done()
		res["banner"] = p.GetBanner(c)
	}()
	// notice
	go func() {
		defer wg.Done()
		res["notice"] = p.GetNotice(c)
	}()
	// coin list
	var coinErr error = nil
	go func() {
		defer wg.Done()
		res["coin"], coinErr = p.GetCoinList(c)
	}()
	// pool rank
	var rankErr error = nil
	go func() {
		defer wg.Done()
		res["poolRank"], rankErr = service.ExplorerService.GetPoolRank(c, "btc")
	}()

	wg.Wait()
	if coinErr != nil {
		output.ShowErr(c, coinErr)
	} else if rankErr != nil {
		output.ShowErr(c, rankErr)
	} else {
		output.Succ(c, res)
	}
}

// LinkData 链上数据
func LinkData(c *gin.Context) {
	coinType := strings.ToLower(c.Query("coinType"))
	if len(coinType) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	res := make(map[string]interface{})
	// coin info
	if list, err := service.HomeService.GetCoinList(c); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		for _, info := range list {
			if strings.ToLower(info.Id) == strings.ToLower(coinType) {
				res["coinInfo"] = info
				break
			}
		}
		if res["coinInfo"] == nil {
			output.ShowErr(c, errs.ApiErrSystem) // 应该能解析到响应的coin
			return
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//pool rank
	var rankErr error = nil
	go func() {
		defer wg.Done()
		res["poolRank"], rankErr = service.ExplorerService.GetPoolRank(c, coinType)
	}()
	// latest blocks
	var latestBlockErr error = nil
	go func() {
		defer wg.Done()
		res["blocks"], latestBlockErr = service.ExplorerService.GetLatestBlocks(c, coinType)
	}()
	wg.Wait()

	if rankErr != nil {
		output.ShowErr(c, rankErr)
	} else if latestBlockErr != nil {
		output.ShowErr(c, latestBlockErr)
	} else {
		output.Succ(c, res)
	}
}

// PoolHashrates 矿池图表数据
func PoolHashrates(c *gin.Context) {

	type PoolHashrateParam struct {
		CoinType  string `form:"coinType" binding:"required"`
		Dimension string `form:"dimension" binding:"required"`
	}
	var pa PoolHashrateParam
	if err := c.ShouldBindQuery(&pa); err != nil {
		output.ShowErr(c, err)
		return
	}

	res := make(map[string]interface{})
	params := make(map[string]interface{})
	params["dimension"] = "1h" //pa.Dimension
	params["count"] = 72
	params["real_point"] = "1"
	params["coin_type"] = pa.CoinType
	//var p struct {
	//	Dimension string `json:"dimension"`
	//	Count     int    `json:"count"`
	//	RealPoint string `json:"real_point"`
	//	CoinType  string `json:"coin_type"`
	//}
	//p.CoinType = coin
	//p.Count = 72
	//p.Dimension = "1h"
	//p.RealPoint = "1"
	if data, err := service.HomeService.GetShareHashrate(c, params); err != nil {
		output.ShowErr(c, err)
	} else {
		res["histories"] = data.Tickets
		res["unit"] = data.Unit
		output.Succ(c, res)
	}
}

// GetPoolBaseInfo 矿池数据
func GetPoolBaseInfo(c *gin.Context) {
	coin := c.Query("coinType")
	if len(coin) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if info, err := service.HomeService.GetPoolDataBaseInfo(c, coin); err != nil {
		output.ShowErr(c, err)
	} else {
		res := make(map[string]interface{})
		res["baseInfo"] = info
		output.Succ(c, res)
	}
}
