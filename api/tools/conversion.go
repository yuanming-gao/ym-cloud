//Package tools
//@ Author: Gao YuanMing
//@ Data: 2019/12/17 5:17 下午
//@ Description:

package tools

import (
	"strconv"
)

const (
	KB = 1024
	MB = 1048576
	GB = 1073741824
	TB = 1099511627776
)

func ConversionBitSize(size int64) string {
	if size > 0 && size < KB {
		return strconv.FormatInt(size, 10) + "B"
	}
	if size >= KB && size < MB {
		return strconv.FormatInt(size/KB, 10) + "." + strconv.FormatInt(size%KB/100, 10) + "KB"
	}
	if size >= MB && size < GB {
		return strconv.FormatInt(size/MB, 10) + "." + strconv.FormatInt(size%MB/10000, 10) + "MB"
	}
	if size >= GB && size < TB {
		return strconv.FormatInt(size/GB, 10) + "." + strconv.FormatInt(size%GB/100, 10) + "GB"
	} else {
		return strconv.FormatInt(size/TB, 10) + "." + strconv.FormatInt(size%TB/100, 10) + "TB"
	}
}
