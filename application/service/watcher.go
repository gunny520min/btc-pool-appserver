package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type watcherHandler struct{}

var WatcherService = &watcherHandler{}

func (p *watcherHandler) AddOtherWatcher(c *gin.Context, params interface{}) (btcpoolclient.Watcher, error) {

	var w btcpoolclient.Watcher
	//check watcher
	if d, err := btcpoolclient.CheckWatcher(c, params); err != nil {
		return w, err
	} else if len(d) > 0 {
		w = d[0]
	} else {
		return w, nil
	}
	// get watcher info
	pms := map[string]string{}
	pms["puid"] = w.Puid
	pms["access_key"] = w.Token
	if d, err := btcpoolclient.WatcherRegionInfo(c, pms); err != nil {
		return w, err
	} else {
		w.RegionNameCN = d.Text["zh-cn"]
		w.RegionNameEN = d.Text["en"]
	}
	return w, nil
}

func (p *watcherHandler) GetWatcherHashrate(c *gin.Context, params interface{}) ([]btcpoolclient.Watcher, error) {

	if d, err := btcpoolclient.CheckWatcher(c, params); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

// GetDashboardAccounts 观察者面板，权限
func (p *watcherHandler) GetWatcherAuthority(c *gin.Context, accessKey string, puid string) (*clientModel.WatcherAuthority, error) {
	var params = map[string]interface{}{
		"access_key": accessKey,
		"puid":       puid,
	}
	if d, err := btcpoolclient.WatcherAuthority(c, params); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

// GetWatcherDashboardIncome 获取观察者面板收益数据
func (p *watcherHandler) GetWatcherDashboardIncome(c *gin.Context, accessKey string, puid string) (model.Income, error) {
	var incomeParams = map[string]interface{}{}
	incomeParams["access_key"] = accessKey
	incomeParams["puid"] = puid
	incomeParams["is_decimal"] = 1
	var income model.Income
	var normalIncome model.NormalIncome
	var smartIncome model.SmartIncome
	if earnState, err := btcpoolclient.GetEarnstats(c, incomeParams); err != nil {
		return income, err
	} else {
		if earnState.IsSmart() {
			smartIncome.IncomeToday.Value, smartIncome.IncomeToday.Unit = countIncome(earnState.EarningsToday)
			smartIncome.IncomeToday.Coin = earnState.GetCoin()
			smartIncome.IsOtc = earnState.EarningsYesterdayIsOtc
			smartIncome.IncomeYesterday.Btc.Value, smartIncome.IncomeYesterday.Btc.Unit = countIncome(earnState.EarningsYesterdayCoins.Btc)
			smartIncome.IncomeYesterday.Btc.Coin = "BTC"
			smartIncome.IncomeYesterday.Bch.Value, smartIncome.IncomeYesterday.Bch.Unit = countIncome(earnState.EarningsYesterdayCoins.Bch)
			smartIncome.IncomeYesterday.Bch.Coin = "BCH"
			smartIncome.IncomeYesterday.Bsv.Value, smartIncome.IncomeYesterday.Bsv.Unit = countIncome(earnState.EarningsYesterdayCoins.Bsv)
			smartIncome.IncomeYesterday.Bsv.Coin = "BSV"
		} else {
			normalIncome.IncomeUnpaid.Value, normalIncome.IncomeUnpaid.Unit = countIncome(earnState.Unpaid)
			normalIncome.IncomeUnpaid.Coin = earnState.GetCoin()
			normalIncome.IncomePaid.Value, normalIncome.IncomePaid.Unit = countIncome(earnState.TotalPaid)
			normalIncome.IncomePaid.Coin = earnState.GetCoin()
			normalIncome.IncomeToday.Value, normalIncome.IncomeToday.Unit = countIncome(earnState.EarningsToday)
			normalIncome.IncomeToday.Coin = earnState.GetCoin()
			normalIncome.IncomeYesterday.Value, normalIncome.IncomeYesterday.Unit = countIncome(earnState.EarningsYesterday)
			normalIncome.IncomeYesterday.Coin = earnState.GetCoin()
		}
		income.HasIncome = earnState.EarningsBefore
		income.SmartIncome = smartIncome
		income.Income = normalIncome
		income.CoinType = earnState.CoinType
		return income, nil
	}
}

// GetWatcherDashboardWorkerStates  观察者面板矿工状态数据
func (p *watcherHandler) GetWatcherDashboardWorkerStates(c *gin.Context, accessKey string, puid string) (model.WorkerStatus, error) {
	var params = map[string]interface{}{}
	params["access_key"] = accessKey
	params["puid"] = puid
	var workerStats model.WorkerStatus
	if ws, err := btcpoolclient.GetWorkerStats(c, params); err != nil {
		return workerStats, err
	} else {
		workerStats.TotalHashrate.Value = ws.Shares15m
		workerStats.TotalHashrate.Unit = ws.SharesUnit // TODO add CoinType
		workerStats.WorkerInactive = ws.WorkersInactive
		workerStats.WorkerActive = ws.WorkersActive

		return workerStats, nil
	}
}

// GetWatcherDashboardWorkerGroup 观察者面板群组数据
func (p *watcherHandler) GetWatcherDashboardWorkerGroup(c *gin.Context, accessKey string, puid string) ([]model.WorkerGroup, error) {
	var groupParams = struct {
		AccessKey string `json:"access_key"`
		Puid      string `json:"puid"`
		Page      string `json:"page"`
		PageSize  string `json:"page_size"`
	}{
		AccessKey: accessKey,
		Puid:      puid,
		Page:      "1",
		PageSize:  "5",
	}

	var workerGroups []model.WorkerGroup
	if list, err := btcpoolclient.WorkerGroups(c, groupParams); err != nil {
		return workerGroups, err
	} else {
		for _, g := range list {
			mWG := model.WorkerGroup{
				Gid:          g.Gid,
				Name:         g.Name,
				WorkerActive: g.WorkersActive,
				WorkerTotal:  g.WorkersTotal,
			}
			mWG.Hashrate.Value = g.Shares15m
			mWG.Hashrate.Unit = g.SharesUnit

			workerGroups = append(workerGroups, mWG)
		}
		return workerGroups, nil
	}
}

// GetWatcherDashboardMiningAddress 观察者面板挖坑地址数据
func (p *watcherHandler) GetWatcherDashboardMiningAddress(c *gin.Context, accessKey string, puid string) (model.MiningAddress, error) {
	var addrParams = struct {
		AccessKey string `json:"access_key"`
		Puid      string `json:"puid"`
	}{
		AccessKey: accessKey,
		Puid:      puid,
	}
	var miningAddress model.MiningAddress
	if sba, err := btcpoolclient.GetSubAccountInfo(c, addrParams); err != nil {
		return miningAddress, err
	} else {
		var list = sba.StratumUrlConf
		miningAddress.Title = "挖坑地址" // TODO 国际化
		miningAddress.Desc = "矿机名设置参考：" + sba.Name + ".001"
		miningAddress.Tips = "注：矿工名须由数字或小写字母组成，密码可任意设置"
		miningAddress.Address = make([]model.MiningAddressDetail, 0)
		for _, g := range list {
			addD := model.MiningAddressDetail{
				Addr: g,
				Tips: "",
			}

			miningAddress.Address = append(miningAddress.Address, addD)
		}
		return miningAddress, nil
	}
}

// GetWatcherDashboardWorkerShareHistory 获取用户面板算力图表数据
func (p *watcherHandler) GetWatcherDashboardWorkerShareHistory(c *gin.Context, accessKey string, puid string, dimension string) (model.WorkerShareHistory, error) {
	var startTs = time.Now().Unix()
	var count = 30
	if dimension == "1d" {
		// 以天为单位，统计30天
		startTs = startTs - 60*60*24*30
		count = 30
	} else {
		// 以小时为单位，统计72小时
		startTs = startTs - 60*60*24*3
		count = 72
	}
	params := struct {
		AccessKey string `json:"access_key"`
		Puid      string `json:"puid"`
		StartTs   string `json:"start_ts"`
		Dimension string `json:"dimension"`
		Count     int    `json:"count"`
		RealPoint bool   `json:"real_point"`
	}{
		AccessKey: accessKey,
		Puid:      puid,
		StartTs:   strconv.FormatInt(startTs, 10),
		Dimension: dimension,
		Count:     count,
		RealPoint: true,
	}

	var workerShareHistory model.WorkerShareHistory
	if wsh, err := btcpoolclient.GetWorkerShareHistory(c, params); err != nil {
		return workerShareHistory, err
	} else {
		workerShareHistory = model.WorkerShareHistory{
			Unit: wsh.SharesUnit,
			List: make([]model.WorkerShareHistoryEntity, len(wsh.Tickers)),
		}
		for index, entity := range wsh.Tickers {
			var workerShareHistoryEntity = model.WorkerShareHistoryEntity{
				Timestamp: entity[0],
				Hashrate:  entity[1],
				Reject:    entity[2],
			}
			workerShareHistory.List[index] = workerShareHistoryEntity
		}

		return workerShareHistory, nil
	}
}

// GetWatcherDashboardIncome 获取观察者用户面板收益数据
func (p *poolHandler) GetWatcherDashboardIncome(c *gin.Context, accessKey string, puid string) (model.Income, error) {
	var incomeParams = map[string]interface{}{}
	incomeParams["access_key"] = accessKey
	incomeParams["puid"] = puid
	incomeParams["is_decimal"] = 1
	var income model.Income
	var normalIncome model.NormalIncome
	var smartIncome model.SmartIncome
	if earnState, err := btcpoolclient.GetEarnstats(c, incomeParams); err != nil {
		return income, err
	} else {
		if earnState.IsSmart() {
			smartIncome.IncomeToday.Value, smartIncome.IncomeToday.Unit = countIncome(earnState.EarningsToday)
			smartIncome.IncomeToday.Coin = earnState.GetCoin()
			smartIncome.IsOtc = earnState.EarningsYesterdayIsOtc
			smartIncome.IncomeYesterday.Btc.Value, smartIncome.IncomeYesterday.Btc.Unit = countIncome(earnState.EarningsYesterdayCoins.Btc)
			smartIncome.IncomeYesterday.Btc.Coin = "BTC"
			smartIncome.IncomeYesterday.Bch.Value, smartIncome.IncomeYesterday.Bch.Unit = countIncome(earnState.EarningsYesterdayCoins.Bch)
			smartIncome.IncomeYesterday.Bch.Coin = "BCH"
			smartIncome.IncomeYesterday.Bsv.Value, smartIncome.IncomeYesterday.Bsv.Unit = countIncome(earnState.EarningsYesterdayCoins.Bsv)
			smartIncome.IncomeYesterday.Bsv.Coin = "BSV"
		} else {
			normalIncome.IncomeUnpaid.Value, normalIncome.IncomeUnpaid.Unit = countIncome(earnState.Unpaid)
			normalIncome.IncomeUnpaid.Coin = earnState.GetCoin()
			normalIncome.IncomePaid.Value, normalIncome.IncomePaid.Unit = countIncome(earnState.TotalPaid)
			normalIncome.IncomePaid.Coin = earnState.GetCoin()
			normalIncome.IncomeToday.Value, normalIncome.IncomeToday.Unit = countIncome(earnState.EarningsToday)
			normalIncome.IncomeToday.Coin = earnState.GetCoin()
			normalIncome.IncomeYesterday.Value, normalIncome.IncomeYesterday.Unit = countIncome(earnState.EarningsYesterday)
			normalIncome.IncomeYesterday.Coin = earnState.GetCoin()
		}
		income.HasIncome = earnState.EarningsBefore
		income.SmartIncome = smartIncome
		income.Income = normalIncome
		return income, nil
	}
}
