package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AppVersionCheck(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "app.versionCheck", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error  AppVersionCheck: %v", err)
	}

	return dest.Data, nil
}

func UrlConfig(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "app.urlConfig", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error UrlConfig: %v", err)
	}

	return dest.Data, nil
}

// Banner ..
type BannerList []Banner
type Banner struct {
	Id    string `json:"id"`
	Pic   string `json:"pic"`
	Link  string `json:"link"`
	Title string `json:"Title"`
	I18n  int    `json:"i18n"`
}

// GetBannerList ..
func GetBannerList(c *gin.Context, params interface{}) ([]Banner, error) {

	// fake data
	// res := make([]Banner, 1)
	// res[0] = Banner{"0", "http://xxx.png", "http://xxx.png"}
	// return res, nil

	var dest = struct {
		BtcpoolRescomm
		Data []Banner `json:"data"`
	}{}

	_, err := doRequest(c, "app.banner", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error getting banner list: %v", err)
	}

	return dest.Data, nil
}

// Notice
type NoticeList []Notice
type Notice struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`
}

// Get notice list
func GetNoticeList(c *gin.Context, params interface{}) ([]Notice, error) {
	var dest = struct {
		BtcpoolRescomm
		Data []Notice `json:"data"`
	}{}
	_, err := doRequest(c, "app.notice", params, &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error getting notice list: %v", err)
	}
	return dest.Data, nil
}

type CoinStat struct {
	Stats        CoinInnerStat  `json:"stats"`
	CoinType     string         `json:"coin_type"`
	CoinPayMode  string         `json:"coin_pay_mode"`
	CoinPayLimit string         `json:"coin_pay_limit"`
	CoinSuffix   string         `json:"coin_suffix"`
	BlocksCount  string         `json:"blocks_count"`
	RewardsCount string         `json:"rewards_count"`
	IndexCoin    IndexCoinModel `json:"index_coin"`
}
type CoinInnerStat struct {
	Shares CoinStatShares `json:"shares"`
}
type CoinStatShares struct {
	Shares15m  string `json:"shares_15m"`
	SharesUnit string `json:"shares_unit"`
}
type IndexCoinModel struct {
	Algorithm string `json:"algorithm"`
	Mode      string `json:"mode"`
	Text      string `json:"text"`
	PayLimit  string `json:"pay_limit"`
}

// ?????????????????????
func GetPoolMultiCoinStats(c *gin.Context) (map[string]CoinStat, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]CoinStat `json:"data"`
	}{}

	_, err := doRequest(c, "public.multiCoinStats", "dimension=1h&is_decimal=1&no_share_history=1", &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error multiCoinStats: %v", err)
	}
	return dest.Data, nil
}

type CoinIncomList []CoinIncome
type CoinIncome struct {
	CoinType                 string  `json:"-"`
	Hashrate                 float64 `json:"hashrate"`
	Diff                     float64 `json:"diff"`
	IncomeCoin               float64 `json:"income_coin"`
	IncomeUsd                float64 `json:"income_usd"`
	IncomeCny                float64 `json:"income_cny"`
	IncomeOptimizeCoin       float64 `json:"income_optimize_coin"`
	IncomeOptimizeUsd        float64 `json:"income_optimize_usd"`
	IncomeOptimizeCny        float64 `json:"income_optimize_cny"`
	DiffAdjustTime           string  `json:"diff_adjust_time"`
	NextDiff                 string  `json:"next_diff"`
	NextIncomeCoin           string  `json:"next_income_coin"`
	NextIncomeUsd            string  `json:"next_income_usd"`
	NextIncomeCny            string  `json:"next_income_cny"`
	PaymentMode              string  `json:"payment_mode"`
	IncomeHashrateUnit       string  `json:"income_hashrate_unit"`
	IncomeHashrateUnitSuffix string  `json:"income_hashrate_unit_suffix"`
	IncomeRealCoin           float64 `json:"income_real_coin"`
	IncomeRealUsd            float64 `json:"income_real_usd"`
	IncomeRealCny            float64 `json:"income_real_cny"`
}

// ???????????????????????????
func GetCoinIncome(c *gin.Context) (map[string](CoinIncome), error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string](CoinIncome) `json:"data"`
	}{}

	_, err := doRequest(c, "public.coinsIncome", "", &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error GetCoinIncome: %v", err)
	}
	return dest.Data, nil
}

func GetCpatcha(c *gin.Context, params interface{}, typeStr string) (map[string]interface{}, error) {
	var dest struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}

	if typeStr == "sms" {
		_, err := doRequest(c, "app.captchaSMS", params, &dest)
		if err != nil {
			return nil, err //fmt.Errorf("error GetCoinIncome: %v", err)
		}
	} else if typeStr == "email" {
		_, err := doRequest(c, "app.captchaEmail", params, &dest)
		if err != nil {
			return nil, err //fmt.Errorf("error GetCoinIncome: %v", err)
		}
	} else {
		return nil, fmt.Errorf("err GetCpatcha: wrong type")
	}
	return dest.Data, nil
}
