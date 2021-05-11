package clientModel

type WatcherAuthority struct {
	EarningDownload bool     `json:"earning_download"`
	MinerExportAuth bool     `json:"miner_export_auth"`
	PageAuthorities []string `json:"page_authorities"`
}
