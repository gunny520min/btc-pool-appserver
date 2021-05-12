package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"
	"btc-pool-appserver/application/model"
	"btc-pool-appserver/application/service"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetWatcherList(c *gin.Context) {
	var params AccountParams
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.GetWatcherList(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func CreateWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		Note        string `form:"note" binding:"required"`
		Lang        string `form:"lang" binding:"required"`
		Authorities string `form:"authorities" binding:"required"`
		GrinValue   string `form:"grin_value" binding:"-"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.CreateWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		res := d
		if len(params.GrinValue) > 0 {
			res["grin_value"] = params.GrinValue
		} else {
			res["grin_value"] = ""
		}
		output.Succ(c, d)
	}
}

func DeleteWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		WatcherId string `form:"watcher_id" binding:"required"`
		Lang      string `form:"lang" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.DeleteWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func UpdateWatcher(c *gin.Context) {
	var params struct {
		AccountParams
		WatcherId   string `form:"watcher_id" binding:"required"`
		Authorities string `form:"authorities" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.UpdateWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

func WatcherAuthority(c *gin.Context) {
	var params struct {
		AccountParams
		AccessKey string `form:"access_key" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if d, err := btcpoolclient.WatcherAuthority(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		// res := make(map[string]interface{})
		output.Succ(c, d)
	}
}

// AddOtherWatcher 添加查看别人的观察者链接
func AddOtherWatcher(c *gin.Context) {
	var params struct {
		Puids        string `form:"puids" binding:"required"`
		WatcherToken string `form:"watcher_token" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := service.WatcherService.AddOtherWatcher(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetWatcherHashrate(c *gin.Context) {
	var params struct {
		Puids        string `form:"puids" binding:"required"`
		WatcherToken string `form:"watcher_token" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := service.WatcherService.GetWatcherHashrate(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetWatcherDashboard(c *gin.Context) {
	var params struct {
		AccessKey string `form:"accessKey"`
		Puid      string `form:"puid"`
	}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		panic(err)
	}
	var res model.WatcherDashboard
	var resErr error = nil
	var resIncomeErr error = nil
	var resWorkerStateErr error = nil
	var resWorkerGroupErr error = nil
	var resMiningAddressErr error = nil
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//notice
		defer wg.Done()

		if userAuthority, err := service.WatcherService.GetWatcherAuthority(c, params.AccessKey, params.Puid); err != nil {
			resErr = err
			return
		} else {
			res.Authorities = userAuthority.PageAuthorities
		}
	}()
	wg.Wait()
	if resErr != nil {
		output.ShowErr(c, resErr)
		return
	}
	wg.Add(4)
	go func() {
		defer wg.Done()
		if income, err := service.WatcherService.GetWatcherDashboardIncome(c, params.AccessKey, params.Puid); err != nil {
			resIncomeErr = err
		} else {
			res.Income = income
		}
	}()

	go func() {
		defer wg.Done()
		if workerStats, err := service.WatcherService.GetWatcherDashboardWorkerStates(c, params.AccessKey, params.Puid); err != nil {
			resWorkerStateErr = err
		} else {
			res.WorkerStatus = workerStats
		}
	}()

	go func() {
		defer wg.Done()
		if wgs, err := service.WatcherService.GetWatcherDashboardWorkerGroup(c, params.AccessKey, params.Puid); err != nil {
			resWorkerGroupErr = err
		} else {
			res.WorkerGroup = wgs
		}
	}()

	go func() {
		defer wg.Done()
		if sba, err := service.WatcherService.GetWatcherDashboardMiningAddress(c, params.AccessKey, params.Puid); err != nil {
			resMiningAddressErr = err
		} else {
			res.MiningAddress = sba
		}
	}()

	wg.Wait()
	if resIncomeErr != nil {
		output.ShowErr(c, resIncomeErr)
		return
	}
	// CoinType只能由income接口获取，所以延后处理一些单位的逻辑
	if resWorkerStateErr == nil {
		res.WorkerStatus.TotalHashrate.Unit = res.WorkerStatus.TotalHashrate.Unit + model.GetCoinSuffixByCoinType(res.Income.CoinType)
	}
	if resWorkerGroupErr == nil {
		for index, group := range res.WorkerGroup {
			group.Hashrate.Unit = group.Hashrate.Unit + model.GetCoinSuffixByCoinType(res.Income.CoinType)
			res.WorkerGroup[index] = group
		}
	}
	if resMiningAddressErr == nil {
		for index, address := range res.MiningAddress.Address {
			url := strings.Split(address.Addr, " ")
			if index == 0 && len(url) == 2 && len(url[1]) > 0 {
				if strings.ToUpper(res.Income.CoinType) == "BEAM" {
					address.Tips = url[0] + " 支持普通显卡矿机\n" + url[1] + " 为Nicehash特别优化"
				}
				if strings.ToUpper(res.Income.CoinType) == "DCR" {
					address.Tips = url[0] + " 已知支持蚂蚁矿机DR3，Claymore，gominer\n" + url[1] + " 已知支持Nicehash，芯动科技等矿机"
				}
			}
		}
	}

	output.Succ(c, res)
}

func GetWatcherDashboardWorkerShareHistory(c *gin.Context) {
	var params struct {
		AccessKey string `form:"accessKey"`
		Puid      string `form:"puid"`
		Dimension string `form:"dimension"`
	}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		output.ShowErr(c, err)
		return
	}
	if workerShareHistory, err := service.WatcherService.GetWatcherDashboardWorkerShareHistory(c, params.AccessKey, params.Puid, params.Dimension); err != nil {
		output.ShowErr(c, err)
	} else {
		output.Succ(c, workerShareHistory)
	}
}

func GetWatcherEarnStats(c *gin.Context) {
	var params struct {
		Puid      string `form:"puid"`
		AccessKey string `form:"accessKey"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	if stats, err := service.WatcherService.GetWatcherDashboardIncome(c, params.AccessKey, params.Puid); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, stats)
	}
}


func GetWatcherEarnHistory(c *gin.Context) {
	var params struct {
		AccountParams
		PageParams
		IsDecimal string `form:"is_decimal" binding:"-"`
		AccessKey string `form:"access_key"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	params.IsDecimal = "1"
	if d, err := btcpoolclient.GetEarnHistory(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

