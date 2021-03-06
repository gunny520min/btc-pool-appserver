package btcpoolclient

import (
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/library/errs"

	"github.com/gin-gonic/gin"
)

type WorkerStats struct {
	Gid             string  `json:"gid"`
	Name            string  `json:"name"`
	Hashrate        float64 `json:"shares_15m"`
	Unit            string  `json:"shares_unit"`
	WorkersActive   string  `json:"workers_active"`
	WorkersInactive string  `json:"workers_inactive"`
	WorkersDead     string  `json:"workers_dead"`
	WorkersTotal    string  `json:"workers_total"`
	SortId          int     `json:"sort_id"`
}
type Worker struct {
	Gid             string  `json:"gid"`
	GroupName       string  `json:"group_name"`
	WorkerId        string  `json:"worker_id"`
	Name            string  `json:"worker_name"`
	SharesUnit      string  `json:"shares_unit"`
	WorkersActive   string  `json:"workers_active"`
	WorkersInactive string  `json:"workers_inactive"`
	WorkersDead     string  `json:"workers_dead"`
	WorkersTotal    string  `json:"workers_total"`
	SortId          int     `json:"sort_id"`
	Shares1m        float64 `json:"shares_1m"`
	Shares5m        float64 `json:"shares_5m"`
	Shares15m       float64 `json:"shares_15m"`
	Shares1d        float64 `json:"shares_1d"`
	Shares1dUnit    string  `json:"shares1dUnit"`
	AcceptCount     float64 `json:"accept_count"`
	AcceptPercent   float64 `json:"accept_percent"`
	TotalAccept     float64 `json:"total_accept"`
	RejectPercent   float64 `json:"reject_percent"`
	LastShareTime   float64 `json:"last_share_time"`
	LastShareIp     string  `json:"last_share_ip"`
	Ip2location     string  `json:"ip2location"`
	Status          string  `json:"status"`
	WorkerStatus    string  `json:"worker_status"`
}

func WorkerGroups(c *gin.Context, params interface{}) ([]clientModel.WorkerGroupEntity, error) {
	var dest = struct {
		BtcpoolRescomm
		Data struct {
			BtcpoolPageRescomm
			List []clientModel.WorkerGroupEntity `json:"list"`
		} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.groups", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerGroups: %v", err)
	}
	return dest.Data.List, nil
}

func WorkerGroupsDelete(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.deleteGroup", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerGroupsDelete: %v", err)
	}
	return dest.Data, nil
}

func WorkerGroupsCreate(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.createGruop", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerGroupsCreate: %v", err)
	}
	return dest.Data, nil
}

// ????????????
func WorkerDelete(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.update", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerDelete: %v", err)
	}
	return dest.Data, nil
}

// ????????????
func WorkerMove(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.move", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerMove: %v", err)
	}
	return dest.Data, nil
}

//  ????????????
func WorkerList(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "worker.list", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerList: %v", err)
	}
	return dest.Data, nil
}

//  ????????????
func WorkerDetail(c *gin.Context, params interface{}) (map[string]Worker, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]Worker `json:"data"`
	}{}

	workerId := c.Query("workerId")
	if len(workerId) == 0 {
		return nil, errs.ApiErrParams //fmt.Errorf("request WorkerDetail paramaters no workerId")
	}

	_, err := doRequestEx(c, "worker.detailInfo", "/"+workerId, params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerDetail: %v", err)
	}
	return dest.Data, nil
}

//  ????????????
func WorkerHashrate(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	workerId := c.Query("workerId")
	if len(workerId) == 0 {
		return nil, errs.ApiErrParams //fmt.Errorf("request WorkerHashrate paramaters no workerId")
	}

	_, err := doRequestEx(c, "worker.hashrateHistory", "/"+workerId+"/share-history/", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error WorkerHashrate: %v", err)
	}
	return dest.Data, nil
}
