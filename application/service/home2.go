package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

type homeHandler struct{
	Lan string
}

var HomeService = &homeHandler{}

// GetBanner Get Banner
func (p *homeHandler) GetBanner(c *gin.Context) []model.Banner {
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

// GetNotice Get Notice
func (p *homeHandler) GetNotice(c *gin.Context) []model.Notice {
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

// GetCoinList coin list
func (p *homeHandler) GetCoinList(c *gin.Context) ([]*model.HomeCoin, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var coinStats map[string]btcpoolclient.CoinStat
	var coinIncome btcpoolclient.CoinIncomList
	var errS error = nil
	var errI error = nil
	//get mul coin stat
	go func() {
		defer wg.Done()
		coinStats, errS = p.GetCoinStats(c)
	}()
	go func() {
		defer wg.Done()
		coinIncome, errI = p.GetIncomes(c)
	}()

	wg.Wait()
	if errS == nil && errI == nil {
		// 成功
		return p.FormatStatsAndIncomes(coinStats, coinIncome), nil
	} else {
		if errS != nil {
			return nil, errS
		} else {
			return nil, errI
		}
	}
}

// GetIncomes 获取收益
func (p *homeHandler) GetIncomes(c *gin.Context) (btcpoolclient.CoinIncomList, error) {

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

// GetCoinStats 全网状态
func (p *homeHandler) GetCoinStats(c *gin.Context) (map[string]btcpoolclient.CoinStat, error) {
	return btcpoolclient.GetPoolMultiCoinStats(c)
}

// FormatStatsAndIncomes 2 model.HomeCoin
func (p *homeHandler) FormatStatsAndIncomes(stats map[string]btcpoolclient.CoinStat, incomes btcpoolclient.CoinIncomList) []*model.HomeCoin {
	ret := make([]*model.HomeCoin, 0)
	for k, stat := range  stats {
		for _, income := range incomes {
			if strings.ToLower(k) == strings.ToLower(income.CoinType) {
				item := &model.HomeCoin{}
				item.SetData(stat, income, p.Lan)
				ret = append(ret, item)
			}
		}
	}
	return ret
}

// GetModules module 小工具
func (p *homeHandler) GetModules(c *gin.Context) []model.HomeModule {
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

type ShareHashrate struct {
	Tickets []model.HashrateData
	Unit string
}
// GetShareHashrate 获取首页算力图表数据
func (p *homeHandler) GetShareHashrate(c *gin.Context, params interface{}) (ShareHashrate, error) {
	var ret ShareHashrate
	if data, err := btcpoolclient.GetPoolShareHashrate(c, params); err != nil {
		return ret, err
	} else {
		ret.Tickets = p.FormatHashrateChartData(data)
		ret.Unit = p.FormatHashrateChartUnit(data)
	}
	return ret, nil
}

func (p *homeHandler) FormatHashrateChartUnit(params btcpoolclient.ShareHashrateData) string {
	return params.Unit
}

func (p *homeHandler) FormatHashrateChartData(params btcpoolclient.ShareHashrateData) []model.HashrateData {
	res := make([]model.HashrateData, 0)
	for _, v := range params.Tickers {
		var item model.HashrateData
		item.Hashrate = v[1]
		item.Timestamp = v[0]
		res = append(res, item)
	}
	return res
}

// GetPoolDataBaseInfo 矿池数据
func (p *homeHandler) GetPoolDataBaseInfo (c *gin.Context, coin string) (model.PoolBaseInfo, error) {
	info := model.PoolBaseInfo{
		MiningGift: "挖BTC免费送NMC、ELA",
		Paymode: "FPPS",
		PayLimit: "0.005BTC",
		PayType: "达到起付额后自动支付",
		PayTime: "每日 08:00-16:00(UTC)支付昨日收益",
		MiningTutorial: "https://btc.com",
		Address: []string{"us.ss.btc.com:1800", "us.ss.btc.com:443", "us.ss.btc.com:25"},
		AccountName: "youdao1.001",
	}
	return info, nil
}

