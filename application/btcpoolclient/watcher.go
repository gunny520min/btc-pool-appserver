package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Watcher struct {
	Id              string   `json:"id"`
	Uid             string   `json:"uid"`
	Puid            string   `json:"puid"`
	WatcherKey      string   `json:"watcher_key"`
	Note            string   `json:"note"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	Token           string   `json:"token"` // access key
	Name            string   `json:"name"`
	Currency        string   `json:"currency"`
	HashrateValue   string   `json:"hashrate_value"`
	HashrateUnit    string   `json:"hashrate_unit"`
	Hashrate_suffix string   `json:"hashrate_suffix"`
	Coin_type       string   `json:"coin_type"`
	Region_id       string   `json:"region_id"`
	Default_url     string   `json:"default_url"`
	Redirect        string   `json:"redirect"`
	Authorities     []string `json:"authorities"`
	RegionNameCN    string   `json:"-"`
	RegionNameEN    string   `json:"-"`
	GrinValue       string   `json:"-"`
}

func GetWatcherList(c *gin.Context, params interface{}) ([]Watcher, error) {
	var dest = struct {
		BtcpoolRescomm
		Data ([]Watcher) `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.list", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.list: %v", err)
	}
	return dest.Data, nil
}

func CreateWatcher(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.create", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.create: %v", err)
	}
	return dest.Data, nil
}

func DeleteWatcher(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.delete", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.delete: %v", err)
	}
	return dest.Data, nil
}

func UpdateWatcher(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.update", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.update: %v", err)
	}
	return dest.Data, nil
}

func WatcherAuthority(c *gin.Context, params interface{}) ([]string, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]([]string) `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.authority", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.authority: %v", err)
	}

	return dest.Data["page_authorities"], nil
}

/// 添加一个观察者链接
func CheckWatcher(c *gin.Context, params interface{}) ([]Watcher, error) {
	var dest = struct {
		BtcpoolRescomm
		Data []Watcher `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.check", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error watcher.check: %v", err)
	}
	return dest.Data, nil
}

type WatcherRegion struct {
	SupportCoin string            `json:"region_support_coin"`
	Text        map[string]string `json:"text"`
}

func WatcherRegionInfo(c *gin.Context, params interface{}) (WatcherRegion, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]WatcherRegion `json:"data"`
	}{}

	_, err := doRequest(c, "watcher.info", params, &dest)
	if err != nil {
		var region WatcherRegion
		return region, fmt.Errorf("error watcher.info: %v", err)
	}
	return dest.Data["region_conf"], nil
}