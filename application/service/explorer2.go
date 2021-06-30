package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/library/tool"
	"btc-pool-appserver/application/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type explorerHandler struct {}

var ExplorerService = &explorerHandler{}

// GetPoolRank 矿池排名
func (p *explorerHandler) GetPoolRank(c *gin.Context, coin string) ([]model.PoolRankInfo, error) {
	platform := c.GetHeader("platform")
	if len(platform) == 0 {
		platform = ""
	}
	param := map[string]string{
		"coins": coin,
		"from":  platform,
	}
	if dic, err := btcpoolclient.GetPoolRank(c, param); err != nil {
		return []model.PoolRankInfo{}, err
	} else {
		// get btc 's pool rank
		v := dic[coin]
		ranks := make([]model.PoolRankInfo, 0)
		for i, rankInfo := range v.Realtime.List {
			var hash string
			var unit string
			if len(rankInfo.RealtimeHashrate) == 0 {
				hash = tool.KeepStringNum(rankInfo.EstimateHashrate, 2)
				unit = rankInfo.EstimateHashrateUnit + rankInfo.HashrateSuffix
			} else {
				hash = tool.KeepStringNum(rankInfo.RealtimeHashrate, 2)
				unit = rankInfo.HashrateUnit + rankInfo.HashrateSuffix
			}

			var diff string
			if len(rankInfo.RealtimeHashrate) == 0 {
				diff = "-"
			} else {
				diff = tool.KeepStringNum(rankInfo.RealtimeDiff24hPercent, 2)
			}

			ranks = append(ranks, model.PoolRankInfo{
				Index:                 strconv.Itoa(i + 1),
				Name:                  rankInfo.PoolName,
				Icon:                  rankInfo.IconLink,
				HashratePercent:       rankInfo.Cur2maxPercent,
				Hashrate:              hash,
				HashrateUnit:          unit,
				HashrateChangePercent: diff,
				Lucy7Day:              rankInfo.Lucky3d,
			})
		}

		return ranks, nil
	}
}
//func (p *explorerHandler) GetPoolRank(c *gin.Context, params map[string]string) (map[string][]model.PoolRank, error) {
//	res := make(map[string][]model.PoolRank, 0)
//	if dic, err := btcpoolclient.GetPoolRank(c, params); err != nil {
//		return res, err
//	} else {
//		for k, v := range dic {
//			ranks := make([]model.PoolRank, 0)
//			for i, rankInfo := range v.Realtime.List {
//
//				var hash string
//				if len(rankInfo.RealtimeHashrate) == 0 {
//					hash = tool.KeepStringNum(rankInfo.EstimateHashrate, 2)
//				} else {
//					hash = tool.KeepStringNum(rankInfo.RealtimeHashrate, 2)
//				}
//
//				var diff string
//				if len(rankInfo.RealtimeHashrate) == 0 {
//					diff = "-"
//				} else {
//					diff = tool.KeepStringNum(rankInfo.RealtimeDiff24hPercent, 2)
//				}
//
//				ranks = append(ranks, model.PoolRank{
//					Rank:                   strconv.Itoa(i + 1),
//					PoolName:               rankInfo.PoolName,
//					IconLink:               rankInfo.IconLink,
//					Progress:				rankInfo.Cur2maxPercent,
//					Hashrate:               hash,
//					Diff: diff,
//				})
//			}
//			res[k] = ranks
//		}
//		return res, nil
//	}
//}

// GetLatestBlocks 最近出块
func (p *explorerHandler) GetLatestBlocks(c *gin.Context, coin string) ([]model.LatestBlock, error) {
	params := map[string]string{
		"coins": coin,
		"show_unconfirm_info":  "true",
	}
	data := make(map[string][]model.LatestBlock)
	if res, err := btcpoolclient.GetLatestBlockList(c, params); err != nil {
		return []model.LatestBlock{}, err
	} else {
		for k, v := range res {
			blkArr := make([]model.LatestBlock, 0)
			for _, blkInfo := range v {
				blkArr = append(blkArr, model.LatestBlock{
					Timestamp: blkInfo.Timestamp,
					Reward:    blkInfo.Reward,
					Height:    blkInfo.Height,
					PoolName:  blkInfo.PoolName,
					Hash:      blkInfo.Hash,
					Size:      blkInfo.Size,
				})
			}
			data[k] = blkArr
		}
		return data[coin], nil
	}
}