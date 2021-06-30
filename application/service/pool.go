package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/library/tool"
	"btc-pool-appserver/application/model"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type poolHandler struct{}

var PoolService = &poolHandler{}

// GetDashboardSubaccounts 用户面包，用户的子账户数据
func (p *poolHandler) GetDashboardSubaccounts(c *gin.Context, puid string) (*clientModel.SubAccountAlgorithmList, *clientModel.SubAccountCoinEntity, error) {
	var subaccounts *clientModel.SubAccountAlgorithmList = nil
	var currentCoinEntity *clientModel.SubAccountCoinEntity = nil

	var algorithmsParams = map[string]interface{}{}
	algorithmsParams["puid"] = puid
	algorithmsParams["is_hidden"] = 0
	algorithmsParams["is_guardian"] = 0
	algorithmsParams["order_by"] = "puid"
	if subAccountAlgorithms, e := btcpoolclient.GetSubAccountAlgorithms(c, algorithmsParams); e != nil {
		return nil, nil, e
	} else {
		if len(subAccountAlgorithms.SubAccounts)==0 {
			return &subAccountAlgorithms, nil, nil
		}
		// 统一机枪池 与 普通账户
		for _, subaccount := range subAccountAlgorithms.SubAccounts {
			for _, algorithm := range subaccount.Algorithms {
				j := 0
				var l = make([]clientModel.SubAccountCoinEntity, 0)
				for _, coinAccount := range algorithm.CoinAccounts {
					if algorithm.IsSmart() == coinAccount.IsSmart() {
						l = append(l, coinAccount)
						//subaccount.Algorithms[j] = algorithm
						j++
					}
				}
				a := &algorithm
				a.CoinAccounts = l
			}
		}
		var cce clientModel.SubAccountCoinEntity
		// 取puid相同的coinAccount
		if len(puid) > 0 {
			for _, subaccount := range subAccountAlgorithms.SubAccounts {
				for _, algorithm := range subaccount.Algorithms {
					for _, coinAccount := range algorithm.CoinAccounts {
						if puid == coinAccount.Puid {
							cce = coinAccount
						}
					}
				}
			}
		}
		currentCoinEntity = &cce
		// 如果没有取到，默认使用第一个作为账户默认coinAccount
		if currentCoinEntity == nil {
			currentCoinEntity = &subAccountAlgorithms.SubAccounts[0].Algorithms[0].CoinAccounts[0]
		}
		subaccounts = &subAccountAlgorithms
		return subaccounts, currentCoinEntity, nil
	}
}

// GetDashboardIncome 获取用户面板收益数据
func (p *poolHandler) GetDashboardIncome(c *gin.Context, puid string) (model.Income, error) {
	var incomeParams = map[string]interface{}{}
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

// 计算收益数值和单位
func countIncome(earn string) (string, string) {
	eNum := decimal.RequireFromString(earn)
	eStr := decimal.RequireFromString(earn).Truncate(0).String()
	var intCount = len(eStr)
	unit := ""
	if len(eStr) > 10 {
		eNum = eNum.Div(decimal.NewFromFloat(math.Pow10(14)))
		intCount = len(eNum.Truncate(0).String())
		unit = "M"
	} else if len(eStr) > 7 {
		eNum = eNum.Div(decimal.NewFromFloat(math.Pow10(11)))
		intCount = len(eNum.Truncate(0).String())
		unit = "K"
	}
	dCount := 7 - intCount
	if dCount < 0 {
		dCount = 0
	}

	return tool.KeepStringNum(eNum.Truncate(int32(dCount)).String(), int32(dCount)), unit
}

// GetDashboardWorkerStates 用户面板矿工状态数据
func (p *poolHandler) GetDashboardWorkerStates(c *gin.Context, puid string, coinType string) (model.WorkerStatus, error) {
	var incomeParams = map[string]interface{}{}
	incomeParams["puid"] = puid
	var workerStats model.WorkerStatus
	if ws, err := btcpoolclient.GetWorkerStats(c, incomeParams); err != nil {
		return workerStats, err
	} else {
		workerStats.TotalHashrate.Value = ws.Shares15m
		workerStats.TotalHashrate.Unit = ws.SharesUnit + model.GetCoinSuffixByCoinType(coinType)
		workerStats.WorkerInactive = ws.WorkersInactive
		workerStats.WorkerActive = ws.WorkersActive

		return workerStats, nil
	}
}

// GetDashboardWorkerGroup 用户面板群组数据
func (p *poolHandler) GetDashboardWorkerGroup(c *gin.Context, puid string, coinType string) ([]model.WorkerGroup, error) {
	var groupParams = struct {
		Puid     string `json:"puid"`
		Page     string `json:"page"`
		PageSize string `json:"page_size"`
	}{
		Puid:     puid,
		Page:     "1",
		PageSize: "5",
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
			mWG.Hashrate.Unit = g.SharesUnit + model.GetCoinSuffixByCoinType(coinType)

			workerGroups = append(workerGroups, mWG)
		}
		return workerGroups, nil
	}
}

// GetDashboardMiningAddress 用户面板挖坑地址数据
func (p *poolHandler) GetDashboardMiningAddress(c *gin.Context, puid string, coinType string) (model.MiningAddress, error) {
	var addrParams = struct {
		Puid string `json:"puid"`
	}{
		Puid: puid,
	}

	var miningAddress model.MiningAddress
	if sba, err := btcpoolclient.GetSubAccountInfo(c, addrParams); err != nil {
		return miningAddress, err
	} else {
		var list = sba.StratumUrlConf
		miningAddress.Title = "挖坑地址" // 国际化
		miningAddress.Desc = "矿机名设置参考：" + sba.Name + ".001"
		miningAddress.Tips = "注：矿工名须由数字或小写字母组成，密码可任意设置"
		miningAddress.Address = make([]model.MiningAddressDetail, 0)
		for index, g := range list {
			tip := ""
			url := strings.Split(g, " ")
			if index == 0 && len(url) == 2 && len(url[1]) > 0 {
				if strings.ToUpper(coinType) == "BEAM" {
					tip = url[0] + " 支持普通显卡矿机\n" + url[1] + " 为Nicehash特别优化"
				}
				if strings.ToUpper(coinType) == "DCR" {
					tip = url[0] + " 已知支持蚂蚁矿机DR3，Claymore，gominer\n" + url[1] + " 已知支持Nicehash，芯动科技等矿机"
				}
			}

			addD := model.MiningAddressDetail{
				Addr: g,
				Tips: tip,
			}

			miningAddress.Address = append(miningAddress.Address, addD)
		}
		return miningAddress, nil
	}
}

// GetDashboardWorkerShareHistory 获取用户面板算力图表数据
func (p *poolHandler) GetDashboardWorkerShareHistory(c *gin.Context, puid string, dimension string) (model.WorkerShareHistory, error) {
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
	var params = struct {
		Puid      string `json:"puid"`
		StartTs   string `json:"start_ts"`
		Dimension string `json:"dimension"`
		Count     int    `json:"count"`
		RealPoint bool   `json:"real_point"`
	}{
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
