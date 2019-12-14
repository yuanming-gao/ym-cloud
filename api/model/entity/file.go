//Package entity
// @Time : 2019/11/26 8:29 下午
// @Author : GaoYuanMing
// @Package : entity
// @FileName : file.go

package entity

//File :
type File struct {
	Name         string
	Size         int64 //文件大小
	Status       bool  //文件状态是否失效
	Path         string
	Sha256String string
	CreateTime   int64
	UpdateTime   int64
}
