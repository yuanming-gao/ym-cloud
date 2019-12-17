//Package file
//@ Author: Gao YuanMing
//@ Data: 2019/12/17 4:20 下午
//@ Description:

package file

import "api/dao/mysql"

func DeleteFile(sha string) error {
	return mysql.DeleteFileInfoBySha256Value(sha)
}
