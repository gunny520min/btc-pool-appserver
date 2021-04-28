package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"fmt"
	"github.com/shopspring/decimal"
	"math"

	"github.com/gin-gonic/gin"
)

type poolHandler struct{}

var PoolService = &poolHandler{}

// 获取首页算力图表数据
func (p *poolHandler) GetShareHashrate(c *gin.Context, params interface{}) btcpoolclient.ShareHashrateData {
	var ret btcpoolclient.ShareHashrateData
	if list, err := btcpoolclient.GetPoolShareHashrate(c, params); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeNu)
	} else {
		ret = list
	}
	return ret
}

func (p *poolHandler) FormatHashrateChartUnit(params btcpoolclient.ShareHashrateData) string {
	return params.Unit
}

func (p *poolHandler) FormatHashrateChartData(params btcpoolclient.ShareHashrateData) []model.HashrateData {
	res := make([]model.HashrateData, 0)
	for _, v := range params.Tickers {
		var item model.HashrateData
		item.Hashrate = v[1]
		item.Timestamp = v[0]
		res = append(res, item)
	}
	fmt.Printf("FormatHashrateChartData count= %v, %v", len(res), len(params.Tickers))
	return res
}

// 用户面包，用户的子账户数据
func (p *poolHandler) GetDashboardSubaccounts(c *gin.Context, puid string) (*clientModel.SubAccountAlgorithmList, *clientModel.SubAccountCoinEntity, error) {
	var subaccounts *clientModel.SubAccountAlgorithmList = nil
	var currentCoinEntity *clientModel.SubAccountCoinEntity = nil

	var algorithmsParams = map[string]interface{}{}
	algorithmsParams["puid"] = puid
	algorithmsParams["is_hidden"] = 0
	algorithmsParams["is_guardian"] = 0
	algorithmsParams["order_by"] = "puid"
	if subAccountAlgorithms, e := btcpoolclient.GetSubAccountAlgorithms(c, algorithmsParams); e != nil {
		output.ShowErr(c, e)
		return nil, nil, e
	} else {
		// 统一机枪池 与 普通账户
		for _, subaccount := range subAccountAlgorithms.SubAccounts {
			for _, algorithm := range subaccount.Algorithms {
				j := 0
				for _, coinAccount := range algorithm.CoinAccounts {
					if algorithm.IsSmart() == coinAccount.IsSmart() {
						subaccount.Algorithms[j] = algorithm
						j++
					}
				}
				subaccount.Algorithms = subaccount.Algorithms[:j]
			}
		}
		// 取puid相同的coinAccount
		if len(puid) > 0 {
			for _, subaccount := range subAccountAlgorithms.SubAccounts {
				for _, algorithm := range subaccount.Algorithms {
					for _, coinAccount := range algorithm.CoinAccounts {
						if puid == coinAccount.Puid {
							currentCoinEntity = &coinAccount
						}
					}
				}
			}
		}
		// 如果没有取到，默认使用第一个作为账户默认coinAccount
		if currentCoinEntity == nil {
			currentCoinEntity = &subAccountAlgorithms.SubAccounts[0].Algorithms[0].CoinAccounts[0]
		}
		subaccounts = &subAccountAlgorithms
		return subaccounts, currentCoinEntity, nil
	}
}

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
		return income, nil
	}
}

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
	return eNum.Truncate(int32(dCount)).String(), unit
}


