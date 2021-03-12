package btcpoolclient

import (
	"btc-pool-appserver/application/config"
	"btc-pool-appserver/application/library/third"
	"fmt"

	"github.com/gin-gonic/gin"
)

// BtcpoolRescomm ..
type BtcpoolRescomm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// GetCode ..
func (b *BtcpoolRescomm) GetCode() int {
	return b.Code
}

// GetMessage ..
func (b *BtcpoolRescomm) GetMessage() string {
	return b.Msg
}

func doRequest(c *gin.Context, action string, params interface{}, dest interface{}) ([]byte, error) {
	// 根据action获取配置请求
	apiConfig, exist := config.Btcpool.Apis[action]
	if !exist {
		return nil, fmt.Errorf("btcpool dorequest unknown action: %s", action)
	}

	// 发起http请求
	if res, err := third.DoActionRequest(c, &apiConfig, params, nil, dest); err != nil {
		return nil, fmt.Errorf("btcpool: %w", err)
	} else {
		return res, nil
	}
}
