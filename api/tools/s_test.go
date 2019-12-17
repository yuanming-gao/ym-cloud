//Package tools
//@ Author: Gao YuanMing
//@ Data: 2019/12/17 5:46 下午
//@ Description:

package tools

import "testing"

func TestConversionBitSize(t *testing.T) {
	t.Log("结果是：", ConversionBitSize(1024*1024*29))
}
