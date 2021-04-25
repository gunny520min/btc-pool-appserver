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

type MergeEarnstats struct {
	EarningsYesterday string `json:"earningsYesterday"`
	EarningsToday     string `json:"earningsToday"`
	Unpaid            string `json:"unpaid"`
	Paid              string `json:"paid"`
	EarnUnit          string `json:"earnUnit"`
}

type MergeEarnHistory struct {
	DiffRate        string `json:"diffRate"`
	PaymentTime     string `json:"paymentTime"`
	Address         string `json:"address"`
	PaymentRedirect string `json:"paymentRedirect"`
	Date            string `json:"date"`
	Stats           string `json:"stats"`
	EarnMode        string `json:"earnMode"`
	Unit            string `json:"unit"`
	EarnModeMore    string `json:"earnModeMore"`
	Earn            string `json:"earn"`
}

type Earnstats struct {
	EarningsYesterday string `json:"earningsYesterday"`
	EarningsToday     string `json:"earningsToday"`
	Unpaid            string `json:"unpaid"`
	Paid              string `json:"paid"`
	EarnUnit          string `json:"earnUnit"`
}

type EarnHistory struct {
	DiffRate        string `json:"diffRate"`
	PaymentTime     string `json:"paymentTime"`
	Address         string `json:"address"`
	PaymentRedirect string `json:"paymentRedirect"`
	Date            string `json:"date"`
	Stats           string `json:"stats"`
	EarnMode        string `json:"earnMode"`
	Unit            string `json:"unit"`
	EarnModeMore    string `json:"earnModeMore"`
	Earn            string `json:"earn"`
}
