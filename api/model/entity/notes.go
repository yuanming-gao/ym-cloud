//Package entity
// @Time : 2019/11/26 8:29 下午
// @Author : GaoYuanMing
// @Package : entity
// @FileName : file.go

package entity

//Notes :笔记
type Notes struct {
	ID         int
	UserID     int
	UserName   string
	Title      string
	Content    string
	Tags       string
	CreateTime int64
	UpdateTime int64
}

type NotesTag struct {
	NotesID int
	Content string
}
