package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"fmt"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type publicHandler struct{}

var PublicService = &publicHandler{}

type BannerAndNoticeData struct {
	Banner btcpoolclient.BannerList
	Notice btcpoolclient.NoticeList
}

func (p *publicHandler) GetBannerAndNotice(c *gin.Context, params interface{}) (BannerAndNoticeData, error) {

	bannerList := make([]btcpoolclient.Banner, 0)
	var eBanner error
	noticeList := make([]btcpoolclient.Notice, 0)
	var eNotice error

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		//notice
		defer wg.Done()
		var p = make(map[string]string)
		p["test"] = "1"
		noticeList, eNotice = btcpoolclient.GetNoticeList(c, p)
	}()
	go func() {
		//banner
		defer wg.Done()
		bannerList, eBanner = btcpoolclient.GetBannerList(c, params)
	}()
	wg.Wait()
	var res BannerAndNoticeData
	if eBanner != nil {
		return res, eBanner
	}
	if eNotice != nil {
		return res, eNotice
	}
	res.Banner = bannerList
	res.Notice = noticeList
	return res, nil
}

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
		n.Id = v.Id
		n.Content = v.Title
		n.Link = v.HtmlUrl
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
func (p *publicHandler) AsnycGetPoolRank(c *gin.Context, coin string, params interface{}) <-chan map[string]btcpoolclient.PoolRankData {
	ch := make(chan map[string]btcpoolclient.PoolRankData, 0)
	go func() {
		var res map[string]btcpoolclient.PoolRankData
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("AsnycGetPoolrank err %v", err))
			}
			ch <- res
		}()
		if dic, err := btcpoolclient.GetPoolRank(c, params); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			res = dic
			// res = dic[strings.ToLower(coin)].Realtime.List
		}
	}()
	return ch
}

func (p *publicHandler) FormatPoolRankList(rankData map[string]btcpoolclient.PoolRankData) map[string]([]model.PoolRank) {
	res := make(map[string]([]model.PoolRank), 0)

	for k, v := range rankData {
		ranks := make([]model.PoolRank, 0)
		for _, rankInfo := range v.Realtime.List {
			ranks = append(ranks, model.PoolRank{
				PoolName:               rankInfo.PoolName,
				IconLink:               rankInfo.IconLink,
				RealtimeHashrate:       rankInfo.RealtimeHashrate,
				EstimateHashrate:       rankInfo.EstimateHashrate,
				RealtimeCur2maxPercent: rankInfo.RealtimeCur2maxPercent,
				EstimateCur2max:        rankInfo.EstimateCur2max,
				HashSuffix:             rankInfo.HashSuffix,
				RealtimeDiff24hPercent: rankInfo.RealtimeDiff24hPercent,
			})
		}
		res[k] = ranks
	}
	return res
}

// get latest block
func (p *publicHandler) AsnycGetLatestBlocks(c *gin.Context, coin string, params interface{}) <-chan (map[string]btcpoolclient.LatestBlockList) {
	ch := make(chan (map[string]btcpoolclient.LatestBlockList), 0)
	go func() {
		var res (map[string]btcpoolclient.LatestBlockList)
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("AsnycGetLatestBlocks err %v", err))
			}
			ch <- res
		}()
		if dic, err := btcpoolclient.GetLatestBlockList(c, params); err != nil {
			_ = c.Error(err).SetType(gin.ErrorTypeNu)
		} else {
			res = dic
		}
	}()
	return ch
}

func (p *publicHandler) FormatLatestBlockList(blkListInfo map[string]btcpoolclient.LatestBlockList) map[string]([]model.LatestBlock) {

	res := make(map[string]([]model.LatestBlock), 0)
	blkArr := make([]model.LatestBlock, 0)
	for k, v := range blkListInfo {
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
		res[k] = blkArr
	}
	return res
}
