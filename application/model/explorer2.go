package model

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"poolName"`
	Hash      string `json:"hash"`
	Size      string `json:"size"`
}

type PoolRankInfo struct {
	Index                 string `json:"index"`
	Icon                  string `json:"icon"`
	Name                  string `json:"name"`
	HashratePercent       string `json:"hashratePercent"`
	Hashrate              string `json:"hashrate"`
	HashrateUnit          string `json:"hashrateUnit"`
	HashrateChangePercent string `json:"hashrateChangePercent"`
	Lucy7Day              string `json:"lucy7Day"`
}