//Package middleware :
// @Time : 2019/11/26 11:27 上午
// @Author : GaoYuanMing
// @Package : middleware
// @FileName : authUser.go
package middleware

import (
	"api/cmd/common"
	"api/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Auth :
func Auth(c *gin.Context) {
	token := c.Request.Header.Get("X-AccessToken")
	userID, err := strconv.Atoi(c.Request.Header.Get("X-UserID"))
	if err != nil {
		responseBody := new(common.APIResponseBody)
		responseBody.Msg = "请求参数缺失"
		responseBody.Status = http.StatusBadRequest
		common.SendAPIResponse(c, responseBody)
		c.Abort()
		return
	}
	//校验token
	err = auth.CheckAccessToken(token, userID)
	if err != nil {
		responseBody := new(common.APIResponseBody)
		responseBody.Msg = err.Error()
		responseBody.Status = http.StatusForbidden
		common.SendAPIResponse(c, responseBody)
		c.Abort()
	}
	c.Next()
}
