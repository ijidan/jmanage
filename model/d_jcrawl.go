package model

import (
	"gorm.io/gorm"
)

// TUser 用户表
type TUser struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"` // 邮箱
	Password string `json:"password"`            // 密码
	IsActive int    `json:"is_active"`           // 是否激活：0-未激活，1-已激活
}
