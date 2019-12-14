//Package mysql :
// @Time : 2019/12/2 3:36 下午
// @Author : GaoYuanMing
// @Package : mysql
// @FileName : login_record.go
package mysql

import (
	"api/model/entity"
	"log"
)

//InsertUserLoginRecord :插入用户登录记录
func InsertUserLoginRecord(record *entity.LoginRecord) error {
	sql := "INSERT INTO user_login_record_table (login_phone, login_ip, login_time) values (?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("InsertUserLoginRecord() db.Prepare error:", err)
		return err
	}
	_, err = stmt.Exec(record.LoginPhone, record.LoginIP, record.LoginTime)
	if err != nil {
		log.Println("InsertUserLoginRecord() stmt.Exec error:", err)
		return err
	}
	return err
}

//QueryUser :查询用户登录记录
func QueryUserLoginRecord(user *entity.User) []*entity.LoginRecord {
	sql := "SELECT login_ip, login_time FROM user_login_record_table WHERE login_phone=?"
	queryRows, err := db.Query(sql, user.Phone)
	defer func() {
		if err = queryRows.Close(); err != nil {
			log.Println("defer close queryRows error:", err)
		}
	}()
	if err != nil {
		log.Println("查询用户登录信息异常:", err)
		return nil
	}

	//初始化用户登录记录列表
	recordList := make([]*entity.LoginRecord, 0)
	for queryRows.Next() {
		//得到每一条记录
		record := new(entity.LoginRecord)
		record.LoginPhone = user.Phone
		if err = queryRows.Scan(&record.LoginIP, &record.LoginTime); err != nil {
			log.Println("查询用户登录信息异常:", err)
			return nil
		}
		recordList = append(recordList, record)
	}

	return recordList
}

//DeleteUserLoginRecord :删除登录记录
func DeleteUserLoginRecord(user *entity.User) error {
	sql := "DELETE FROM user_login_record_table WHERE login_phone=?"
	_, err := db.Exec(sql, user.Phone)
	if err != nil {
		log.Println("DeleteUserLoginRecord error:", err)
		return err
	}
	return nil
}
