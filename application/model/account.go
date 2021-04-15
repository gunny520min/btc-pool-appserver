package model

type MinerAddress struct {
	Addr string `json:"addr"`
	Tips string `json:"tips"`
}

type MinerConfig struct {
	Title   string         `json:"title"`
	Address []MinerAddress `json:"address"`
	Desc    string         `json:"desc"`
	Tips    string         `json:"tips"`
}
