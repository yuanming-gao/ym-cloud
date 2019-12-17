/**
* @Time : 2019/11/25 4:48 下午
* @Author : GaoYuanMing
* @Package : router
* @FileName : router.go
 */
package router

import (
	"api/cmd/handler"
	"api/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes() *gin.Engine {
	r := gin.Default()
	//全局中间件,允许跨域访问
	r.Use(middleware.Cors)

	//公共路由不需要身份验证
	public := r.Group("/api/v1")
	public.GET("/user", handler.SignInHandler)
	public.POST("/user", handler.SignUpHandler)

	apiRoutes := r.Group("/api/v1")
	//所有路由都需身份验证中间件
	apiRoutes.Use(middleware.Auth)
	//在该用户id下新建笔记
	apiRoutes.POST("/user/:user_id/notes", handler.CreateNotes)
	//该用户id下的所有笔记
	apiRoutes.GET("/user/:user_id/notes", handler.ListAllNotes)
	apiRoutes.GET("/user/:user_id/notes/:notes_id", handler.SelectNotes)
	apiRoutes.DELETE("/user/:user_id/notes/:notes_id", handler.DeleteNotes)
	apiRoutes.PUT("/user/:user_id/notes/:notes_id", handler.UpdateNodes)

	apiRoutes.GET("/user/:user_id/file", handler.ListFile)
	apiRoutes.POST("/user/:user_id/file", handler.UploadFile)
	apiRoutes.DELETE("/user/:user_id/file/:sha", handler.DeleteFile)
	apiRoutes.GET("/user/:user_id/file/:name", handler.DownloadFile)
	return r
}
