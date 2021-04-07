package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type poolHandler struct{}

var PoolService = &poolHandler{}

func (p *poolHandler) GetShareHashrate(c *gin.Context, params interface{}) btcpoolclient.ShareHashrateData {
	var ret btcpoolclient.ShareHashrateData
	if list, err := btcpoolclient.GetPoolShareHashrate(c, params); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeNu)
	} else {
		ret = list
	}
	return ret
}

func (p *poolHandler) FormatHashrateChartUnit(params btcpoolclient.ShareHashrateData) string {
	return params.Unit
}

func (p *poolHandler) FormatHashrateChartData(params btcpoolclient.ShareHashrateData) []model.HashrateData {
	res := make([]model.HashrateData, 0)
	for _, v := range params.Tickers {
		var item model.HashrateData
		item.Hashrate = v[1]
		item.Timestamp = v[0]
		res = append(res, item)
	}
	fmt.Printf("FormatHashrateChartData count= %v, %v", len(res), len(params.Tickers))
	return res
}
