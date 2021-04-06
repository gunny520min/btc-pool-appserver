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

func urlEncoded(params map[string]interface{}) (string, error) {
	ue := url.Values{}
	for k, v := range params {
		str := fmt.Sprintf("%v", v)
		ue.Add(k, str)
	}
	return url.QueryUnescape(ue.Encode())
}
