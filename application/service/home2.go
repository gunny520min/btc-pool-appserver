package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/tool"
	"btc-pool-appserver/application/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"sync"
)

type homeHandler struct{
	lan string
}

var HomeService = &homeHandler{}

func (p *homeHandler) GetHomeInfo(c *gin.Context) (map[string]interface{}, error) {
	langStr := c.GetHeader("Accept-Language")
	if len(langStr) == 0 {
		langStr = "en-us"
	}
	p.lan = langStr

	res := make(map[string]interface{})
	// module  小工具
	res["module"] = p.getModules(c)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	// banner
	go func() {
		defer wg.Done()
		res["banner"] = p.getBanner(c)
	}()
	// notice
	go func() {
		defer wg.Done()
		res["notice"] = p.getNotice(c)
	}()
	// coin list
	var coinErr error = nil
	go func() {
		defer wg.Done()
		res["coin"], coinErr = p.getCoinList(c)
	}()
	// pool rank
	var rankErr error = nil
	go func() {
		defer wg.Done()
		res["poolRank"], rankErr = p.getPoolRank(c)
	}()

	wg.Wait()
	if coinErr != nil {
		return res, coinErr
	} else if rankErr != nil {
		return res, rankErr
	} else {
		return res, nil
	}
}

// getBanner
func (p *homeHandler) getBanner(c *gin.Context) []model.Banner {
	ret := make([]model.Banner, 0)
	if list, err := btcpoolclient.GetBannerList(c, ""); err == nil {
		for _, item := range list {
			banner := model.Banner{}
			banner.Id = item.Id
			banner.Url = item.Link
			banner.ImgUrl = item.Pic
			ret = append(ret, banner)
		}
	}
	return ret
}

// getNotice
func (p *homeHandler) getNotice(c *gin.Context) []model.Notice {
	ret := make([]model.Notice, 0)
	if list, err := btcpoolclient.GetNoticeList(c, ""); err == nil {
		for _, item := range  list {
			notice := model.Notice{}
			notice.Id = item.Id
			notice.Url = item.HtmlUrl
			notice.Content = item.Title
			ret = append(ret, notice)
		}
	}
	return ret
}

// coin list
func (p *homeHandler) getCoinList(c *gin.Context) ([]*model.HomeCoin, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var coinStats map[string]btcpoolclient.CoinStat
	var coinIncome btcpoolclient.CoinIncomList
	var errS error = nil
	var errI error = nil
	//get mul coin stat
	go func() {
		defer wg.Done()
		coinStats, errS = p.getCoinStats(c)
	}()
	go func() {
		defer wg.Done()
		coinIncome, errI = p.getIncomes(c)
	}()

	wg.Wait()
	if errS == nil && errI == nil {
		// 成功
		return p.formatStatsAndIncomes(coinStats, coinIncome), nil
	} else {
		if errS != nil {
			return nil, errS
		} else {
			return nil, errI
		}
	}
}

// 获取收益
func (p *homeHandler) getIncomes(c *gin.Context) (btcpoolclient.CoinIncomList, error) {

	incomeList := make(btcpoolclient.CoinIncomList, 0)
	if dic, err := btcpoolclient.GetCoinIncome(c); err != nil {
		return incomeList, err
	} else {
		for k, info := range dic {
			info.CoinType = k
			incomeList = append(incomeList, info)
		}
		return incomeList, nil
	}
}

// 全网状态
func (p *homeHandler) getCoinStats(c *gin.Context) (map[string]btcpoolclient.CoinStat, error) {
	return btcpoolclient.GetPoolMultiCoinStats(c)
}

// 2 model.HomeCoin
func (p *homeHandler) formatStatsAndIncomes(stats map[string]btcpoolclient.CoinStat, incomes btcpoolclient.CoinIncomList) []*model.HomeCoin {
	ret := make([]*model.HomeCoin, 0)
	for k, stat := range  stats {
		for _, income := range incomes {
			if strings.ToLower(k) == strings.ToLower(income.CoinType) {
				item := &model.HomeCoin{}
				item.SetData(stat, income, p.lan)
				ret = append(ret, item)
			}
		}
	}
	return ret
}

// module 小工具
func (p *homeHandler) getModules(c *gin.Context) []model.HomeModule {
	observable := model.HomeModule{
		Id: "",
		Icon: "",
		Title: "",
		Url: "",
	}
	counter := model.HomeModule{
		Id: "",
		Icon: "",
		Title: "",
		Url: "",
	}
	minerRank := model.HomeModule{
		Id: "",
		Icon: "",
		Title: "",
		Url: "",
	}
	chainData := model.HomeModule{
		Id: "",
		Icon: "",
		Title: "",
		Url: "",
	}
	return []model.HomeModule{observable, counter, minerRank, chainData}
}

// pool rank
func (p *homeHandler) getPoolRank(c *gin.Context) ([]model.PoolRankInfo, error) {
	platform := c.GetHeader("platform")
	if len(platform) == 0 {
		platform = ""
	}
	param := map[string]string{
		"coins": "btc",
		"from":  platform,
	}
	if dic, err := btcpoolclient.GetPoolRank(c, param); err != nil {
		return []model.PoolRankInfo{}, err
	} else {
		// get btc 's pool rank
		v := dic["btc"]
		ranks := make([]model.PoolRankInfo, 0)
		for i, rankInfo := range v.Realtime.List {
			var hash string
			var unit string
			if len(rankInfo.RealtimeHashrate) == 0 {
				hash = tool.KeepStringNum(rankInfo.EstimateHashrate, 2)
				unit = rankInfo.EstimateHashrateUnit + rankInfo.HashrateSuffix
			} else {
				hash = tool.KeepStringNum(rankInfo.RealtimeHashrate, 2)
				unit = rankInfo.HashrateUnit + rankInfo.HashrateSuffix
			}

			var diff string
			if len(rankInfo.RealtimeHashrate) == 0 {
				diff = "-"
			} else {
				diff = tool.KeepStringNum(rankInfo.RealtimeDiff24hPercent, 2)
			}

			ranks = append(ranks, model.PoolRankInfo{
				Index:                 strconv.Itoa(i + 1),
				Name:                  rankInfo.PoolName,
				Icon:                  rankInfo.IconLink,
				HashratePercent:       rankInfo.Cur2maxPercent,
				Hashrate:              hash,
				HashrateUnit:          unit,
				HashrateChangePercent: diff,
				Lucy7Day:              rankInfo.Lucky3d,
			})
		}

		return ranks, nil
	}
}

//func ExplorerPoolRank(c *gin.Context) {
//
//	var poolRankParams PoolRankParameters
//	if err := c.ShouldBindQuery(&poolRankParams); err != nil {
//		fmt.Printf(">> explorer params2 = %v %v\n", poolRankParams, err)
//		output.ShowErr(c, errs.ApiErrParams)
//		return
//	}
//	p := make(map[string]string)
//	p["coins"] = poolRankParams.Coins
//	p["from"] = poolRankParams.From
//	if res, err := service.PublicService.GetPoolRank(c, p); err!=nil {
//		output.ShowErr(c, err)
//	} else {
//		output.Succ(c, res)
//	}
//}