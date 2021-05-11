package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/model"
	"github.com/gin-gonic/gin"
)

type accountHandler struct{}

var AccountService = &accountHandler{}

func (p *accountHandler) GetSubAccountList(c *gin.Context, puid string, isHidden int) ([]model.SubaccountEntity, error) {
	var algorithmsParams = map[string]interface{}{}
	algorithmsParams["puid"] = puid
	algorithmsParams["is_hidden"] = isHidden
	algorithmsParams["is_guardian"] = 0
	algorithmsParams["order_by"] = "puid"

	var subaccountList = make([]model.SubaccountEntity, 0)
	if subAccountAlgorithms, e := btcpoolclient.GetSubAccountAlgorithms(c, algorithmsParams); e != nil {
		return subaccountList, e
	} else {
		for _, subaccount := range subAccountAlgorithms.SubAccounts {
			for _, algorithm := range subaccount.Algorithms {
				j := 0
				var l = make([]clientModel.SubAccountCoinEntity, 0)
				for _, coinAccount := range algorithm.CoinAccounts {
					if algorithm.IsSmart() == coinAccount.IsSmart() && coinAccount.IsHidden == isHidden{
						l = append(l, coinAccount)
						//subaccount.Algorithms[j] = algorithm
						j++
					}
				}
				algorithm.CoinAccounts = l
			}
		}

		for _, account := range subAccountAlgorithms.SubAccounts {
			var subaccountEntity = model.SubaccountEntity{}
			subaccountEntity.Title = account.Name+" "+account.RegionText
			subaccountEntity.SearchKey = account.Name+" "+account.RegionText // TODO searchKey 中文情况下可以加拼音
			subaccountEntity.Algorithms = make([]model.SubaccountAlgorithmEntity, 0)
			for _, algorithm := range account.Algorithms {
				var a = model.SubaccountAlgorithmEntity{}
				a.AlgorithmText = algorithm.AlgorithmText
				a.CurrentCoin = algorithm.CurrentCoin
				a.CurrentPuid = algorithm.CurrentPuid
				a.IsSmart = algorithm.IsSmart()
				a.SupportCoins = algorithm.SupportCoins
				a.SubAccount = make([]model.SubaccountAlgorithmCoinAccountEntity, 0)
				for _, coinAccount := range algorithm.CoinAccounts {
					var coinA = model.SubaccountAlgorithmCoinAccountEntity{}
					coinA.CoinType = coinAccount.CoinType
					coinA.Puid = coinAccount.Puid
					coinA.IsHidden = coinAccount.IsHidden == 1
					coinA.IsCurrent = coinAccount.IsCurrent == 1

					a.SubAccount = append(a.SubAccount, coinA)
				}

				subaccountEntity.Algorithms = append(subaccountEntity.Algorithms, a)
			}

			subaccountList = append(subaccountList, subaccountEntity)
		}

		return subaccountList, nil
	}
}

func (p *accountHandler) GetSubAccountHashrates(c *gin.Context, puids string) (map[string]model.SubaccountHashrateEntity, error) {
	var params = map[string]interface{}{}
	params["puids"] = puids
	var data = make(map[string]model.SubaccountHashrateEntity)
	if subAccountHashrate, e := btcpoolclient.GetSubaccountHashrate(c, params); e != nil {
		return nil, e
	} else {
		for puid, h := range subAccountHashrate {
			var hashrateEntity = model.SubaccountHashrateEntity{}
			hashrateEntity.Puid = h.Puid
			hashrateEntity.WorkerTotal = h.Workers
			hashrateEntity.WorkerActive = h.WorkersActive
			hashrateEntity.Hashrate = h.Shares1d
			hashrateEntity.HashrateUnit = h.Shares1dUnit+h.HashrateSuffix
			hashrateEntity.LastAlertTrans = h.LatestAlert.Trans

			data[puid] = hashrateEntity
		}

		return data, nil
	}
}
