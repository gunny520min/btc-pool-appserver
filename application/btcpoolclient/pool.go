package btcpoolclient

import (
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"fmt"

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

type Earnstats struct {
	EarningsYesterday string `json:"earnings_yesterday"`
	EarningsToday     string `json:"earnings_today"`
	Unpaid            string `json:"unpaid"`
	Paid              string `json:"paid"`
}

func GetEarnstats(c *gin.Context, params interface{}) (map[string]Earnstats, error) {
	var dest = struct {
		BtcpoolRescomm
		Data Earnstats `json:"data"`
	}{}

	_, err := doRequest(c, "account.earnStats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.earnStats: %v", err)
	}
	res := make(map[string]Earnstats)
	res["earnstats"] = dest.Data
	return res, nil
}

func GetMergeEarnstats(c *gin.Context, params interface{}) (map[string]Earnstats, error) {
	var dest = struct {
		BtcpoolRescomm
		Data Earnstats `json:"data"`
	}{}

	_, err := doRequest(c, "account.mergeEarnStats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.mergeEarnStats: %v", err)
	}
	res := make(map[string]Earnstats)
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

func GetWorkerStats(c *gin.Context, params interface{}) (clientModel.WorkGroupEntity, error) {
	var dest = struct {
		BtcpoolRescomm
		Data clientModel.WorkGroupEntity `json:"data"`
	}{}

	_, err := doRequest(c, "worker.stats", params, &dest)
	if err != nil {
		var res clientModel.WorkGroupEntity
		return res, fmt.Errorf("error GetSubAccountAlgorithms: %v", err)
	}
	return dest.Data, nil
}
