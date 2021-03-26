package service

import (
	"btc-pool-appserver/application/btcpoolclient"
	"btc-pool-appserver/application/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type publicHandler struct{}

var PublicService = &publicHandler{}

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
		if lang=="en_US" {
			if v.I18n==2 {
				var b model.Banner
				b.Id = v.Id
				b.Link = v.Link
				b.Lang = "en_US"
				b.Title = v.Title
				banners = append(banners, b)
			}
		} else {
			if v.I18n==1 {
				var b model.Banner
				b.Id = v.Id
				b.Link = v.Link
				b.Lang = "zh_CN"
				b.Title = v.Title
				banners = append(banners, b)
			}
		}

	}
	return banners
}
