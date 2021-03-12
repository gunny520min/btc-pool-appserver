package controller

import "github.com/gin-gonic/gin"

// GetLang ...
func GetLang(c *gin.Context) string {
	if l, exit := c.Get("lang"); exit {
		if langStr, ok := l.(string); ok {
			return langStr
		}
	}
	return "en_US"
}
