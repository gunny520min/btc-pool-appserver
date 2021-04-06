package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"fmt"

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
