package admin

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/pkg"
)

//用户相关
type UserController struct {
	controller.BaseController
}

//首页
func (c *UserController) Reg(context *gin.Context) {
	type Req struct {
		Email       string `form:"email" binding:"required,email"`
		Password    string `form:"password" binding:"required,min=6,max=18"`
		RptPassword string `form:"rpt_password" binding:"eqfield=Password"`
	}
	var req Req
	if err := context.ShouldBind(&req); err != nil {
		message:=c.Trans(context,err)
		c.JsonFail(context, pkg.InvalidParams, message, nil, "")
		return
	}
	//userLogic:=logic.NewUser()
	//user:=userLogic.FindUser(email)
	//if user.ID>0{
	//	if user.IsActive==model.IsActiveYes{
	//		c.JsonFail(context, pkg.ERROR, "该邮箱已经注册", nil, "")
	//		return
	//	}
	//	//发送激活邮件
	//	userLogic.SendActiveEmail(email)
	//}else{
	//	userLogic.AddUser(email,password)
	//	//发送激活邮件
	//	userLogic.SendActiveEmail(email)
	//}
	c.JsonSuccess(context, "注册成功", nil, "")
	return
}

//Ping
func (c *UserController) Login(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})

}
