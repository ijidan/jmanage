package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"net/http"
)

type IndexController struct {
	controller.BaseController
}

//首页
func (c *IndexController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "layout/main.html", nil)
}

//Ping
func (c *IndexController) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})

}
