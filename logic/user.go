package logic

import (
	"github.com/ijidan/jmanage/model"
	"github.com/ijidan/jutil/jutil"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//用户逻辑
type User struct {
	model.User
	db *gorm.DB
}

//发送激活邮件
func (u *User) SendActiveEmail(email string) {
	user:=u.FindUser(email)
	url:=u.GenActiveUrl(user.ID)
	emailUtil:=jutil.NewEmailUtil()
	emailUtil.Send(email,"激活邮件","请点击下面链接进行激活：<br>"+url)
}

//生成激活链接
func (u *User) GenActiveUrl(userId uint) string  {
	return ""
}

//根据邮箱查询用户
func (u *User) FindUser(email string) *model.User {
	var user model.User
	u.db.Where("email = ?", email).First(&user)
	return &user
}

//添加用户
func (u *User) AddUser(email string,password string) *model.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	user:=model.User{
		Email:    email,
		Password: string(hash),
		IsActive: model.IsActiveNo,
	}
	u.db.Create(&user)
	return &user
}

//密码加密
func (u *User) HashPassword(password string)string  {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}

//密码比较
func (u *User) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

//获取实例
func NewUser() *User {
	db := jutil.NewDb("")
	user := &User{
		User: model.User{},
		db:   db,
	}
	return user
}
