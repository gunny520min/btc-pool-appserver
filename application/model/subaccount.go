package model

type SubaccountEntity struct {
	Title      string                      `json:"title"`
	SearchKey  string                      `json:"searchKey"`
	Algorithms []SubaccountAlgorithmEntity `json:"algorithms"`
}

type SubaccountAlgorithmEntity struct {
	AlgorithmText string                                 `json:"algorithmText"`
	CurrentCoin   string                                 `json:"currentCoin"`
	CurrentPuid   string                                 `json:"currentPuid"`
	IsSmart       bool                                   `json:"isSmart"`
	SupportCoins  []string                               `json:"supportCoins"`
	SubAccount    []SubaccountAlgorithmCoinAccountEntity `json:"subAccount"`
}

type SubaccountAlgorithmCoinAccountEntity struct {
	Puid      string `json:"puid"`
	CoinType  string `json:"coinType"`
	IsHidden  bool   `json:"isHidden"`
	IsCurrent bool   `json:"IsCurrent"`
}
