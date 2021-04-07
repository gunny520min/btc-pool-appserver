package controller

import (
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func MultiCoinStats(c *gin.Context) {

}

func HashrateHistory(c *gin.Context) {
	params := make(map[string]interface{})
	res := make(map[string]interface{})
	params["dimension"] = "1h"
	params["count"] = 72
	params["real_point"] = "1"

	coin := c.Query("coin")
	if len(coin) == 0 {
		output.ShowErr(c, errs.ApiErrParams)
	} else {
		fmt.Println("get hashrate")
		params["coin_type"] = coin

		if p, err := urlEncoded(params); err != nil {
			res["histories"] = make(map[string]interface{})
			res["unit"] = ""
			output.Succ(c, res)
		} else {
			shareData := service.PoolService.GetShareHashrate(c, p)
			res["histories"] = service.PoolService.FormatHashrateChartData(shareData)
			res["unit"] = service.PoolService.FormatHashrateChartUnit(shareData)
			output.Succ(c, res)
		}
	}
}
