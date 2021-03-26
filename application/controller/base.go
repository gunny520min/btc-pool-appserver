package controller

import "github.com/gin-gonic/gin"

// GetLang ...
func GetLang(c *gin.Context) string {
	if l, exit := c.Get("lang"); exit {
		if langStr, ok := l.(string); ok {
			return langStr
		}
	}
	langStr := c.Query("lang")
	if len(langStr)==0 {
		return "en_US"
	} else {
		return langStr
	}
}
