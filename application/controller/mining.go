package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"

	"github.com/gin-gonic/gin"
)

func GetMinerGroups(c *gin.Context) {

	var params AccountParams
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	var p = struct {
		AccountParams
		Page     string `form:"page"`
		PageSize string `form:"page_size"`
	}{
		AccountParams: params,
		Page:          "1",
		PageSize:      "100",
	}
	if d, err := btcpoolclient.WorkerGroups(c, p); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, map[string][]clientModel.WorkerGroupEntity{
			"list": d,
		})
	}
}

func MinerGroupDelete(c *gin.Context) {
	var params struct {
		AccountParams
		Gid string `form:"gid"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.WorkerGroupsDelete(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func MinerGroupCreate(c *gin.Context) {
	var params struct {
		AccountParams
		Name string `form:"name"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.WorkerGroupsCreate(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func MinerWorkerDelete(c *gin.Context) {
	var params struct {
		AccountParams
		WorkerIds string `form:"workerIds"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	var p = struct {
		AccountParams
		GroupId   string `form:"group_id"`
		WorkerIds string `form:"worker_id"`
	}{
		AccountParams: params.AccountParams,
		GroupId:       "0",
		WorkerIds:     params.WorkerIds,
	}
	if d, err := btcpoolclient.WorkerDelete(c, p); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func MinerWorkerMove(c *gin.Context) {
	var params struct {
		AccountParams
		WorkerIds string `form:"workerIds"`
		GroupId   string `form:"groupId"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	var p = struct {
		AccountParams
		GroupId   string `form:"group_id"`
		WorkerIds string `form:"worker_id"`
	}{
		AccountParams: params.AccountParams,
		GroupId:       params.GroupId,
		WorkerIds:     params.WorkerIds,
	}
	if d, err := btcpoolclient.WorkerMove(c, p); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetMinerWorkerList(c *gin.Context) {
	var params struct {
		AccountParams
		PageParams
		Status    string `form:"status"`
		Order_by  string `form:"order_by"`
		Asc       string `form:"asc"`
		Group     string `form:"group"`
		Filter    string `form:"filter" binding:"-"`
		AccessKey string `form:"access_key" binding:"-"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.WorkerList(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetMinerWorkerDetail(c *gin.Context) {
	var params struct {
		AccountParams
		WorkerId  string `form:"workerId"`
		AccessKey string `form:"access_key" binding:"-"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	if d, err := btcpoolclient.WorkerDetail(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func GetMinerWorkerHashrate(c *gin.Context) {
	var params struct {
		AccountParams
		Dimension string `form:"dimension"`
		Start_ts  string `form:"start_ts"`
		Count     string `form:"count"`
		RealPoint string `form:"real_point" binding:"-"`
		AccessKey string `form:"access_key" binding:"-"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	params.RealPoint = "true"
	if d, err := btcpoolclient.WorkerHashrate(c, params); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}
