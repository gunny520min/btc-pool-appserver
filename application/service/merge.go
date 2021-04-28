package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"

	"github.com/gin-gonic/gin"
)

type mergeHandler struct{}

var MergeService = &mergeHandler{}

func (p *mergeHandler) GetMergeCoinData(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	if d, err := btcpoolclient.GetAccountInfo(c, params); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

func (p *mergeHandler) FormatMergeCoinInfos(data map[string]interface{}, lang string) map[string]*model.MergeCoin {
	var res = make(map[string]*model.MergeCoin)
	//
	coins := data["merge_mining_coins"].(map[string]string)
	coinConfig := data["merge_mining_coins_config"].(map[string](map[string]string))
	addressInfo := data["merge_mining_addresses"].(map[string]string)

	for k, v := range coins {

		info := model.MergeCoin{}
		info.CoinType = k
		info.Offline = v == "offline"
		info.HelpUrl = coinConfig[k]["helper_link"]
		info.Address = addressInfo[k]
		// TODO: localized
		info.Title = k
		info.HelpTitle = k
		res[k] = &info
	}
	return res
}
