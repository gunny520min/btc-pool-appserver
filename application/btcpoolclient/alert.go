package btcpoolclient

import (
	"github.com/gin-gonic/gin"
)

type AlertSetting struct {
	HashrateValue  string `json:"hashrate_value"`
	HashrateUnit   string `json:"hashrate_unit"`
	HashrateSuffix string `json:"hashrate_suffix"`
	HashrateAlert  string `json:"hashrate_alert"`
	MinerValue     string `json:"miner_value"`
	MinerAlert     string `json:"miner_alert"`
	AlertInterval  string `json:"alert_interval"`
}

func AlertSettings(c *gin.Context, params interface{}) (map[string]AlertSetting, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]AlertSetting `json:"data"`
	}{}

	_, err := doRequest(c, "alert.settings", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.settings: %v", err)
	}
	return dest.Data, nil
}

type AlertContact struct {
	Id         string `json:"id"`
	Uid        string `json:"uid"`
	Puid       string `json:"puid"`
	Note       string `json:"note"`
	Email      string `json:"email"`
	RegionCode string `json:"region_code"`
	Phone      string `json:"phone"`
	IsDefault  string `json:"is_default"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func AlertContacts(c *gin.Context, params interface{}) (map[string][]AlertContact, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string][]AlertContact `json:"data"`
	}{}

	_, err := doRequest(c, "alert.contacts", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.contacts: %v", err)
	}
	return dest.Data, nil
}

func UpdateAlertHashrate(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.hashrateUpdate", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.hashrateUpdate: %v", err)
	}
	return dest.Data, nil
}

func UpdateAlertMiners(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.minersUpdate", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.minersUpdate: %v", err)
	}
	return dest.Data, nil
}

func UpdateAlertInterval(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.intervalUpdate", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.intervalUpdate: %v", err)
	}
	return dest.Data, nil
}

func DeleteAlertContact(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.contactDelete", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.contactDelete: %v", err)
	}
	return dest.Data, nil
}

func CreateAlertContact(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.contactCreate", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.contactCreate: %v", err)
	}
	return dest.Data, nil
}

func UpdateAlertContact(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.contactUpdate", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.contactUpdate: %v", err)
	}
	return dest.Data, nil
}

/// 报警列表
type AlarmClassification struct {
	Id        string `json:"id"`
	Uid       string `json:"uid"`
	CreatedAt string `json:"created_at"`
	Puid      string `json:"puid"`
	Content   string `json:"content"`
}

func GetAlerMerge(c *gin.Context, params interface{}) ([]AlarmClassification, error) {
	var dest = struct {
		BtcpoolRescomm
		Data ([]AlarmClassification) `json:"data"`
	}{}

	_, err := doRequest(c, "alert.list", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.list: %v", err)
	}
	return dest.Data, nil
}

func AlertRead(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "alert.read", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error alert.read: %v", err)
	}
	return dest.Data, nil
}
