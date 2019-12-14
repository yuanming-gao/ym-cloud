//Package auth
//@ Author: Gao YuanMing
//@ Data: 2019/12/9 10:05 下午
//@ Description:检验token

package auth

import (
	db "api/dao/mysql"
	"api/model/entity"
	"api/tools"
	"errors"
	"log"
)

//CreateAccessToken :身份验证后的token创建函数
func CreateAccessToken(userID int) (*entity.AccessToken, error) {
	token := entity.NewAccessToken(userID)
	err := db.InsertAccessToken(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return token, nil
}

//CheckAccessToken : 检验token是否可用
func CheckAccessToken(tokenID string, userID int) error {
	token := db.RetrieveAccessToken(tokenID)
	if token == nil {
		return errors.New("不存在的tokenID")
	}
	//id不匹配直接退出
	if token.UserID != userID {
		return errors.New("tokenID与userID不匹配")
	}
	//token过期
	if token.ExpirationTime < tools.NowTimeToUnixNano() {
		//删除数据库的token
		err := db.DeleteAccessToken(token)
		if err != nil {
			log.Println(err)
		}
		return errors.New("token已过期请重新获取")
	}
	return nil
}
