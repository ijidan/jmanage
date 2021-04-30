package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	_ "github.com/gogf/gf/util/gconv"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/model"
	"github.com/ijidan/jutil/jutil"
	"net/http"
)

//资讯
type NewsController struct {
	controller.BaseController
}

//资讯列表
func (c *NewsController) List(context *gin.Context) {
	context.HTML(http.StatusOK, "news/list.html", nil)
}

//获取资讯数据
func (c *NewsController) GetList(context *gin.Context) {
	pageSize := c.GetPageSize(context)
	offset := c.GetPageOffset(context)
	//查询
	var cnt int64
	db := jutil.NewDbPartner().Order("f_id desc").Offset(offset).Limit(pageSize)
	newsMgr := model.TNewsMgr(db)
	list, _ := newsMgr.Gets()
	newsMgr.Count(&cnt)
	count := gconv.Uint(cnt)
	data := gconv.SliceMap(list)
	c.JsonSuccess2(context, "", count, data, "")
}

//编辑
func (c *NewsController) Edit(context *gin.Context) {

}
