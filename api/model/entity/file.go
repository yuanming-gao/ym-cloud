//Package entity
// @Time : 2019/11/26 8:29 下午
// @Author : GaoYuanMing
// @Package : entity
// @FileName : file.go

package entity

//File :
type File struct {
	Sha256Value string //作为文件的主键
	Name        string
	Type        string //文件类型
	UserID      int
	Size        int64 //文件大小
	Status      int   //文件状态应为枚举类型
	LocalAt     string
	UploadAt    string
	CreateTime  int64
}

func CreateFile() *File {
	return nil
}
