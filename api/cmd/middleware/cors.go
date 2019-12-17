package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Cors :允许跨域访问
func Cors(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,X-UserID,X-AccessToken, x-requested-with, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE,PUT,OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
