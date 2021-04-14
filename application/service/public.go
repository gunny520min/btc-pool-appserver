package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type publicHandler struct{}

var PublicService = &publicHandler{}

func (p *publicHandler) AsyncGetNotice(c *gin.Context, params interface{}) <-chan btcpoolclient.NoticeList {
	ch := make(chan btcpoolclient.NoticeList)
	go func() {
		var noticeList btcpoolclient.NoticeList
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("%v", err))
			}
			ch <- noticeList
		}()
		if list, err := btcpoolclient.GetNoticeList(c, params); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			noticeList = list
		}
	}()
	return ch
}

func (p *publicHandler) AsyncGetBanner(c *gin.Context, params interface{}) <-chan btcpoolclient.BannerList {
	ch := make(chan btcpoolclient.BannerList, 1)
	go func() {
		var bannerlist btcpoolclient.BannerList
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("get banner list async panic: %v", err))
			}

			ch <- bannerlist
		}()
		if list, err := btcpoolclient.GetBannerList(c, params); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			bannerlist = list
		}
	}()

	return ch
}

func (p *publicHandler) FormatBannerList(ads btcpoolclient.BannerList, lang string) []model.Banner {
	var banners = make([]model.Banner, 0)
	if len(ads) <= 0 {
		return banners
	}

	// ad := ads[0]
	for _, v := range ads {
		var b model.Banner
		b.Id = v.Id
		b.ImgUrl = v.Pic
		b.Link = v.Link
		banners = append(banners, b)
	}
	return banners
}

func (p *publicHandler) FormatNoticeList(list btcpoolclient.NoticeList, lang string) []model.Notice {
	var noticeList = make([]model.Notice, 0)
	for _, v := range list {
		var n model.Notice
		n.Content = v.Title
		n.Link = v.Url
		noticeList = append(noticeList, n)
	}
	return noticeList
}

// 获取全网币收益
func (p *publicHandler) AsyncGetAllCoinIncome(c *gin.Context) <-chan btcpoolclient.CoinIncomList {
	ch := make(chan btcpoolclient.CoinIncomList)
	go func() {
		incomeList := make(btcpoolclient.CoinIncomList, 0)
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("get all coin income async panic: %v", err))
			}
			ch <- incomeList
		}()
		if dic, err := btcpoolclient.GetCoinIncome(c); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			for k, info := range dic {
				info.CoinType = k
				incomeList = append(incomeList, info)
			}
		}
	}()
	return ch
}

// 获取多币种信息
func (p *publicHandler) AsnycGetMultiCoinStats(c *gin.Context) <-chan (map[string](btcpoolclient.CoinStat)) {
	ch := make(chan map[string](btcpoolclient.CoinStat))
	go func() {
		var res map[string](btcpoolclient.CoinStat)
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("get multicoin_stats async panic %v", err))
			}
			ch <- res
		}()
		if info, err := btcpoolclient.GetPoolMultiCoinStats(c); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			res = info
		}
	}()
	return ch
}

func (p *publicHandler) FormatHomeCoinList(mulStats map[string](btcpoolclient.CoinStat), incomeList btcpoolclient.CoinIncomList) []*model.HomeCoinInfo {
	var ret []*model.HomeCoinInfo = make([]*model.HomeCoinInfo, 0)
	for k, stats := range mulStats {
		for _, income := range incomeList {
			if strings.ToLower(k) == strings.ToLower(income.CoinType) {
				hci := new(model.HomeCoinInfo)
				hci.Coin = income.CoinType
				hci.SetData(stats, income)
				ret = append(ret, hci)
				break
			}
		}
	}
	return ret
}

// get pool rank
func (p *publicHandler) AsnycGetPoolRank(c *gin.Context, coin string) <-chan btcpoolclient.PoolRankList {
	ch := make(chan btcpoolclient.PoolRankList, 0)
	go func() {
		var res btcpoolclient.PoolRankList
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("AsnycGetPoolrank err %v", err))
			}
			ch <- res
		}()
		if dic, err := btcpoolclient.GetPoolRank(c); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			res = dic[strings.ToLower(coin)].Realtime.List
		}
	}()
	return ch
}

/*
type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"poolName"`
	Hash      string `json:"hash"`
	Size      int    `json:"size"`
}*/
func (p *publicHandler) FormatPoolRankList(params btcpoolclient.PoolRankList) []model.PoolRank {
	res := make([]model.PoolRank, 0)
	for _, v := range params {
		res = append(res, model.PoolRank{
			PoolName:               v.PoolName,
			IconLink:               v.IconLink,
			RealtimeHashrate:       v.RealtimeHashrate,
			EstimateHashrate:       v.EstimateHashrate,
			RealtimeCur2maxPercent: v.RealtimeCur2maxPercent,
			EstimateCur2max:        v.EstimateCur2max,
			HashSuffix:             v.HashSuffix,
			RealtimeDiff24hPercent: v.RealtimeDiff24hPercent,
		})
	}
	return res
}

// get latest block
func (p *publicHandler) AsnycGetLatestBlocks(c *gin.Context, coin string) <-chan btcpoolclient.LatestBlockList {
	ch := make(chan btcpoolclient.LatestBlockList, 0)
	go func() {
		var res btcpoolclient.LatestBlockList
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("AsnycGetLatestBlocks err %v", err))
			}
			ch <- res
		}()
		if dic, err := btcpoolclient.GetLatestBlockList(c); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			res = dic[strings.ToLower(coin)].List
		}
	}()
	return ch
}

func (p *publicHandler) FormatLatestBlockList(params btcpoolclient.LatestBlockList) []model.LatestBlock {
	res := make([]model.LatestBlock, 0)
	for _, v := range params {
		res = append(res, model.LatestBlock{
			Timestamp: v.Timestamp,
			Reward:    v.Reward,
			Height:    v.Height,
			PoolName:  v.PoolName,
			Hash:      v.Hash,
			Size:      v.Size,
		})
	}
	return res
}
