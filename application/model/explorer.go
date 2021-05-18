package model

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"pool_name"`
	Hash      string `json:"hash"`
	Size      string `json:"size"`
}

type PoolRank struct {
	PoolName               string `json:"pool_name"`
	IconLink               string `json:"icon_link"`
	RealtimeHashrate       string `json:"realtime_hashrate"`
	EstimateHashrate       string `json:"estimate_hashrate"`
	RealtimeCur2maxPercent string `json:"realtime_cur2max_percent"`
	EstimateCur2max        string `json:"estimate_cur2max"`
	HashSuffix             string `json:"hashrate_suffix"`
	RealtimeDiff24hPercent string `json:"realtime_diff_24h_percent"`
}
