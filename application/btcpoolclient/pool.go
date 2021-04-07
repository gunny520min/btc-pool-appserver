package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ShareHashrateData struct {
	Unit    string        `json:"unit"`
	Tickers []([]float64) `json:"tickers"`
}

func GetPoolShareHashrate(c *gin.Context, params interface{}) (ShareHashrateData, error) {
	var dest = struct {
		BtcpoolRescomm
		Data ShareHashrateData `json:"data"`
	}{}

	_, err := doRequest(c, "pool.hashrateHistory", params, &dest)
	if err != nil {
		var res ShareHashrateData
		return res, fmt.Errorf("error getting banner list: %v", err)
	}
	return dest.Data, nil

}
