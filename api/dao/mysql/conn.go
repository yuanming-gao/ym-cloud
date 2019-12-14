//Package mysql :
// @Time : 2019/12/2 2:24 下午
// @Author : GaoYuanMing
// @Package : mysql
// @FileName : conn.go
package mysql

import (
	"api/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	db  *sql.DB
	err error
)

//ConnDatabaseAndCreateTables :
func ConnDatabaseAndCreateTables() {
	db, err = sql.Open("mysql", config.DataSourceName)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(config.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetMaxOpenConns(config.MaxOpenConn)

}
