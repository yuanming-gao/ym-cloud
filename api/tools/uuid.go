package tools

import (
	"crypto/rand"
	"fmt"
	"io"
)

//CreateUUID :
func CreateUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if len(uuid) != n || err != nil {
		panic(err)
		return ""
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%X", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:12])
}
