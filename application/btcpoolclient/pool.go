package btcpoolclient

import (
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type ShareHashrateData struct {
	Unit    string        `json:"unit"`
	Tickers []([]float64) `json:"tickers"`
}

// 获取首页算力图表数据
func GetPoolShareHashrate(c *gin.Context, params interface{}) (ShareHashrateData, error) {
	var dest = struct {
		BtcpoolRescomm
		Data ShareHashrateData `json:"data"`
	}{}

	_, err := doRequest(c, "pool.hashrateHistory", params, &dest)
	if err != nil {
		var res ShareHashrateData
		return res, fmt.Errorf("error getting banner list: %v", err)
	}
	return dest.Data, nil
}

type EarnStats struct {
	TotalPaid                    string             `json:"total_paid"`
	PendingPayouts               string             `json:"pending_payouts"`
	EarningsYesterdayPaymentTime string             `json:"earnings_yesterday_payment_time"`
	EarningsYesterday            string             `json:"earnings_yesterday"`
	EarningsYesterdayIsOtc       bool               `json:"earnings_yesterday_is_otc"`
	LastPaymentTime              string             `json:"last_payment_time"`
	EarningsToday                string             `json:"earnings_today"`
	Unpaid                       string             `json:"unpaid"`
	RelativePpsRate              string             `json:"relative_pps_rate"`
	Amount100t                   string             `json:"amount_100t"`
	AmountStandardEarn           string             `json:"amount_standard_earn"`
	AmountStandardUnit           string             `json:"amount_standard_unit"`
	HashrateYesterday            HashrateUnitEntity `json:"hashrate_yesterday"`
	Shares1d                     HashrateUnitEntity `json:"shares_1d"`
	CoinType                     string             `json:"coin_type"`
	EarningsYesterdayCoins       EarnStatsSmartItem `json:"earnings_yesterday_coins"`
	PaymentMode                  string             `json:"payment_mode"`
	EarningsBefore               bool               `json:"earnings_before"`
}

func (e *EarnStats) IsSmart() bool {
	return strings.HasPrefix(e.CoinType, "smart_")
}

func (e *EarnStats) GetCoin() string {
	if e.IsSmart() {
		return "BTC"
	} else {
		return strings.ToUpper(e.CoinType)
	}
}

type EarnStatsSmartItem struct {
	Btc string `json:"btc"`
	Bch string `json:"bch"`
	Bsv string `json:"bsv"`
}

type HashrateUnitEntity struct {
	Size string `json:"size"`
	Unit string `json:"unit"`
	Pure string `json:"pure"`
}

func GetEarnstats(c *gin.Context, params interface{}) (*EarnStats, error) {
	var dest = struct {
		BtcpoolRescomm
		Data EarnStats `json:"data"`
	}{}

	_, err := doRequest(c, "account.earnStats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.earnStats: %v", err)
	}

	return &dest.Data, nil
}

func GetMergeEarnstats(c *gin.Context, params interface{}) (map[string]EarnStats, error) {
	var dest = struct {
		BtcpoolRescomm
		Data EarnStats `json:"data"`
	}{}

	_, err := doRequest(c, "account.mergeEarnStats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.mergeEarnStats: %v", err)
	}
	res := make(map[string]EarnStats)
	res["earnstats"] = dest.Data
	return res, nil
}

type EarnHistory struct {
	DiffRate        string `json:"diff_rate"`
	PaymentTime     string `json:"payment_time"`
	Address         string `json:"address"`
	PaymentRedirect string `json:"payment_redirect"`
	Date            string `json:"date"`
	Stats           string `json:"stats"`
	EarnMode        string `json:"payment_mode"`
	Unit            string `json:"unit"`
	EarnModeMore    string `json:"more_than_pps96"`
	Earn            string `json:"earn"`
}

func GetEarnHistory(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.earnHistory", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.earnHistory: %v", err)
	}
	return dest.Data, nil
}

func GetMergeEarnHistory(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.mergeEarnHistory", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetMergeEarnHistory: %v", err)
	}
	return dest.Data, nil
}

func GetSubAccountAlgorithms(c *gin.Context, params interface{}) (clientModel.SubAccountAlgorithmList, error) {
	var dest = struct {
		BtcpoolRescomm
		Data clientModel.SubAccountAlgorithmList `json:"data"`
	}{}

	_, err := doRequest(c, "subaccount.algorithms", params, &dest)
	if err != nil {
		var res clientModel.SubAccountAlgorithmList
		return res, fmt.Errorf("error GetSubAccountAlgorithms: %v", err)
	}
	return dest.Data, nil
}

// page of clientModel.WorkGroupEntity
func GetWorkerStats(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.stats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetSubAccountAlgorithms: %v", err)
	}
	return dest.Data, nil
}
