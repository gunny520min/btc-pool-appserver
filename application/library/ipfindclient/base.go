package ipfindclient

import (
	"btc-pool-appserver/application/config"
	"btc-pool-appserver/application/library/third"
	"fmt"
	"net/url"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

type IpfindResComm struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
	Msg    string `json:"message"`
}

func (b IpfindResComm) GetCode() int {
	return b.Errno
}

func (b IpfindResComm) GetMessage() string {
	if b.Msg != "" {
		return b.Msg
	}
	return b.Errmsg
}

func (b IpfindResComm) GetMessageSuffix() string {
	return "ipfind"
}

// 判断code是否表示调用成功, 如果不成功会直接返回错误
func (b IpfindResComm) IsSucc() bool {
	if b.Errno != 0 || b.Msg != "" {
		return false
	}
	return true
}

func QueryIp(c *gin.Context, ip string) (*simplejson.Json, error) {
	if ip == "" {
		return nil, fmt.Errorf("ipfind query ip: empty ip")
	}

	params := url.Values{
		"ip": []string{ip},
	}

	if j, err := DoRequestWithSjRes(c, params.Encode(), "query"); err != nil {
		return nil, fmt.Errorf("ipfind query ip: %w", err)
	} else {
		return j, nil
	}
}

func DoRequestWithSjRes(c *gin.Context, params interface{}, action string) (*simplejson.Json, error) {
	dest := new(IpfindResComm) // 用于检查返回错误码是否为0
	res, err := doRequest(c, action, params, dest)
	if err != nil {
		return nil, err
	}

	j, err := simplejson.NewJson(res)
	if err != nil {
		return nil, fmt.Errorf("simple json fail: %w", err)
	}

	return j.Get("data"), nil
}

// 发起一个ipfind接口请求
func doRequest(c *gin.Context, action string, params interface{}, dest interface{}) ([]byte, error) {
	// 根据action获取请求配置
	apiConfig, exist := config.Third.Ipfind.Apis[action]
	if !exist {
		return nil, fmt.Errorf("ipfind dorequest unknown action: %s", action)
	}

	// 发起http请求
	if res, err := third.DoActionRequest(c, &apiConfig, params, nil, dest); err != nil {
		return nil, fmt.Errorf("ipfind: %w", err)
	} else {
		return res, nil
	}
}

func GetIpCountryName(sj *simplejson.Json, ip string, languages []string) string {
	if sj != nil {
		names := sj.Get(ip).Get("Country").Get("Names")
		if names != nil {
			for _, l := range languages {
				n, err := names.Get(l).String()
				if err == nil && n != "" {
					return n
				}
			}
		}
	}

	return ""
}

func GetIpCountryCode(sj *simplejson.Json, ip string) string {
	if sj != nil {
		return sj.Get(ip).Get("Country").Get("IsoCode").MustString()
	}

	return ""
}
