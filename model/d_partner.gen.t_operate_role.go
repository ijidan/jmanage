package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TOperateRoleMgr struct {
	*_BaseMgr
}

// TOperateRoleMgr open func
func TOperateRoleMgr(db *gorm.DB) *_TOperateRoleMgr {
	if db == nil {
		panic(fmt.Errorf("TOperateRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TOperateRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_operate_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TOperateRoleMgr) GetTableName() string {
	return "t_operate_role"
}

// Get 获取
func (obj *_TOperateRoleMgr) Get() (result TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TOperateRoleMgr) Gets() (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TOperateRoleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_TOperateRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPermissions permissions获取 权限菜单
func (obj *_TOperateRoleMgr) WithPermissions(permissions string) Option {
	return optionFunc(func(o *options) { o.query["permissions"] = permissions })
}

// WithCreateTime create_time获取
func (obj *_TOperateRoleMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_TOperateRoleMgr) WithUpdateTime(updateTime int) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithDesc desc获取 备注
func (obj *_TOperateRoleMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// GetByOption 功能选项模式获取
func (obj *_TOperateRoleMgr) GetByOption(opts ...Option) (result TOperateRole, err error) {
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
func (obj *_TOperateRoleMgr) GetByOptions(opts ...Option) (results []*TOperateRole, err error) {
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

// GetFromID 通过id获取内容
func (obj *_TOperateRoleMgr) GetFromID(id int) (result TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TOperateRoleMgr) GetBatchFromID(ids []int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TOperateRoleMgr) GetFromName(name string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TOperateRoleMgr) GetBatchFromName(names []string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPermissions 通过permissions获取内容 权限菜单
func (obj *_TOperateRoleMgr) GetFromPermissions(permissions string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permissions = ?", permissions).Find(&results).Error

	return
}

// GetBatchFromPermissions 批量唯一主键查找 权限菜单
func (obj *_TOperateRoleMgr) GetBatchFromPermissions(permissionss []string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permissions IN (?)", permissionss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_TOperateRoleMgr) GetFromCreateTime(createTime int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_TOperateRoleMgr) GetBatchFromCreateTime(createTimes []int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_TOperateRoleMgr) GetFromUpdateTime(updateTime int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_TOperateRoleMgr) GetBatchFromUpdateTime(updateTimes []int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 备注
func (obj *_TOperateRoleMgr) GetFromDesc(desc string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 备注
func (obj *_TOperateRoleMgr) GetBatchFromDesc(descs []string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TOperateRoleMgr) FetchByPrimaryKey(id int) (result TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByID  获取多个内容
func (obj *_TOperateRoleMgr) FetchIndexByID(id int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_TOperateRoleMgr) FetchIndexByName(name string) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// FetchIndexByCreateTime  获取多个内容
func (obj *_TOperateRoleMgr) FetchIndexByCreateTime(createTime int) (results []*TOperateRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}
