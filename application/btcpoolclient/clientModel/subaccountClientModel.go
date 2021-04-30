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
