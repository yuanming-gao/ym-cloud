//Package dto
//@ Author: Gao YuanMing
//@ Data: 2019/12/11 9:21 下午
//@ Description:

package dto

type CreateNotesDto struct {
	UserName string   `json:"user_name"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

type NotesInfoDto struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Tags       []string `json:"tags"`
	CreateTime string   `json:"create_time"` //返回前端不需要修改的的格式
}

type NotesDto struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Tags       []string `json:"tags"`
	CreateTime string   `json:"create_time"`
	UpdateTime string   `json:"update_time"`
}
