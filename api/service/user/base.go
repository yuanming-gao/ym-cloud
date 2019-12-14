//Package user :
// @Time : 2019/12/2 4:44 下午
// @Author : GaoYuanMing
// @Package : user
// @FileName : base.go
package user

import (
	db "api/dao/mysql"
	"api/model/dto"
	"api/model/entity"
	"api/service/auth"
	"api/tools"
	"errors"
	"log"
)

//LoginService :用户登录
func LoginService(dto *dto.SignInDto) (*entity.User, *entity.AccessToken, error) {
	exist := db.IsExistUserPhone(dto.Phone)
	if exist == false {
		return nil, nil, errors.New("未注册的手机号")
	}
	user := db.SelectUser(dto.Phone)
	if user.Password != dto.Password {
		return nil, nil, errors.New("未找到用户,用户名或密码错误")
	}
	//保存登录记录
	record := &entity.LoginRecord{
		LoginPhone: dto.Phone,
		LoginTime:  tools.NowTimeToUnixNano(),
		LoginIP:    dto.IP,
	}
	if err := db.InsertUserLoginRecord(record); err != nil {
		log.Println(err)
		return nil, nil, errors.New("登录异常,无法保存登录记录")
	}
	token, err := auth.CreateAccessToken(user.ID)
	if err != nil {
		return nil, nil, err
	}
	return user, token, nil
}

func RegisterService(dto *dto.SignUpDto) error {
	if len(dto.Phone) != 11 || len(dto.Password) < 6 || len(dto.Name) < 2 {
		return errors.New("传入的数据有误")
	}
	exist := db.IsExistUserPhone(dto.Phone)
	if exist == true {
		return errors.New("该手机号已经注册")
	}
	user := &entity.User{
		Name:       dto.Name,
		Phone:      dto.Phone,
		Password:   dto.Password,
		Position:   dto.Position,
		CreateTime: tools.NowTimeToUnixNano(),
		UpdateTime: tools.NowTimeToUnixNano(),
	}
	return db.InsertUser(user)
}

func UpdateUserInfoService() bool {
	return false
}
