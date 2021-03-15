package btcpoolclient

import (
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

type BannerList []Banner

// Banner ..
type Banner struct {
	Id   string `json:"id"`
	Pic  string `json:"pic"`
	Link string `json:"link"`
}

// GetBannerList ..
func GetBannerList(c *gin.Context, params interface{}) ([]Banner, error) {

	// fake data
	res := make([]Banner, 1)
	res[0] = Banner{"0", "http://xxx.png", "http://xxx.png"}
	return res, nil

	// var dest = struct {
	// 	BtcpoolRescomm
	// 	Data []Banner `json:"data"`
	// }{}

	// _, err := doRequest(c, "bannerlist", params, &dest)
	// if err != nil {
	// 	return nil, fmt.Errorf("banner list: %w", err)
	// }

	// return dest.Data, nil
}
