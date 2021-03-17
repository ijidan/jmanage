package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller/admin"
	"github.com/ijidan/jmanage/controller/www"
	"net/http"
	"time"
)

//注册路由
func WWWRegister(r *gin.Engine) {
	index := www.IndexController{}
	//
	r.GET("/", index.Index)
	r.GET("/ping", index.Ping)

}

//注册路由
func AdminRegister(r *gin.Engine) {
	index := admin.IndexController{}
	menu := admin.MenuController{}
	redis := admin.RedisController{}
	user := admin.UserController{}
	doc := admin.DocController{}
	news:=admin.NewsController{}

	r.GET("/", index.Index)
	r.GET("/menu/getMenu", menu.GetMenu)
	r.GET("/r", redis.GetDb)
	//用户相关
	r.Any("/user/reg", user.Reg)
	r.Any("/user/login", user.Login)
	//资讯相关
	r.GET("/news/getList",news.GetList)
	r.POST("news/edit",news.Edit)
	//文件相关
	r.POST("/doc/upload", doc.Upload)
	r.POST("/doc/multiUpload", doc.MultiUpload)
}

//前台服务
func WWWServer() *http.Server {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("template/www/**/*")
	r.StaticFS("/static", http.Dir("static"))
	WWWRegister(r)
	server := &http.Server{
		Addr:         ":80",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server

}

//后台服务
func AdminServer() *http.Server {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(GoroutineTest())
	r.Use(TimeStat())
	r.LoadHTMLGlob("template/admin/**/*")
	r.StaticFS("/static", http.Dir("static"))
	//记录日志
	//f,_:=os.Create("runtime/j_manage.log")
	//gin.DefaultWriter=io.MultiWriter(f)
	gin.ForceConsoleColor()
	//路由注册
	AdminRegister(r)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server
}
