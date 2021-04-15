package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAccountInfo(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}
	if _, err := doRequest(c, "account.info", params, &dest); err != nil {
		return nil, fmt.Errorf("error GetAccountInfo: %v", err)
	}
	return dest.Data, nil
}

type SubaccountPayset struct {
	Address     string   `json:"address"`
	UpdateAt    float64  `json:"update_at"`
	PayLimit    string   `json:"pay_limit"`
	PayLimitSet []string `json:"pay_limit_set"`
}

func GetSubaccountPayset(c *gin.Context, params interface{}) (map[string]SubaccountPayset, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]SubaccountPayset `json:"data"`
	}{}

	_, err := doRequest(c, "account.payset", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetSubaccountPayset: %v", err)
	}
	return dest.Data, nil
}

func UpdateSubaccountPayAddress(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.paysetAddressUpdate", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error UpdateSubaccountPayAddress: %v", err)
	}
	return dest.Data, nil
}

func UpdateSubaccountPayLimit(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.paylimitUpdate", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error UpdateSubaccountPayLimit: %v", err)
	}
	return dest.Data, nil
}

func GetAccountMinerConfig(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.minerConfig", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetAccountMinerConfig: %v", err)
	}
	return dest.Data, nil
}

type Earnstats struct {
	EarningsYesterday string `json:"earnings_yesterday"`
	EarningsToday     string `json:"earnings_today"`
	Unpaid            string `json:"unpaid"`
	Paid              string `json:"paid"`
}

func GetMergeEarnstats(c *gin.Context, params interface{}) (map[string]Earnstats, error) {
	var dest = struct {
		BtcpoolRescomm
		Data Earnstats `json:"data"`
	}{}

	_, err := doRequest(c, "account.mergeEarnStats", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetMergeEarnstats: %v", err)
	}
	res := make(map[string]Earnstats)
	res["earnstats"] = dest.Data
	return res, nil
}

type MergeEarnHistory struct {
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

func GetMergeEarnHistory(c *gin.Context, params interface{}) (map[string]([]MergeEarnHistory), error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]([]MergeEarnHistory) `json:"data"`
	}{}

	_, err := doRequest(c, "account.mergeEarnHistory", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetMergeEarnHistory: %v", err)
	}
	return dest.Data, nil
}
