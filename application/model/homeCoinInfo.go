package model

import (
	"btc-pool-appserver/application/btcpoolclient"

	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type HomeCoinInfo struct {
	Coin              string `json:"coin"`
	CurrencyUsd       string `json:"currency_usd"`        // 币值
	CurrencyCny       string `json:"currency_cny"`        // 币值rmb
	PoolHashrate      string `json:"pool_hashrate"`       // 总算力
	HashrateUnit      string `json:"hashrate_unit"`       // 算力单位
	TotalCount        string `json:"total_count"`         // 总币数
	TotalBlocks       string `json:"total_blocks"`        // 总区块数
	AllHashrate       string `json:"all_hashrate"`        // 全网算力
	AllHashrateUnit   string `json:"all_hashrate_unit"`   // 全网算力单位
	Diff              string `json:"diff"`                // 全网难度
	DiffUnit          string `json:"diff_unit"`           // 难度单位
	NextDiff          string `json:"next_iff"`            // 下次调整难度
	NextDiffUnit      string `json:"next_diff_unit"`      // 难度单位
	NextDiffChange    string `json:"next_diff_change"`    // 下次调整难度百分比
	NextDiffTime      string `json:"next_diff_time"`      // 下次调整难度时间
	IncomeCoin        string `json:"income_coin"`         // 币单位收益
	IncomeUnit        string `json:"income_unit"`         // 单位收益单位： th/s
	IncomeCurrencyUsd string `json:"income_currency_usd"` // 币值单位收益
	IncomeCurrencyCny string `json:"income_currency_cny"` // 币值单位收益rmb

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

	info.IncomeUnit = income.IncomeHashrateUnit + income.IncomeHashrateUnitSuffix

	var diff, diffUnit = calculateHashRate(income.Diff, 2)
	info.Diff = diff
	info.DiffUnit = diffUnit

	if strings.ToLower(statInfo.Coin_type) == "btc" {
		if nDiff, err := strconv.ParseFloat(income.NextDiff, 64); err == nil {
			change := ((nDiff - income.Diff) / income.Diff) * 100
			var nDiffStr, nDiffUnit = calculateHashRate(nDiff, 2)
			info.NextDiff = nDiffStr
			info.NextDiffUnit = nDiffUnit
			info.NextDiffChange = fmt.Sprintf("%0.3f%%", change)
		}
		info.NextDiffTime = income.DiffAdjustTime
	} else {
		info.NextDiff = income.NextDiff
		info.NextDiffTime = income.DiffAdjustTime
		info.NextDiffChange = ""
	}

	if strings.ToLower(statInfo.Coin_type) == "btc" && strings.ToLower(income.PaymentMode) == "fpps" {
		info.IncomeCoin = keepFloatNum(income.IncomeOptimizeCoin, 8)
		info.IncomeCurrencyCny = keepFloatNum(income.IncomeOptimizeCny, 2)
		info.IncomeCurrencyUsd = keepFloatNum(income.IncomeOptimizeUsd, 2)
	} else {
		info.IncomeCoin = keepFloatNum(income.IncomeCoin, 8)
		info.IncomeCurrencyCny = keepFloatNum(income.IncomeCny, 2)
		info.IncomeCurrencyUsd = keepFloatNum(income.IncomeUsd, 2)
	}
	info.CurrencyCny = fmt.Sprintf("%.2f", income.IncomeCny/income.IncomeCoin)
	info.CurrencyUsd = fmt.Sprintf("%.2f", income.IncomeUsd/income.IncomeCoin)
}

func keepStringNum(value string, l int32) string {
	s := fmt.Sprintf("%%.%df", l)
	if v, err := strconv.ParseFloat(value, 64); err != nil {
		return "-"
	} else {
		return fmt.Sprintf(s, v)
	}
	//if d, e := decimal.NewFromString(value); e != nil {
	//	return "-"
	//} else {
	//	return d.Round(l).String()
	//}
}

func keepFloatNum(value float64, l int32) string {
	s := fmt.Sprintf("%%.%df", l)
	return fmt.Sprintf(s, value)

	//return decimal.NewFromFloat(value).Round(l).String()
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
