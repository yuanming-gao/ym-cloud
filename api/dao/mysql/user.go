//Package mysql :
// @Time : 2019/12/2 2:24 下午
// @Author : GaoYuanMing
// @Package : mysql
// @FileName : user.go
package mysql

import (
	"api/model/entity"
	"database/sql"
	"log"
)

//InsertUser :插入用户
func InsertUser(user *entity.User) error {
	sqlStr := "INSERT INTO user_table (user_phone, user_name, user_password, user_position, create_time, update_time) VALUES (?,?,?,?,?,?)"
	insertStmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = insertStmt.Exec(user.Phone, user.Name, user.Password, user.Position, user.CreateTime, user.UpdateTime)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err = insertStmt.Close(); err != nil {
			log.Println(err)
		}
	}()
	return nil
}

//DeleteUser :删除用户
func DeleteUser(user *entity.User) error {
	sqlStr := "DELETE FROM user_table WHERE user_id=?"
	_, err := db.Exec(sqlStr, user.ID)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return err
	}
	return nil
}

//IsExistUserPhone :手机号码是否注册
func IsExistUserPhone(phone string) bool {
	sqlStr := "SELECT * FROM user_table WHERE user_phone=? LIMIT 1"
	err := db.QueryRow(sqlStr, phone).Scan()
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

//SelectUser :查询用户
func SelectUser(phone string) *entity.User {
	sqlStr := "SELECT * FROM user_table WHERE user_phone=?"
	queryRows := db.QueryRow(sqlStr, phone)
	user := new(entity.User)
	err := queryRows.Scan(&user.ID, &user.Phone, &user.Name, &user.Password, &user.Position, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		log.Println("db.QueryRow error:", err)
		return nil
	}
	return user
}

//UpdateUserName :修改用户名字
func UpdateUserName(user *entity.User) error {
	sqlStr := "UPDATE user_table SET user_name=? WHERE id=?"
	_, err := db.Exec(sqlStr, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateUserPhone :修改用户手机
func UpdateUserPhone(user *entity.User) error {
	sqlStr := "UPDATE user_table SET user_phone=? WHERE id=?"
	_, err := db.Exec(sqlStr, user.Phone, user.ID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateUserPassword :修改用户密码
func UpdateUserPassword(user *entity.User) error {
	sqlStr := "UPDATE user_table SET user_password=? WHERE id=?"
	_, err := db.Exec(sqlStr, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}
