//Package dto :
// @Time : 2019/11/26 8:29 下午
// @Author : GaoYuanMing
// @Package : dto
// @FileName : file_dto.go
package dto

//FileInfoDto :
type FileInfoDto struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Size int    `json:"size"`
}
