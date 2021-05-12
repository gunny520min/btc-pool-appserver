package model

import (
	"btc-pool-appserver/application/btcpoolclient"
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

type HomeCoinInfo struct {
	Coin              string  `json:"coin"`
	CurrencyUsd       string  `json:"currencyUsd"`       // 币值
	CurrencyCny       string  `json:"currencyCny"`       // 币值rmb
	PoolHashrate      string  `json:"poolHashrate"`      // 总算力
	HashrateUnit      string  `json:"hashrateUnit"`      // 算力单位
	TotalCount        string  `json:"totalCount"`        // 总币数
	TotalBlocks       string  `json:"totalBlocks"`       // 总区块数
	AllHashrate       string  `json:"allHashrate"`       // 全网算力
	AllHashrateUnit   string  `json:"allHashrateUnit"`   // 全网算力单位
	Diff              float64 `json:"diff"`              // 全网难度
	NextDiff          string  `json:"nextDiff"`          // 下次调整难度
	NextDiffTime      string  `json:"nextDiffTime"`      // 下次调整难度时间
	IncomeCoin        string  `json:"incomeCoin"`        // 币单位收益
	IncomeUnit        string  `json:"incomeUnit"`        // 单位收益单位： th/s
	IncomeCurrencyUsd string  `json:"incomeCurrencyUsd"` // 币值单位收益
	IncomeCurrencyCny string  `json:"incomeCurrencyCny"` // 币值单位收益rmb
	// DiffUnit        string `json:"diffUnit"`        // 全网难度单位
	// NextDiffPer     string  `json:"nextDiffPer"`     // 下次调整难度百分比
}

func (info *HomeCoinInfo) SetData(statInfo btcpoolclient.CoinStat, income btcpoolclient.CoinIncome) {
	info.Coin = statInfo.Coin_type
	info.TotalCount = keepStringNum(statInfo.Rewards_count, 0)
	info.TotalBlocks = keepStringNum(statInfo.Blocks_count, 0)
	info.PoolHashrate = keepStringNum(statInfo.Stats.Shares.Shares_15m, 3)
	info.HashrateUnit = statInfo.Stats.Shares.Shares_unit + statInfo.Coin_suffix

	var hash, unit = calculateHashRate(income.Hashrate, 3)
	info.AllHashrate = hash
	info.AllHashrateUnit = unit + statInfo.Coin_suffix
	info.Diff = income.Diff
	info.NextDiff = income.NextDiff
	info.NextDiffTime = income.DiffAdjustTime
	info.IncomeUnit = income.IncomeHashrateUnit + income.IncomeHashrateUnitSuffix

	if strings.ToLower(statInfo.Coin_type) == "btc" && strings.ToLower(income.PaymentMode) == "fpps" {
		info.IncomeCoin = keepFloatNum(income.IncomeOptimizeCoin, 8)
		info.IncomeCurrencyCny = keepFloatNum(income.IncomeOptimizeCny, 2)
		info.IncomeCurrencyUsd = keepFloatNum(income.IncomeOptimizeUsd, 2)
	} else {
		info.IncomeCoin = keepFloatNum(income.IncomeCoin, 8)
		info.IncomeCurrencyCny = keepFloatNum(income.IncomeCny, 2)
		info.IncomeCurrencyUsd = keepFloatNum(income.IncomeUsd, 2)
	}

}

func keepStringNum(value string, l int32) string {
	if d, e := decimal.NewFromString(value); e != nil {
		return "-"
	} else {
		return d.Round(l).String()
	}
}

func keepFloatNum(value float64, l int32) string {
	return decimal.NewFromFloat(value).Round(l).String()
}

// value 要转化的hashrate，l小数点后位数
func calculateHashRate(value float64, l int32) (string, string) {
	d := decimal.NewFromFloatWithExponent(value, 0) // 取整数部分
	switch len(d.String()) {
	case 0, 1, 2, 3:
		return decimal.NewFromFloat(value).Round(l).String(), ""
	case 4, 5, 6:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(3))).Round(l).String(), "K"
	case 7, 8, 9:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(6))).Round(l).String(), "M"
	case 10, 11, 12:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(9))).Round(l).String(), "G"
	case 13, 14, 15:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(12))).Round(l).String(), "T"
	case 16, 17, 18:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(15))).Round(l).String(), "P"
	case 19, 20, 21:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(18))).Round(l).String(), "E"
	case 22, 23, 24:
		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(21))).Round(l).String(), "Z"
	}
	return "", ""
}
