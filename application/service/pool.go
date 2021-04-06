package service

import (
	"btc-pool-appserver/application/btcpoolclient"

	"github.com/gin-gonic/gin"
)

type poolHandler struct{}

var PoolService = &poolHandler{}

func (p *poolHandler) GetShareHashrate(c *gin.Context, params interface{}) interface{} {
	var ret interface{}
	if list, err := btcpoolclient.GetPoolShareHashrate(c, params); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeNu)
	} else {
		ret = list
	}
	return ret
}
