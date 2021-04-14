package controller

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

// GetLang ...
func GetLang(c *gin.Context) string {
	if l, exit := c.Get("lang"); exit {
		if langStr, ok := l.(string); ok {
			return langStr
		}
	}
	langStr := c.Query("lang")
	if len(langStr) == 0 {
		return "en_US"
	} else {
		return langStr
	}
}

func GetParam(c *gin.Context, key string) string {
	if v, exit := c.Get(key); exit {
		if val, ok := v.(string); ok {
			return val
		}
	}
	val := c.Query(key)
	if len(val) == 0 {
		return ""
	} else {
		return val
	}
	// post？postfrom
}

func urlEncoded(params map[string]interface{}) (string, error) {
	ue := url.Values{}
	for k, v := range params {
		str := fmt.Sprintf("%v", v)
		ue.Add(k, str)
	}
	return url.QueryUnescape(ue.Encode())
}
