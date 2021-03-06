package model
//
//import (
//	"btc-pool-appserver/application/btcpoolclient"
//	"btc-pool-appserver/application/library/tool"
//
//	"fmt"
//	"math"
//	"strconv"
//	"strings"
//
//	"github.com/shopspring/decimal"
//)
//
//type HomeCoinInfo struct {
//	Coin              string `json:"coin"`
//	CurrencyCny       string `json:"currencyCny"`
//	CurrencyUsd       string `json:"currencyUsd"`
//	PoolHashrate      string `json:"poolHashrate"`
//	HashrateUnit      string `json:"hashrateUnit"`
//	TotalCount        string `json:"totalCount"`
//	TotalBlocks       string `json:"totalBlocks"`
//	AllHashrate       string `json:"allHashrate"`
//	AllHashrateUnit   string `json:"allHashrateUnit"`
//	Diff              string `json:"diff"`
//	DiffUnit          string `json:"diffUnit"`
//	NextDiff          string `json:"nextDiff"`
//	NextDiffUnit      string `json:"nextDiffUnit"`
//	NextDiffChange    string `json:"nextDiffChange"`
//	NextDiffTime      string `json:"nextDiffTime"`
//	IncomeCoin        string `json:"incomeCoin"`
//	IncomeUnit        string `json:"incomeUnit"`
//	IncomeCurrencyUsd string `json:"incomeCurrencyUsd"`
//	IncomeCurrencyCny string `json:"incomeCurrencyCny"`
//}
//
//func (info *HomeCoinInfo) SetData(statInfo btcpoolclient.CoinStat, income btcpoolclient.CoinIncome) {
//	info.Coin = statInfo.CoinType
//
//	info.TotalCount = tool.KeepStringNum(statInfo.RewardsCount, 0)
//	info.TotalBlocks = tool.KeepStringNum(statInfo.BlocksCount, 0)
//	info.PoolHashrate = tool.KeepStringNum(statInfo.Stats.Shares.Shares15m, 3)
//
//	info.HashrateUnit = statInfo.Stats.Shares.SharesUnit + statInfo.CoinSuffix
//
//	var hash, unit = calculateHashRate(income.Hashrate, 3)
//	info.AllHashrate = hash
//	info.AllHashrateUnit = unit + statInfo.CoinSuffix
//
//	info.IncomeUnit = income.IncomeHashrateUnit + income.IncomeHashrateUnitSuffix
//
//	var diff, diffUnit = calculateHashRate(income.Diff, 2)
//	info.Diff = diff
//	info.DiffUnit = diffUnit
//
//	if strings.Contains("btc,dcr,ltc,BTC,DCR,LTC", statInfo.CoinType) {
//		if nDiff, err := strconv.ParseFloat(income.NextDiff, 64); err == nil {
//			change := ((nDiff - income.Diff) / income.Diff) * 100
//			var nDiffStr, nDiffUnit = calculateHashRate(nDiff, 3)
//			info.NextDiff = nDiffStr
//			info.NextDiffUnit = nDiffUnit
//			info.NextDiffChange = fmt.Sprintf("%0.3f%%", change)
//			info.NextDiffTime = income.DiffAdjustTime
//		} else {
//			info.NextDiff = "-"
//			info.NextDiffTime = "-"
//			info.NextDiffChange = ""
//		}
//	} else {
//		info.NextDiff = "-"//income.NextDiff
//		info.NextDiffTime = "-"//income.DiffAdjustTime
//		info.NextDiffChange = ""
//	}
//
//	if strings.ToLower(statInfo.CoinType) == "btc" && strings.ToLower(income.PaymentMode) == "fpps" {
//		info.IncomeCoin = tool.KeepFloatNum(income.IncomeOptimizeCoin, 8)
//		info.IncomeCurrencyCny = tool.KeepFloatNum(income.IncomeOptimizeCny, 2)
//		info.IncomeCurrencyUsd = tool.KeepFloatNum(income.IncomeOptimizeUsd, 2)
//	} else {
//		if strings.ToLower(statInfo.CoinType) == "ubtc" {
//			info.IncomeCoin = "-"
//			info.IncomeCurrencyCny = tool.KeepFloatNum(income.IncomeCny, 2)
//			info.IncomeCurrencyUsd = tool.KeepFloatNum(income.IncomeUsd, 2)
//		}else {
//			info.IncomeCoin = tool.KeepFloatNum(income.IncomeCoin, 8)
//			info.IncomeCurrencyCny = tool.KeepFloatNum(income.IncomeCny, 2)
//			info.IncomeCurrencyUsd = tool.KeepFloatNum(income.IncomeUsd, 2)
//		}
//	}
//	info.CurrencyCny = fmt.Sprintf("%.2f", income.IncomeCny/income.IncomeCoin)
//	info.CurrencyUsd = fmt.Sprintf("%.2f", income.IncomeUsd/income.IncomeCoin)
//}
//
//// value ????????????hashrate???l??????????????????
//func calculateHashRate(value float64, l int32) (string, string) {
//	d := decimal.NewFromFloatWithExponent(value, 0) // ???????????????
//	switch len(d.String()) {
//	case 0, 1, 2, 3:
//		return decimal.NewFromFloat(value).Round(l).String(), ""
//	case 4, 5, 6:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(3))).Round(l).String(), "K"
//	case 7, 8, 9:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(6))).Round(l).String(), "M"
//	case 10, 11, 12:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(9))).Round(l).String(), "G"
//	case 13, 14, 15:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(12))).Round(l).String(), "T"
//	case 16, 17, 18:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(15))).Round(l).String(), "P"
//	case 19, 20, 21:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(18))).Round(l).String(), "E"
//	case 22, 23, 24:
//		return decimal.NewFromFloat(value).Div(decimal.NewFromFloat(math.Pow10(21))).Round(l).String(), "Z"
//	}
//	return "", ""
//}
