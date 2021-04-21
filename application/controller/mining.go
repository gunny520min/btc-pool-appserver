package controller

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/output"

	"github.com/gin-gonic/gin"
)

func GetMinerGroups(c *gin.Context) {

	var params AccountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}
	var p = struct {
		AccountParams
		Page     string `json:"page"`
		PageSize string `json:"page_size"`
	}{
		AccountParams: params,
		Page:          "1",
		PageSize:      "100",
	}
	if d, err := btcpoolclient.WorkerGroups(c, p); err != nil {
		output.ShowErr(c, err)
		return
	} else {
		output.Succ(c, d)
	}
}

func MinerGroupDelete(c *gin.Context) {
	var params struct {
		AccountParams
		Gid string `json:"gid"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
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
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
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
		WorkerIds string `json:"workerIds"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	var p = struct {
		AccountParams
		GroupId   string `json:"group_id"`
		WorkerIds string `json:"worker_id"`
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
		WorkerIds string `json:"workerIds"`
		GroupId   string `json:"groupId"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		output.ShowErr(c, errs.ApiErrParams)
		return
	}

	var p = struct {
		AccountParams
		GroupId   string `json:"group_id"`
		WorkerIds string `json:"worker_id"`
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
		Status    string `json:"status"`
		Order_by  string `json:"order_by"`
		Asc       string `json:"asc"`
		Group     string `json:"group"`
		Filter    string `json:"filter" binding:"-"`
		AccessKey string `json:"access_key" binding:"-"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
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
		WorkerId  string `json:"workerId"`
		AccessKey string `json:"access_key" binding:"-"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
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
		Dimension string `json:"dimension"`
		Start_ts  string `json:"start_ts"`
		Count     string `json:"count"`
		RealPoint string `json:"real_point" binding:"-"`
		AccessKey string `json:"access_key" binding:"-"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
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
