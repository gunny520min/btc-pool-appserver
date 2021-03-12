package third

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"

	"btc-pool-appserver/application/config"
	"btc-pool-appserver/application/library/errs"
	"btc-pool-appserver/application/library/log"
	// "gitlab.bitdeer.vip/bitdeer/go-lib/okHttp"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

type Responser interface {
	GetCode() int
	GetMessage() string
	IsSucc() bool
	GetMessageSuffix() string
}

func DoActionRequest(c *gin.Context, api *config.Api, params interface{}, headers map[string]string, dest interface{}) ([]byte, error) {
	var res []byte
	var err error
	var ext map[string]string

	var lang string
	if headers == nil {
		headers = make(map[string]string)
	}
	if c != nil {
		headers["Cookie"] = c.GetHeader("Cookie")
		langV, exists := c.Get("lang")
		if exists {
			lang = langV.(string)
		}
	}

	var finalParams string
	switch p := params.(type) {
	case string:
		finalParams = p
	case []byte:
		finalParams = string(p)
	default:
		pbyte, _ := json.Marshal(p)
		finalParams = string(pbyte)
	}

	if api.Method == "POST" {
		headers["Content-Type"] = "application/json"
		// res, ext, err = okHttp.Post(api.Uri, finalParams, api.Timeout, 3, headers)
	} else {
		// res, ext, err = okHttp.Get(api.Uri+"?"+finalParams, api.Timeout, 3, headers)
	}

	if err != nil {
		return res, fmt.Errorf("third request fail: %w", err)
	}

	if c != nil {
		if _, exist := c.Get("_secret"); exist { // 如果是敏感接口，就不记录请求参数和返回结果了，以免密码等泄漏
			log.ContextInfo(c, "third request secret: ", api, " ext:", ext)
		} else {
			log.ContextWithFields(c, "call_third", map[string]interface{}{
				"api":    api,
				"ext":    ext,
				"params": finalParams,
			})
		}
	}

	if dest != nil {
		fmt.Println("222222222222222222")
		fmt.Println(string(res))
		if err = jsoniter.Unmarshal(res, dest); err != nil {
			return nil, err
		}

		if responser, ok := dest.(Responser); ok {

			if !responser.IsSucc() {
				// 这里需要返回一个可以直接输出的err, 所以把api err包进去
				return nil, fmt.Errorf("third res fail: code %v, msg: %s| %w", responser.GetCode(), responser.GetMessage(), errs.NewApiErrorThird(responser.GetCode(), responser.GetMessage(), nil, lang, responser.GetMessageSuffix()))
			} else {
				switch responser.GetCode() {
				case errs.ErrnoNeedRefreshToken:
					if c != nil {
						c.Set("_forceOutputCode", errs.ErrnoNeedRefreshToken)
					}
				}
			}
		}
	}

	return res, nil
}
