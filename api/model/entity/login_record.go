//Package entity :
// @Time : 2019/12/2 3:37 下午
// @Author : GaoYuanMing
// @Package : entity
// @FileName : login_record.go
package entity

//LoginRecord :用户的登录信息
type LoginRecord struct {
	LoginIP    string
	LoginPhone string
	LoginTime  int64
}
