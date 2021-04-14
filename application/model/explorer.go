package model

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"poolName"`
	Hash      string `json:"hash"`
	Size      int    `json:"size"`
}

type PoolRank struct {
	PoolName               string `json:"poolName"`
	IconLink               string `json:"iconLink"`
	RealtimeHashrate       string `json:"realtimeHashrate"`
	EstimateHashrate       string `json:"estimateHashrate"`
	RealtimeCur2maxPercent string `json:"realtimeCur2maxPercent"`
	EstimateCur2max        string `json:"estimateCur2max"`
	HashSuffix             string `json:"hashSuffix"`
	RealtimeDiff24hPercent string `json:"realtimeDiff24hPercent"`
}
