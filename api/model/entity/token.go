package entity

import (
	"api/config"
	"api/tools"
)

//AccessToken :访问api的令牌
type AccessToken struct {
	ID             string //随机的uuid
	UserID         int
	ExpirationTime int64 //过期时间的时间戳
}

//NewAccessToken :api访问token
func NewAccessToken(userId int) *AccessToken {
	return &AccessToken{
		ID:             tools.CreateUUID(),
		UserID:         userId,
		ExpirationTime: tools.NowTimeToUnixNano() + config.AccessTokenExpirationTime,
	}
}
