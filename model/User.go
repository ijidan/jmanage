package model

import "gorm.io/gorm"

const (
	IsActiveYes = 1 //已激活
	IsActiveNo  = 0 //未激活
)

//房源概况
type User struct {
	gorm.Model
	Email    string
	Password string
	IsActive int `gorm:"type:enum(0,1);default:0"`
}

//表名
func (u *User) TableName() string {
	return "t_user"
}
