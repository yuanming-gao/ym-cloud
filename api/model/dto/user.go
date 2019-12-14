//Package dto :
/**
* @Time : 2019/11/26 8:29 下午
* @Author : GaoYuanMing
* @Package : dto
* @FileName : user_dto.go
 */
package dto

//UserInfoDto :用户信息
type UserInfoDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
	Avatar   string `json:"avatar"`
}

//SignInDto :登录需要的传输数据
type SignInDto struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	IP       string
}

//SignUpDto :注册需要的传输数据
type SignUpDto struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

//LoginRecordDto :登录记录
type LoginRecordDto struct {
	LoginPhone string `json:"login_phone"`
	LoginIP    string `json:"login_ip"`
	LoginTime  string `json:"login_time"`
}
