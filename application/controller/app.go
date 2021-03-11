package controller

import "github.com/gin-gonic/gin"

// AppVersion 获取升级包信息
func AppVersion(c *gin.Context) {
	//TODO
	c.JSON(200, gin.H{"msg": "todo"})
}
