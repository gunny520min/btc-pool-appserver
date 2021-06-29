package service
//
//import (
//	"btc-pool-appserver/application/btcpoolclient"
//	"btc-pool-appserver/application/library/tool"
//	"btc-pool-appserver/application/model"
//	"fmt"
//	"sort"
//	"strconv"
//	"strings"
//	"sync"
//
//	"github.com/gin-gonic/gin"
//)
//
//type publicHandler struct{}
//
//var PublicService = &publicHandler{}
//
//type BannerAndNoticeData struct {
//	Banner btcpoolclient.BannerList
//	Notice btcpoolclient.NoticeList
//}
//
//func (p *publicHandler) GetBannerAndNotice(c *gin.Context, params interface{}) (BannerAndNoticeData, error) {
//
//	bannerList := make([]btcpoolclient.Banner, 0)
//	var eBanner error
//	noticeList := make([]btcpoolclient.Notice, 0)
//	var eNotice error
//
//	wg := &sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		//notice
//		defer wg.Done()
//		var p = make(map[string]string)
//		noticeList, eNotice = btcpoolclient.GetNoticeList(c, p)
//	}()
//	go func() {
//		//banner
//		defer wg.Done()
//		bannerList, eBanner = btcpoolclient.GetBannerList(c, params)
//	}()
//	wg.Wait()
//	var res BannerAndNoticeData
//	if eBanner != nil {
//		return res, eBanner
//	}
//	if eNotice != nil {
//		return res, eNotice
//	}
//	res.Banner = bannerList
//	res.Notice = noticeList
//	return res, nil
//}
//
//func (p *publicHandler) AsyncGetNotice(c *gin.Context, params interface{}) <-chan btcpoolclient.NoticeList {
//	ch := make(chan btcpoolclient.NoticeList)
//	go func() {
//		var noticeList btcpoolclient.NoticeList
//		defer func() {
//			if err := recover(); err != nil {
//				_ = c.Error(fmt.Errorf("%v", err))
//			}
//			ch <- noticeList
//		}()
//		if list, err := btcpoolclient.GetNoticeList(c, params); err != nil {
//			_ = c.Error(err).SetType(gin.ErrorTypeNu)
//		} else {
//			noticeList = list
//		}
//	}()
//	return ch
//}
//
//func (p *publicHandler) AsyncGetBanner(c *gin.Context, params interface{}) <-chan btcpoolclient.BannerList {
//	ch := make(chan btcpoolclient.BannerList, 1)
//	go func() {
//		var bannerlist btcpoolclient.BannerList
//		defer func() {
//			if err := recover(); err != nil {
//				_ = c.Error(fmt.Errorf("get banner list async panic: %v", err))
//			}
//
//			ch <- bannerlist
//		}()
//		if list, err := btcpoolclient.GetBannerList(c, params); err != nil {
//			_ = c.Error(err).SetType(gin.ErrorTypeNu)
//		} else {
//			bannerlist = list
//		}
//	}()
//
//	return ch
//}
//
//func (p *publicHandler) FormatBannerList(ads btcpoolclient.BannerList, lang string) []model.Banner {
//	var banners = make([]model.Banner, 0)
//	if len(ads) <= 0 {
//		return banners
//	}
//
//	// ad := ads[0]
//	for _, v := range ads {
//		var b model.Banner
//		b.Id = v.Id
//		b.ImgUrl = v.Pic
//		b.Url = v.Link
//		banners = append(banners, b)
//	}
//	return banners
//}
//
//func (p *publicHandler) FormatNoticeList(list btcpoolclient.NoticeList, lang string) []model.Notice {
//	var noticeList = make([]model.Notice, 0)
//	for _, v := range list {
//		var n model.Notice
//		n.Id = v.Id
//		n.Content = v.Title
//		n.Url = v.HtmlUrl
//		noticeList = append(noticeList, n)
//	}
//	return noticeList
//}
//
//// GetAllCoinIncome  获取全网币收益
//func (p *publicHandler) GetAllCoinIncome(c *gin.Context) (btcpoolclient.CoinIncomList, error) {
//
//	incomeList := make(btcpoolclient.CoinIncomList, 0)
//	if dic, err := btcpoolclient.GetCoinIncome(c); err != nil {
//		return incomeList, err
//	} else {
//		for k, info := range dic {
//			info.CoinType = k
//			incomeList = append(incomeList, info)
//		}
//		return incomeList, nil
//	}
//
//}
//
//// GetMultiCoinStats 获取多币种信息
//func (p *publicHandler) GetMultiCoinStats(c *gin.Context) (map[string]btcpoolclient.CoinStat, error) {
//	return btcpoolclient.GetPoolMultiCoinStats(c)
//}
//
//func (p *publicHandler) FormatHomeCoinList(mulStats map[string]btcpoolclient.CoinStat, incomeList btcpoolclient.CoinIncomList) []*model.HomeCoinInfo {
//	var ret = make([]*model.HomeCoinInfo, 0)
//	for k, stats := range mulStats {
//		for _, income := range incomeList {
//			if strings.ToLower(k) == strings.ToLower(income.CoinType) {
//				hci := new(model.HomeCoinInfo)
//				hci.Coin = income.CoinType
//				hci.SetData(stats, income)
//				ret = append(ret, hci)
//				break
//			}
//		}
//	}
//	sort.SliceStable(ret, func(i, j int) bool {
//		if ret[i].Coin == "BTC" {
//			return true
//		} else if ret[j].Coin == "BTC" {
//			return false
//		} else {
//			return ret[i].Coin < ret[j].Coin
//		}
//	})
//	return ret
//}
//
//// GetPoolRank get pool rank
//func (p *publicHandler) GetPoolRank(c *gin.Context, params map[string]string) (map[string][]model.PoolRank, error) {
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
//
//// get latest block
//func (p *publicHandler) GetLatestBlocks(c *gin.Context, params map[string]string) (map[string][]model.LatestBlock, error) {
//	data := make(map[string][]model.LatestBlock)
//	if res, err := btcpoolclient.GetLatestBlockList(c, params); err != nil {
//		return data, err
//	} else {
//		for k, v := range res {
//			blkArr := make([]model.LatestBlock, 0)
//			for _, blkInfo := range v {
//				blkArr = append(blkArr, model.LatestBlock{
//					Timestamp: blkInfo.Timestamp,
//					Reward:    blkInfo.Reward,
//					Height:    blkInfo.Height,
//					PoolName:  blkInfo.PoolName,
//					Hash:      blkInfo.Hash,
//					Size:      blkInfo.Size,
//				})
//			}
//			data[k] = blkArr
//		}
//		return data, nil
//	}
//}
