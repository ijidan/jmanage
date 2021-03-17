package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/logic"
	"github.com/ijidan/jutil/jutil"
	"log"
)

type RedisController struct {
	controller.BaseController
}

//获取所有数据库
func (c *RedisController) GetDb(context *gin.Context) {
	client:=jutil.NewRedisClient()
	keys:=client.GetAllKeys()
	log.Println(keys)
	menuMap := logic.NewMenu().GetMenu()
	c.JsonSuccess(context, "", menuMap, "")
}
