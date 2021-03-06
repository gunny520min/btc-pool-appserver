package btcpoolclient

import (
	"btc-pool-appserver/application/library/errs"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateMergeCoinAddress(c *gin.Context, params interface{}) (map[string]interface{}, error) {

	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	coinType := c.Query("coinType")
	if len(coinType) == 0 {
		return nil, errs.ApiErrParams // fmt.Errorf("request UpdateMergeCoinAddress paramaters no coin type")
	}

	_, err := doRequestEx(c, "merge.updateAddress", "/"+strings.ToLower(coinType)+"/address/update", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error UpdateMergeCoinAddress: %v", err)
	}
	return dest.Data, nil
}
