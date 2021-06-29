package model

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/tool"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

type Banner struct {
	Id     string `json:"id"`
	ImgUrl string `json:"imgUrl"`
	Url    string `json:"url"`
}

type Notice struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

// 首页小工具
type HomeModule struct {
	Id    string `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type PoolRankInfo struct {
	Index     string                 `json:"index"`
	Icon string                      `json:"icon"`
	Name string                      `json:"name"`
	HashratePercent string           `json:"hashratePercent"`
	Hashrate string                  `json:"hashrate"`
	HashrateUnit string              `json:"hashrateUnit"`
	HashrateChangePercent     string `json:"hashrateChangePercent"`
	Lucy7Day string                  `json:"lucy7Day"`
}


type HomeCoin struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Algorithm      string `json:"algorithm"`
	DayIncome      string `json:"dayIncome"`
	DayIncomeUnit  string `json:"dayIncomeUnit"`
	Price          string `json:"price"`
	PoolHashrate   string `json:"poolHashrate"`
	AllnetHashrate string `json:"allnetHashrate"`
	HashrateUnit   string `json:"hashrateUnit"`
	PayLimit       string `json:"payLimit"`
	PayMode        string `json:"payMode"`
	Difficult      string `json:"difficult"`
	NextDiff       string `json:"nextDiff"`
	NextAdjustTime string `json:"nextAdjustTime"`
}

func (info *HomeCoin) SetData(statInfo btcpoolclient.CoinStat, income btcpoolclient.CoinIncome, lan string) {
	info.Id = statInfo.CoinType
	info.Name = statInfo.IndexCoin.Text
	info.Algorithm = statInfo.IndexCoin.Algorithm
	isEn := strings.ToLower(lan) == "en_us"
	isFpps := strings.ToLower(income.PaymentMode) == "fpps"
	if isEn {
		if isFpps {
			info.DayIncome = "$" + tool.FormatFloat(income.IncomeOptimizeUsd, 2)
		} else {
			info.DayIncome = "$" + tool.FormatFloat(income.IncomeUsd, 2)
		}
		info.Price = "$" + fmt.Sprintf("%.2f", income.IncomeUsd/income.IncomeCoin)
	} else {
		if isFpps {
			info.DayIncome = "￥" + tool.FormatFloat(income.IncomeOptimizeCny, 2)
		} else {
			info.DayIncome = "￥" + tool.FormatFloat(income.IncomeCny, 2)
		}
		info.Price = "￥" + fmt.Sprintf("%.2f", income.IncomeCny/income.IncomeCoin)
	}
	info.DayIncomeUnit = income.IncomeHashrateUnit + income.IncomeHashrateUnitSuffix
	info.PoolHashrate = tool.KeepStringNum(statInfo.Stats.Shares.Shares15m, 3)
	var hash, unit = calculateHashRate(income.Hashrate, 3)
	info.AllnetHashrate = hash
	info.HashrateUnit = unit
	info.PayLimit = statInfo.CoinPayLimit
	info.PayMode = statInfo.CoinPayMode
	var diff, diffUnit = calculateHashRate(income.Diff, 2)
	info.Difficult = diff + diffUnit
	if strings.Contains("btc,dcr,ltc,BTC,DCR,LTC", strings.ToLower(statInfo.CoinType)) {
		if nDiff, err := strconv.ParseFloat(income.NextDiff, 64); err == nil {
			change := ((nDiff - income.Diff) / income.Diff) * 100
			var nDiffStr, nDiffUnit = calculateHashRate(nDiff, 3)
			info.NextDiff = fmt.Sprintf("%s%s(%0.3f%%)", nDiffStr, nDiffUnit, change)
			info.NextAdjustTime = income.DiffAdjustTime
		} else {
			info.NextDiff = "-"
			info.NextAdjustTime = "-"
		}
	} else {
		info.NextDiff = "-"       //income.NextDiff
		info.NextAdjustTime = "-" //income.DiffAdjustTime
	}
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
