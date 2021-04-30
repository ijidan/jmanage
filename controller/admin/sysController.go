package admin

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gogf/gf/util/gconv"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/logic"
	"net/http"
)

//资讯
type SysController struct {
	controller.BaseController
}

//资讯列表
func (c *SysController) List(context *gin.Context) {
	sys := logic.NewSys()
	data := sys.GetSysInfo()
	context.HTML(http.StatusOK, "sys/list.html", data)
}
