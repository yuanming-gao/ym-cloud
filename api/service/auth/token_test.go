//Package auth
//@ Author: Gao YuanMing
//@ Data: 2019/12/10 2:28 下午
//@ Description:

package auth

import (
	"api/dao/mysql"
	"testing"
)

func TestToken(t *testing.T) {
	mysql.ConnDatabaseAndCreateTables()
	/*token, err := createAccessToken(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)*/
	err := CheckAccessToken("aa6072e8-c776-42bd-87be-9A63", 1)
	t.Log("结果:", err)
}
