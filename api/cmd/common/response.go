//Package handler
//@ Author: Gao YuanMing
//@ Data: 2019/12/10 9:41 上午
//@ Description:

package common

import (
	"github.com/gin-gonic/gin"
	"log"
)

type APIResponseBody struct {
	Status int                    `json:"status"` //http状态码
	Msg    string                 `json:"msg"`
	Data   map[string]interface{} `json:"data"`
}

func SendAPIResponse(c *gin.Context, responseBody *APIResponseBody) {
	data, err := json.Marshal(responseBody)
	if err != nil {
		log.Println(err)
		_, _ = c.Writer.WriteString("json解析错误")
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	_, err = c.Writer.Write(data)
	if err != nil {
		c.Writer.WriteHeader(500)
		log.Println(err)
		return
	}
}
