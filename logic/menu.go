package logic

import (
	"github.com/bitly/go-simplejson"
)

//菜单
type Menu struct {
}

//获取菜单
func (l *Menu) GetMenu() map[string]interface{} {
	menuStr := `{"data":[{
	"menuId": "1",
	"menuName": "数据管理",
	"menuIcon": "",
	"menuHref": "\/service\/Project_Relation\/index?",
	"parentMenuId": "0"
}, {
	"menuId": "101",
	"menuName": "Redis管理",
	"menuIcon": "",
	"menuHref": "\/common_user\/user_sync\/index?",
	"parentMenuId": "1"
}]}`
	var m map[string]interface{}
	js, _ := simplejson.NewJson([]byte(menuStr))
	m,_=js.Map()
	return m
}

//获取实例
func NewMenu() *Menu {
	return &Menu{}
}
