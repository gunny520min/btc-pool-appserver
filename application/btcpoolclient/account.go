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

func SubacountHiiden(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.hidden", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.hidden: %v", err)
	}
	return dest.Data, nil
}

func SubacountHiidenCancel(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "account.hiddenCancel", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error account.hiddenCancel: %v", err)
	}
	return dest.Data, nil
}
