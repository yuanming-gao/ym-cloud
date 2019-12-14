//Package tools
//@ Author: Gao YuanMing
//@ Data: 2019/12/8 9:29 下午
//@ Description:

package tools

import "time"

const (
	timeLayout = "2006-01-02 15:04:05"
)

//NowTimeToUnixNano :获得当前时间的10位时间戳
func NowTimeToUnixNano() int64 {
	return time.Now().UnixNano() / 1000000000
}

func ParseUnixNanoToString(nano int64) string {
	return time.Unix(nano, 0).Format(timeLayout)
}
