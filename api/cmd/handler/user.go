package handler

import (
	"api/cmd/common"
	"api/model/dto"
	userservice "api/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SignInHandler :登录
func SignInHandler(c *gin.Context) {
	signInDto := &dto.SignInDto{
		Phone:    c.Query("phone"),
		Password: c.Query("password"),
		IP:       c.ClientIP(),
	}
	responseBody := new(common.APIResponseBody)
	user, token, err := userservice.LoginService(signInDto)
	if err != nil {
		responseBody.Data = nil
		responseBody.Status = http.StatusNotFound
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	userInfo := &dto.UserInfoDto{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Position: user.Position,
		Avatar:   "",
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "登录成功"
	responseBody.Data = make(map[string]interface{})
	responseBody.Data["user_info"] = userInfo
	responseBody.Data["access_token"] = token.ID
	common.SendAPIResponse(c, responseBody)
}

//SignUpHandler :注册
func SignUpHandler(c *gin.Context) {
	signUpDto := new(dto.SignUpDto)
	responseBody := new(common.APIResponseBody)
	err := common.ParseAPIRequestJSON(c.Request, signUpDto)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = "Json解析错误"
		common.SendAPIResponse(c, responseBody)
		return
	}
	err = userservice.RegisterService(signUpDto)
	if err != nil {
		responseBody.Status = http.StatusBadRequest
		responseBody.Msg = err.Error()
		common.SendAPIResponse(c, responseBody)
		return
	}
	responseBody.Status = http.StatusOK
	responseBody.Msg = "注册成功"
	common.SendAPIResponse(c, responseBody)
}
