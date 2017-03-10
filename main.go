package main

import (
	"GoReadNote/handlers"
	"GoReadNote/logger"
	"GoReadNote/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为 gin.ReleaseMode  gin.DebugMode
	router := gin.Default()      //获得路由实例
	//网页请求----------------------
	//添加中间件
	router.Use(middleware.Middleware)
	//搜索小说
	router.GET("/", handlers.HomeHandler)
	router.GET("/SearchNote", handlers.SearchNoteHandler)
	router.GET("/GetBookInfo", handlers.GetBookInfoHandler)
	//获取章节内容
	router.GET("/GetBookContent", handlers.GetNoteContentHandler)

	//JSON请求----------------------
	router.GET("/GetJson", handlers.GetJsonHandler)
	router.GET("/GetSearchNoteJson", handlers.GetSearchNoteJsonHandler)
	router.GET("/GetBookContentJson", handlers.GetBookContentJsonHandler)
	router.GET("/GetTopNoteListJson", handlers.GetTopNoteListJsonHandler)
	router.GET("/GetNoteInfoJson", handlers.GetNoteInfoJsonHandler)

	//文件上传
	router.GET("/UploadFile", handlers.GetUpLoadPageHandler)
	router.POST("/UploadFile", handlers.UploadFileHandler)
	router.GET("/GetFileListJson", handlers.GetFileListJsonHandler)
	router.StaticFS("/Main", http.Dir("./savefile/main"))
	router.StaticFS("/Weifei", http.Dir("./savefile/wei"))
	//icon
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")

	logger.ALogger().Notice("Listen start.")
	logger.ALogger().Notice("Listen 443 https")
	//监听端口
	//http.ListenAndServe(":8005", router)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	//8000端口是测试之用 实际端口为443
	err := http.ListenAndServeTLS(":4433", "./ca/1_fsnsaber.cn_bundle.crt", "./ca/2_fsnsaber.cn.key", router)
	//http.ListenAndServeTLS(":443","2_fsnsaber.cn.crt","3_fsnsaber.cn.key",router)
	logger.ALogger().Error(err)

}
