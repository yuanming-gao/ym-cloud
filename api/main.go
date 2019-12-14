package main

import (
	"api/cmd/router"
	"api/config"
	db "api/dao/mysql"
)

func main() {
	//注册所有路由
	r := router.RegisterAllRoutes()

	//加载连接数据库
	db.ConnDatabaseAndCreateTables()
	err := r.Run(config.APIServerAddress)
	if err != nil {
		panic(err)
	}
}
