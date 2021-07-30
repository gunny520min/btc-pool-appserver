package btcpoolclient

import (
	"btc-pool-appserver/application/btcpoolclient/clientModel"

	"github.com/gin-gonic/gin"
)

func GetSubAccountInfo(c *gin.Context, params interface{}) (*clientModel.SubAccountEntity, error) {
	var dest = struct {
		BtcpoolRescomm
		Data clientModel.SubAccountEntity `json:"data"`
	}{}
	if _, err := doRequest(c, "subaccount.info", params, &dest); err != nil {
		return nil, err //fmt.Errorf("error GetAccountInfo: %v", err)
	}
	return &dest.Data, nil
}

func GetAccountInfo(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}
	if _, err := doRequest(c, "account.info", params, &dest); err != nil {
		return nil, err //fmt.Errorf("error GetAccountInfo: %v", err)
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
		return nil, err //fmt.Errorf("error GetSubaccountPayset: %v", err)
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
		return nil, err //fmt.Errorf("error UpdateSubaccountPayAddress: %v", err)
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
		return nil, err //fmt.Errorf("error UpdateSubaccountPayLimit: %v", err)
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
		return nil, err //fmt.Errorf("error GetAccountMinerConfig: %v", err)
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
		return nil, err //fmt.Errorf("error account.hidden: %v", err)
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
		return nil, err //fmt.Errorf("error account.hiddenCancel: %v", err)
	}
	return dest.Data, nil
}

func GetSubaccountHashrate(c *gin.Context, params interface{}) (map[string]clientModel.SubAccountHashrateDetail, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]clientModel.SubAccountHashrateDetail `json:"data"`
	}{}

	_, err := doRequest(c, "subaccount.hashrate", params, &dest)
	if err != nil {
		return nil, err
	}
	return dest.Data, nil
}

func SubaccountChangeHashrate(c *gin.Context, params interface{}) (*clientModel.ChangeHashrateRes, error) {
	var dest = struct {
		BtcpoolRescomm
		Data clientModel.ChangeHashrateRes `json:"data"`
	}{}

	_, err := doRequest(c, "subaccount.changeHashrate", params, &dest)
	if err != nil {
		return nil, err
	}
	return &dest.Data, nil
}

func SubaccountCreateInit(c *gin.Context, params interface{}) (*clientModel.CreateSubaccountInitRes, error) {
	var dest = struct {
		BtcpoolRescomm
		Data clientModel.CreateSubaccountInitRes `json:"data"`
	}{}

	_, err := doRequest(c, "subaccount.createInit", params, &dest)
	if err != nil {
		return nil, err
	}
	return &dest.Data, nil
}
