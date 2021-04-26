package service

import (
	"btc-pool-appserver/application/btcpoolclient"

	"github.com/gin-gonic/gin"
)

type watcherHandler struct{}

var WatcherService = &watcherHandler{}

func (p *watcherHandler) AddOtherWatcher(c *gin.Context, params interface{}) (btcpoolclient.Watcher, error) {

	var w btcpoolclient.Watcher
	//check watcher
	if d, err := btcpoolclient.CheckWatcher(c, params); err != nil {
		return w, err
	} else if len(d) > 0 {
		w = d[0]
	} else {
		return w, nil
	}
	// get watcher info
	pms := map[string]string{}
	pms["puid"] = w.Puid
	pms["access_key"] = w.Token
	if d, err := btcpoolclient.WatcherRegionInfo(c, pms); err != nil {
		return w, err
	} else {
		w.RegionNameCN = d.Text["zh-cn"]
		w.RegionNameEN = d.Text["en"]
	}
	return w, nil
}
