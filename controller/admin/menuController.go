package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/logic"
)

type MenuController struct {
	controller.BaseController
}

//获取菜单
func (c *MenuController) GetMenu(context *gin.Context)  {
	menuMap:=logic.NewMenu().GetMenu()
	c.JsonSuccess(context,"",menuMap,"")
}