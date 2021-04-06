package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
var id: String = ""
    var fid: String = ""
    var pic: String = ""
    var link: String = ""
    var i18n: Int?
    var title: String = ""
    var target: Bool = false
    var description: String = ""
    var begin: String = ""
    var end: String = ""
    var beginStr: String = ""
    var endStr: String = ""
    var backgroundColor: String = ""
*/

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
		return nil, fmt.Errorf("error getting banner list: %v", err)
	}

	return dest.Data, nil
}

// Notice
type NoticeList []Notice
type Notice struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

// Get notice list
func GetNoticeList(c *gin.Context, params interface{}) ([]Notice, error) {
	var dest = struct {
		BtcpoolRescomm
		Data []Notice `json:"data"`
	}{}
	_, err := doRequest(c, "app.notice", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error getting notice list: %v", err)
	}
	return dest.Data, nil
}
