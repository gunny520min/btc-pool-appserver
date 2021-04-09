package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取矿池排名
func GetPoolRank(c *gin.Context) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "explorer.poolRank", "", &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetPoolRank: %v", err)
	}
	return dest.Data, nil
}

// 获取最新出块列表
func GetLatestBlockList(c *gin.Context) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "explorer.blockList", "", &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetLatestBlockList: %v", err)
	}
	return dest.Data, nil
}
