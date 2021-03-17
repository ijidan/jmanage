package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TUserMgr struct {
	*_BaseMgr
}

// TUserMgr open func
func TUserMgr(db *gorm.DB) *_TUserMgr {
	if db == nil {
		panic(fmt.Errorf("TUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TUserMgr) GetTableName() string {
	return "t_user"
}

// Get 获取
func (obj *_TUserMgr) Get() (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TUserMgr) Gets() (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_TUserMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_TUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_TUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_TUserMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithEmail email获取 邮箱
func (obj *_TUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPassword password获取 密码
func (obj *_TUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithIsActive is_active获取 是否激活：0-未激活，1-已激活
func (obj *_TUserMgr) WithIsActive(isActive int) Option {
	return optionFunc(func(o *options) { o.query["is_active"] = isActive })
}

// GetByOption 功能选项模式获取
func (obj *_TUserMgr) GetByOption(opts ...Option) (result TUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TUserMgr) GetByOptions(opts ...Option) (results []*TUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_TUserMgr) GetFromID(id int64) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_TUserMgr) GetBatchFromID(ids []int64) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_TUserMgr) GetFromCreatedAt(createdAt time.Time) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_TUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_TUserMgr) GetFromUpdatedAt(updatedAt time.Time) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_TUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_TUserMgr) GetFromDeletedAt(deletedAt time.Time) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_TUserMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_TUserMgr) GetFromEmail(email string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&result).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱
func (obj *_TUserMgr) GetBatchFromEmail(emails []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_TUserMgr) GetFromPassword(password string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 密码
func (obj *_TUserMgr) GetBatchFromPassword(passwords []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromIsActive 通过is_active获取内容 是否激活：0-未激活，1-已激活
func (obj *_TUserMgr) GetFromIsActive(isActive int) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active = ?", isActive).Find(&results).Error

	return
}

// GetBatchFromIsActive 批量唯一主键查找 是否激活：0-未激活，1-已激活
func (obj *_TUserMgr) GetBatchFromIsActive(isActives []int) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active IN (?)", isActives).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TUserMgr) FetchByPrimaryKey(id int64) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByEmail primay or index 获取唯一内容
func (obj *_TUserMgr) FetchUniqueByEmail(email string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&result).Error

	return
}
