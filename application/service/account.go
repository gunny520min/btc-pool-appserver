package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/btcpoolclient/clientModel"
	"btc-pool-appserver/application/model"
	"github.com/gin-gonic/gin"
)

type accountHandler struct{}

var AccountService = &accountHandler{}

func (p *accountHandler) GetSubAccountList(c *gin.Context, puid string, isHidden int) ([]model.SubaccountEntity2, error) {
	var algorithmsParams = map[string]interface{}{}
	algorithmsParams["puid"] = puid
	algorithmsParams["is_hidden"] = isHidden
	algorithmsParams["is_guardian"] = 0
	algorithmsParams["order_by"] = "puid"

	var subaccountList = make([]model.SubaccountEntity2, 0)
	if subAccountAlgorithms, e := btcpoolclient.GetSubAccountAlgorithms(c, algorithmsParams); e != nil {
		return subaccountList, e
	} else {
		for _, subaccount := range subAccountAlgorithms.SubAccounts {
			for _, algorithm := range subaccount.Algorithms {
				j := 0
				var l = make([]clientModel.SubAccountCoinEntity, 0)
				for _, coinAccount := range algorithm.CoinAccounts {
					if algorithm.IsSmart() == coinAccount.IsSmart() && coinAccount.IsHidden == isHidden {
						l = append(l, coinAccount)
						//subaccount.Algorithms[j] = algorithm
						j++
					}
				}
				algorithm.CoinAccounts = l
			}
		}

		for _, account := range subAccountAlgorithms.SubAccounts {
			var subaccountEntity = model.SubaccountEntity2{}
			subaccountEntity.Title = account.Name + " " + account.RegionText
			subaccountEntity.SearchKey = account.Name + " " + account.RegionText // TODO searchKey 中文情况下可以加拼音
			subaccountEntity.SubAccount = make([]model.SubaccountAlgorithmCoinAccountEntity, 0)
			for _, algorithm := range account.Algorithms {
				//var a = model.SubaccountAlgorithmEntity{}
				//a.AlgorithmText = algorithm.AlgorithmText
				//a.CurrentCoin = algorithm.CurrentCoin
				//a.CurrentPuid = algorithm.CurrentPuid
				//a.IsSmart = algorithm.IsSmart()
				//a.SupportCoins = algorithm.SupportCoins
				//a.SubAccount = make([]model.SubaccountAlgorithmCoinAccountEntity, 0)
				for _, coinAccount := range algorithm.CoinAccounts {
					var coinA = model.SubaccountAlgorithmCoinAccountEntity{}
					coinA.CoinType = coinAccount.CoinType
					coinA.Puid = coinAccount.Puid
					coinA.IsHidden = coinAccount.IsHidden == 1
					coinA.IsCurrent = coinAccount.IsCurrent == 1
					coinA.SupportCoins = algorithm.SupportCoins

					subaccountEntity.SubAccount = append(subaccountEntity.SubAccount, coinA)
				}

				//subaccountEntity.Algorithms = append(subaccountEntity.Algorithms, a)
			}

			subaccountList = append(subaccountList, subaccountEntity)
		}

		return subaccountList, nil
	}
}
func (p *accountHandler) GetCurrentSubAccount(c *gin.Context, puid string) (*model.SubaccountEntityCurrent2, error) {
	if subaccountList, err := AccountService.GetSubAccountList(c, puid, 0); err != nil {
		return nil, err
	} else {
		for _, coinAccount := range subaccountList {
			for _, a := range coinAccount.SubAccount {
				if a.Puid == puid {
					entity := model.SubaccountEntityCurrent2{
						Puid:     a.Puid,
						CoinType: a.CoinType,
					}
					entity.SubaccountEntity2 = coinAccount
					return &entity, nil
				}
			}
		}

		entity := model.SubaccountEntityCurrent2{
			Puid:     subaccountList[0].SubAccount[0].Puid,
			CoinType: subaccountList[0].SubAccount[0].CoinType,
		}
		entity.SubaccountEntity2 = subaccountList[0]
		return &entity, nil
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
			hashrateEntity.HashrateUnit = h.Shares1dUnit + h.HashrateSuffix
			hashrateEntity.LastAlertTrans = h.LatestAlert.Trans

			data[puid] = hashrateEntity
		}

		return data, nil
	}
}

func (p *accountHandler) GetSubAccountCreateInit(c *gin.Context) ([]model.CreateSubaccountInitEntity, error) {
	var params = map[string]interface{}{}
	var data = make([]model.CreateSubaccountInitEntity, 0)
	if subaccountCreateInit, e := btcpoolclient.SubaccountCreateInit(c, params); e != nil {
		return data, e
	} else {
		coinTypeList := make([]string, 0)
		for _, coinType := range subaccountCreateInit.CoinTypeList {
			for k, _ := range subaccountCreateInit.NodeList {
				if coinType == k {
					coinTypeList = append(coinTypeList, coinType)
				}
			}
		}

		for _, key := range coinTypeList {
			item := model.CreateSubaccountInitEntity{}
			item.CoinType = key
			item.RegionNode = make([]model.CreateSubaccountInitNodeEntity, 0)
			for name, _ := range subaccountCreateInit.RegionList {
				if n, ok := subaccountCreateInit.NodeList[key]; ok {
					for _, childn := range n {
						if name == childn.Region {
							node := model.CreateSubaccountInitNodeEntity{}
							node.ShowText = childn.Text
							node.RegionName = childn.RegionName
							item.RegionNode = append(item.RegionNode, node)
						}
					}
				}
			}
			data = append(data, item)
		}

		return data, nil
	}
}
