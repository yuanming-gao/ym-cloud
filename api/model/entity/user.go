//Package entity
// @Time : 2019/11/26 8:29 下午
// @Author : GaoYuanMing
// @Package : entity
// @FileName : file.go

package entity

//User :用户
type User struct {
	ID         int
	Name       string
	Phone      string
	Password   string
	Position   string //所在城市位置
	CreateTime int64
	UpdateTime int64
}
