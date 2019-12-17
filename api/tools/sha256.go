//Package tools :
// @Time : 2019/11/25 8:34 下午
// @Author : GaoYuanMing
// @Package : tools
// @FileName : file.go
package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

//NewSha256String :根据字节获取sha256值
func NewSha256String(data []byte) string {
	hs := sha256.New()
	hs.Write(data)
	return hex.EncodeToString(hs.Sum(nil))
}
