package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/model"
	"github.com/ijidan/jutil/jutil"
)

//资讯
type NewsController struct {
	controller.BaseController
}

//列表
func (c *NewsController) GetList(context *gin.Context)  {
	pageSize:=c.GetPageSize(context)
	offset:=c.GetPageOffset(context)
	//查询
	var cnt int64
	db:=jutil.NewDbPartner().Order("f_id desc").Offset(offset).Limit(pageSize)
	newsMgr:=model.TNewsMgr(db)
	list, _ :=newsMgr.Gets()
	newsMgr.Count(&cnt)
	data:=gin.H{
		"cnt":cnt,
		"list":list,
	}
	c.JsonSuccess(context,"",data,"")
}

//编辑
func (c *NewsController) Edit(context *gin.Context)  {

}
