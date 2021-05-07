package clientModel

import "strings"

type SubAccountAlgorithmList struct {
	HaveDefault bool                  `json:"have_default"`
	Account     string                `json:"account"`
	SubAccounts []SubAccountAlgorithm `json:"subaccounts"`
}

type SubAccountAlgorithm struct {
	Name           string                      `json:"name"`
	HaveOtherCoins string                      `json:"have_other_coins"`
	RegionText     string                      `json:"region_text"`
	Algorithms     []SubAccountAlgorithmDetail `json:"algorithms"`
}

type SubAccountAlgorithmDetail struct {
	AlgorithmName   string                 `json:"algorithm_name"`
	AlgorithmText   string                 `json:"algorithm_text"`
	CurrentCoin     string                 `json:"CurrentCoin"`
	CurrentPuid     string                 `json:"current_puid"`
	CurrentMode     string                 `json:"current_mode"`
	CurrentModeText string                 `json:"current_mode_text"`
	OpenCoinType    []string               `json:"open_coin_type"`
	SupportCoins    []string               `json:"support_coins"`
	CoinAccounts    []SubAccountCoinEntity `json:"coin_accounts"`
}

func (sa *SubAccountAlgorithmDetail) IsSmart() bool {
	return strings.HasPrefix(sa.CurrentMode, "smart_")
}

type SubAccountCoinEntity struct {
	RegionServiceName string `json:"region_service_name"`
	Puid              string `json:"puid"`
	RegionId          string `json:"region_id"`
	RegionText        string `json:"region_text"`
	Name              string `json:"name"`
	DefaultUrl        string `json:"default_url"`
	RegionName        string `json:"region_name"`
	CountryName       string `json:"country_name"`
	CoinType          string `json:"coin_type"`
	CoinAlgorithm     string `json:"coin_algorithm"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Shares            string `json:"shares"`
	IsHidden          int    `json:"is_hidden"` // 1 隐藏，0 显示
	IsGenerate        int    `json:"is_generate"`
	IsDefault         int    `json:"is_default"`
	IsCurrent         int    `json:"is_current"`// 1 当前正在挖， 0 未在挖
	ChangeUpdatedAt   string `json:"change_updated_at"`
}

func (sa *SubAccountCoinEntity) IsSmart() bool {
	return strings.HasPrefix(sa.CoinType, "smart_")
}

type WorkerGroupEntity struct {
	Gid             string `json:"gid"`
	WorkersDead     string `json:"workers_dead"`
	WorkersActive   string `json:"workers_active"`
	WorkersInactive string `json:"workers_inactive"`
	Name            string `json:"name"`
	Shares1m        string `json:"shares_1m"`
	Shares5m        string `json:"shares_5m"`
	Shares15m       string `json:"shares_15m"`
	RejectPercent   string `json:"reject_percent"`
	WorkersTotal    string `json:"workers_total"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	SortId          string `json:"sort_id"`
	SharesUnit      string `json:"shares_unit"`
}

type WorkerShareHistory struct {
	SharesUnit string     `json:"shares_unit"`
	Tickers    [][]string `json:"tickers"`
}
