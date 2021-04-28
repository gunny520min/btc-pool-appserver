package model

/*
merge_mining_coins": {
			"nmc": "normal",
			"ela": "normal",
			"rsk": "offline",
			"vcash": "normal"
		},
"merge_mining_coins_config": {
			"nmc": {
				"helper_link": "https:\/\/help.pool.btc.com\/hc\/zh-cn\/articles\/360020805651"
			},
			"ela": {
				"helper_link": "https:\/\/help.pool.btc.com\/hc\/zh-cn\/articles\/360026754171"
			},
			"rsk": {
				"helper_link": "https:\/\/help.pool.btc.com\/hc\/zh-cn\/articles\/360020527351"
			},
			"vcash": {
				"helper_link": "https:\/\/help.pool.btc.com\/hc\/zh-cn\/articles\/360029847871"
			}
		},
"merge_mining_addresses": {
			"nmc": "",
			"ela": "",
			"rsk": "",
			"vcash": ""
		},
*/
type MergeCoin struct {
	CoinType  string `json:"coinType"`
	Title     string `json:"title"`
	Offline   bool   `json:"offline"`
	HelpUrl   string `json:"helpUrl"`
	HelpTitle string `json:"helpTitle"`
	Address   string `json:"address"`
}
