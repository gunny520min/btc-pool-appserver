package controller

import "github.com/gin-gonic/gin"

// HomeIndex 首页数据
func HomeIndex(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Todo"})
}
