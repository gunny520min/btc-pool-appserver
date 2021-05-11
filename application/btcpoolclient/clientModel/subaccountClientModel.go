package clientModel

type SubAccountEntity struct {
	Uid            string               `json:"uid"`
	Puid           string               `json:"puid"`
	RegionId       string               `json:"region_id"`
	RegionName     string               `json:"region_names"`
	Name           string               `json:"name"`
	IsHidden       int                  `json:"is_hidden"`
	StratumUrlConf []string             `json:"stratum_url_conf"`
	CoinType       string               `json:"coin_type"`
	UserInfo       SubAccountUserInfo   `json:"user_info"`
	RegionConf     SubAccountRegionConf `json:"region_conf"`
}

type SubAccountUserInfo struct {
	Mail      string      `json:"mail"`
	Phone     PhoneNumber `json:"phone"`
	AvatarPic string      `json:"avatar_pic"`
}

type PhoneNumber struct {
	Country string `json:"country"`
	Number  string `json:"number"`
}

type SubAccountRegionConf struct {
	Text map[string]string `json:"text"`
}

type SubAccountHashrateDetail struct {
	Puid            string                  `json:"puid"`
	Workers         int                     `json:"workers"`
	WorkersActive   int                     `json:"workers_active"`
	WorkersInactive int                     `json:"workers_inactive"`
	WorkersDead     int                     `json:"workers_dead"`
	Shares1d        string                  `json:"shares_1d"`
	Shares1dUnit    string                  `json:"shares_1d_unit"`
	Shares1dPure    string                  `json:"shares_1d_pure"`
	HashrateSuffix  string                  `json:"hashrate_suffix"`
	Amount          string                  `json:"amount"`
	AmountType      string                  `json:"amount_type"`
	Name            string                  `json:"name"`
	RegionId        string                  `json:"region_id"`
	LatestAlert     SubAccountHashrateAlert `json:"latest_alert"`
}

type SubAccountHashrateAlert struct {
	Actual           string `json:"actual"`
	Expect           string `json:"expect"`
	Trans            string `json:"trans"`
	Type1            string `json:"type"`
	Unit             string `json:"unit"`
	CreateAt         string `json:"create_at"`
	CreatedTimestamp string `json:"created_timestamps"`
}

type ChangeHashrateRes struct {
	DestPuid string `json:"dest_puid"`
	DestPuidName string `json:"dest_puid_name"`
	DestRegionId string `json:"dest_region_id"`
	SwitchMode string `json:"switch_mode"`
	RegionName string `json:"region_name"`
	RegionUrl string `json:"region_url"`
}
