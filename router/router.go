package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
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
	news := admin.NewsController{}
	sys := admin.SysController{}

	r.Any("/", index.Index)
	r.Any("/menu/getMenu", menu.GetMenu)
	r.Any("/r", redis.GetDb)
	//用户相关
	r.Any("/user/reg", user.Reg)
	r.Any("/user/login", user.Login)
	//资讯相关
	r.Any("/news/list", news.List)
	r.Any("/news/getList", news.GetList)
	r.Any("news/edit", news.Edit)
	//系统相关
	r.Any("/sys/list", sys.List)
	//文件相关
	r.Any("/doc/upload", doc.Upload)
	r.Any("/doc/multiUpload", doc.MultiUpload)
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
	r.HTMLRender = controller.LoadTemplates("admin")
	r.StaticFS("/static", http.Dir("static"))
	r.Use(RecoveryHandler())
	r.Use(GoroutineTest())
	r.Use(LogRecord())

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
