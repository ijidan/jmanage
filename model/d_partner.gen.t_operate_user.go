package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TOperateUserMgr struct {
	*_BaseMgr
}

// TOperateUserMgr open func
func TOperateUserMgr(db *gorm.DB) *_TOperateUserMgr {
	if db == nil {
		panic(fmt.Errorf("TOperateUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TOperateUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_operate_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TOperateUserMgr) GetTableName() string {
	return "t_operate_user"
}

// Get 获取
func (obj *_TOperateUserMgr) Get() (result TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TOperateUserMgr) Gets() (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用户id
func (obj *_TOperateUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPasswd passwd获取 账户密码
func (obj *_TOperateUserMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithEmail email获取 工作邮箱
func (obj *_TOperateUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithNameEn name_en获取 英文名
func (obj *_TOperateUserMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithNameCn name_cn获取 中文名
func (obj *_TOperateUserMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithPermissions permissions获取 权限ID集合，逗号分隔
func (obj *_TOperateUserMgr) WithPermissions(permissions string) Option {
	return optionFunc(func(o *options) { o.query["permissions"] = permissions })
}

// WithRoleID role_id获取 角色ID
func (obj *_TOperateUserMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_TOperateUserMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_TOperateUserMgr) WithUpdateTime(updateTime int) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithLowerUserID lower_user_id获取 下级用户ID，逗号分隔
func (obj *_TOperateUserMgr) WithLowerUserID(lowerUserID string) Option {
	return optionFunc(func(o *options) { o.query["lower_user_id"] = lowerUserID })
}

// WithIsDisabled is_disabled获取 是否禁用
func (obj *_TOperateUserMgr) WithIsDisabled(isDisabled int8) Option {
	return optionFunc(func(o *options) { o.query["is_disabled"] = isDisabled })
}

// GetByOption 功能选项模式获取
func (obj *_TOperateUserMgr) GetByOption(opts ...Option) (result TOperateUser, err error) {
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
func (obj *_TOperateUserMgr) GetByOptions(opts ...Option) (results []*TOperateUser, err error) {
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

// GetFromID 通过id获取内容 用户id
func (obj *_TOperateUserMgr) GetFromID(id int) (result TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 用户id
func (obj *_TOperateUserMgr) GetBatchFromID(ids []int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容 账户密码
func (obj *_TOperateUserMgr) GetFromPasswd(passwd string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量唯一主键查找 账户密码
func (obj *_TOperateUserMgr) GetBatchFromPasswd(passwds []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 工作邮箱
func (obj *_TOperateUserMgr) GetFromEmail(email string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 工作邮箱
func (obj *_TOperateUserMgr) GetBatchFromEmail(emails []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromNameEn 通过name_en获取内容 英文名
func (obj *_TOperateUserMgr) GetFromNameEn(nameEn string) (result TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name_en = ?", nameEn).Find(&result).Error

	return
}

// GetBatchFromNameEn 批量唯一主键查找 英文名
func (obj *_TOperateUserMgr) GetBatchFromNameEn(nameEns []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name_en IN (?)", nameEns).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 中文名
func (obj *_TOperateUserMgr) GetFromNameCn(nameCn string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name_cn = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量唯一主键查找 中文名
func (obj *_TOperateUserMgr) GetBatchFromNameCn(nameCns []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name_cn IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromPermissions 通过permissions获取内容 权限ID集合，逗号分隔
func (obj *_TOperateUserMgr) GetFromPermissions(permissions string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permissions = ?", permissions).Find(&results).Error

	return
}

// GetBatchFromPermissions 批量唯一主键查找 权限ID集合，逗号分隔
func (obj *_TOperateUserMgr) GetBatchFromPermissions(permissionss []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("permissions IN (?)", permissionss).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色ID
func (obj *_TOperateUserMgr) GetFromRoleID(roleID int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色ID
func (obj *_TOperateUserMgr) GetBatchFromRoleID(roleIDs []int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_TOperateUserMgr) GetFromCreateTime(createTime int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_TOperateUserMgr) GetBatchFromCreateTime(createTimes []int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_TOperateUserMgr) GetFromUpdateTime(updateTime int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_TOperateUserMgr) GetBatchFromUpdateTime(updateTimes []int) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromLowerUserID 通过lower_user_id获取内容 下级用户ID，逗号分隔
func (obj *_TOperateUserMgr) GetFromLowerUserID(lowerUserID string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lower_user_id = ?", lowerUserID).Find(&results).Error

	return
}

// GetBatchFromLowerUserID 批量唯一主键查找 下级用户ID，逗号分隔
func (obj *_TOperateUserMgr) GetBatchFromLowerUserID(lowerUserIDs []string) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lower_user_id IN (?)", lowerUserIDs).Find(&results).Error

	return
}

// GetFromIsDisabled 通过is_disabled获取内容 是否禁用
func (obj *_TOperateUserMgr) GetFromIsDisabled(isDisabled int8) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_disabled = ?", isDisabled).Find(&results).Error

	return
}

// GetBatchFromIsDisabled 批量唯一主键查找 是否禁用
func (obj *_TOperateUserMgr) GetBatchFromIsDisabled(isDisableds []int8) (results []*TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_disabled IN (?)", isDisableds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TOperateUserMgr) FetchByPrimaryKey(id int) (result TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByNameEn primay or index 获取唯一内容
func (obj *_TOperateUserMgr) FetchUniqueByNameEn(nameEn string) (result TOperateUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name_en = ?", nameEn).Find(&result).Error

	return
}
