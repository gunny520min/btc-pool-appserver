package model

import (
	"btc-pool-appserver/application/btcpoolclient"
	"fmt"
	"strings"
)

type HomeCoinInfo struct {
	Coin            string  `json:"coin"`
	CurrencyUsd     string  `json:"currencyUsd"`     // 币值
	CurrencyCny     string  `json:"currencyCny"`     // 币值rmb
	PoolHashrate    string  `json:"poolHashrate"`    // 总算力
	HashrateUnit    string  `json:"hashrateUnit"`    // 算力单位
	TotalCount      string  `json:"totalCount"`      // 总币数
	TotalBlocks     string  `json:"totalBlocks"`     // 总区块数
	AllHashrate     string  `json:"allHashrate"`     // 全网算力
	AllHashrateUnit string  `json:"allHashrateUnit"` // 全网算力单位
	Diff            float64 `json:"diff"`            // 全网难度
	NextDiff        string  `json:"nextDiff"`        // 下次调整难度
	// DiffUnit        string `json:"diffUnit"`        // 全网难度单位
	// NextDiffPer     string  `json:"nextDiffPer"`     // 下次调整难度百分比
	NextDiffTime      string  `json:"nextDiffTime"`      // 下次调整难度时间
	IncomeCoin        float64 `json:"incomeCoin"`        // 币单位收益
	IncomeUnit        string  `json:"incomeUnit"`        // 单位收益单位： th/s
	IncomeCurrencyUsd float64 `json:"incomeCurrencyUsd"` // 币值单位收益
	IncomeCurrencyCny float64 `json:"incomeCurrencyCny"` // 币值单位收益rmb
}

func (info *HomeCoinInfo) SetData(statInfo btcpoolclient.CoinStat, income btcpoolclient.CoinIncome) {
	info.Coin = statInfo.Coin_type
	info.TotalCount = statInfo.Rewards_count
	info.TotalBlocks = statInfo.Blocks_count
	info.PoolHashrate = statInfo.Stats.Shares_15m
	info.HashrateUnit = statInfo.Stats.Shares_unit + statInfo.Coin_suffix

	info.AllHashrate = fmt.Sprintf("%v", income.Hashrate)
	info.AllHashrateUnit = statInfo.Stats.Shares_unit + statInfo.Coin_suffix
	info.Diff = income.Diff
	info.NextDiff = income.NextDiff
	info.NextDiffTime = income.DiffAdjustTime

	if strings.ToLower(statInfo.Coin_type) == "btc" && strings.ToLower(income.PaymentMode) == "fpps" {
		info.IncomeCoin = income.IncomeOptimizeCoin
		info.IncomeCurrencyCny = income.IncomeOptimizeCny
		info.IncomeCurrencyUsd = income.IncomeOptimizeUsd
	} else {
		info.IncomeCoin = income.IncomeCoin
		info.IncomeCurrencyCny = income.IncomeCny
		info.IncomeCurrencyUsd = income.IncomeUsd
	}

}
