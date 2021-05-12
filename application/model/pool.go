package model

type Income struct {
	Income      NormalIncome `json:"income"`
	SmartIncome SmartIncome  `json:"smartIncome"`
	HasIncome   bool         `json:"hasIncome"`
	CoinType    string       `json:"coinType"`
}

type WatcherDashboard struct {
	Income
	IsSmart       bool          `json:"isSmart"`
	WorkerStatus  WorkerStatus  `json:"workerStatus"`
	WorkerGroup   []WorkerGroup `json:"workerGroup"`
	MiningAddress MiningAddress `json:"miningAddress"`
	Authorities   []string      `json:"authorities"`
}

type Dashboard struct {
	Income
	IsSmart       bool          `json:"isSmart"`
	Puid          string        `json:"puid"`
	//CoinType      string        `json:"coinType"`
	Title         string        `json:"title"`
	WorkerStatus  WorkerStatus  `json:"workerStatus"`
	WorkerGroup   []WorkerGroup `json:"workerGroup"`
	MiningAddress MiningAddress `json:"miningAddress"`
}
type ValueUnit struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
	Coin  string `json:"coin"`
}

type NormalIncome struct {
	IncomeYesterday ValueUnit `json:"incomeYesterday"`
	IncomeToday     ValueUnit `json:"incomeToday"`
	IncomePaid      ValueUnit `json:"incomePaid"`
	IncomeUnpaid    ValueUnit `json:"incomeUnpaid"`
}

type SmartIncome struct {
	IsOtc           bool                 `json:"isOtc"`
	IncomeYesterday SmartIncomeYesterday `json:"incomeYesterday"`
	IncomeToday     ValueUnit            `json:"incomeToday"`
}

type SmartIncomeYesterday struct {
	Btc ValueUnit `json:"btc"`
	Bch ValueUnit `json:"bch"`
	Bsv ValueUnit `json:"bsv"`
}

type WorkerStatus struct {
	TotalHashrate  ValueUnit `json:"totalHashrate"`
	WorkerActive   string    `json:"workerActive"`
	WorkerInactive string    `json:"workerInactive"`
}

type WorkerGroup struct {
	Gid          string    `json:"gid"`
	Name         string    `json:"name"`
	Hashrate     ValueUnit `json:"hashrate"`
	WorkerActive string    `json:"workerActive"`
	WorkerTotal  string    `json:"workerTotal"`
}

type MiningAddress struct {
	Title   string                `json:"title"`
	Address []MiningAddressDetail `json:"address"`
	Desc    string                `json:"desc"`
	Tips    string                `json:"tips"`
}

type MiningAddressDetail struct {
	Addr string `json:"addr"`
	Tips string `json:"tips"`
}

type WorkerShareHistoryEntity struct {
	Timestamp string `json:"timestamp"`
	Hashrate  string `json:"hashrate"`
	Reject    string `json:"reject"`
}

type WorkerShareHistory struct {
	Unit string                     `json:"unit"`
	List []WorkerShareHistoryEntity `json:"list"`
}
