package model

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"poolName"`
	Hash      string `json:"hash"`
	Size      string `json:"size"`
}

type PoolRank struct {
	Rank     string `json:"rank"`
	PoolName string `json:"poolName"`
	IconLink string `json:"iconLink"`
	Progress string `json:"progress"`
	Diff     string `json:"diff"`
	Hashrate string `json:"hashrate"`
}
