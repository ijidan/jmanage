package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TOperateLogMgr struct {
	*_BaseMgr
}

// TOperateLogMgr open func
func TOperateLogMgr(db *gorm.DB) *_TOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("TOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TOperateLogMgr) GetTableName() string {
	return "t_operate_log"
}

// Get 获取
func (obj *_TOperateLogMgr) Get() (result TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TOperateLogMgr) Gets() (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithFID f_id获取
func (obj *_TOperateLogMgr) WithFID(fID int) Option {
	return optionFunc(func(o *options) { o.query["f_id"] = fID })
}

// WithFUId f_uid获取 操作用户ID
func (obj *_TOperateLogMgr) WithFUId(fUId int) Option {
	return optionFunc(func(o *options) { o.query["f_uid"] = fUId })
}

// WithFName f_name获取 操作用户名
func (obj *_TOperateLogMgr) WithFName(fName string) Option {
	return optionFunc(func(o *options) { o.query["f_name"] = fName })
}

// WithFCtime f_ctime获取 操作时间
func (obj *_TOperateLogMgr) WithFCtime(fCtime int) Option {
	return optionFunc(func(o *options) { o.query["f_ctime"] = fCtime })
}

// WithFMethod f_method获取 操作方式；get，post
func (obj *_TOperateLogMgr) WithFMethod(fMethod string) Option {
	return optionFunc(func(o *options) { o.query["f_method"] = fMethod })
}

// WithFURL f_url获取 操作URL
func (obj *_TOperateLogMgr) WithFURL(fURL string) Option {
	return optionFunc(func(o *options) { o.query["f_url"] = fURL })
}

// WithFData f_data获取 操作数据
func (obj *_TOperateLogMgr) WithFData(fData string) Option {
	return optionFunc(func(o *options) { o.query["f_data"] = fData })
}

// GetByOption 功能选项模式获取
func (obj *_TOperateLogMgr) GetByOption(opts ...Option) (result TOperateLog, err error) {
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
func (obj *_TOperateLogMgr) GetByOptions(opts ...Option) (results []*TOperateLog, err error) {
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

// GetFromFID 通过f_id获取内容
func (obj *_TOperateLogMgr) GetFromFID(fID int) (result TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id = ?", fID).Find(&result).Error

	return
}

// GetBatchFromFID 批量唯一主键查找
func (obj *_TOperateLogMgr) GetBatchFromFID(fIDs []int) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id IN (?)", fIDs).Find(&results).Error

	return
}

// GetFromFUId 通过f_uid获取内容 操作用户ID
func (obj *_TOperateLogMgr) GetFromFUId(fUId int) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_uid = ?", fUId).Find(&results).Error

	return
}

// GetBatchFromFUId 批量唯一主键查找 操作用户ID
func (obj *_TOperateLogMgr) GetBatchFromFUId(fUIds []int) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_uid IN (?)", fUIds).Find(&results).Error

	return
}

// GetFromFName 通过f_name获取内容 操作用户名
func (obj *_TOperateLogMgr) GetFromFName(fName string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_name = ?", fName).Find(&results).Error

	return
}

// GetBatchFromFName 批量唯一主键查找 操作用户名
func (obj *_TOperateLogMgr) GetBatchFromFName(fNames []string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_name IN (?)", fNames).Find(&results).Error

	return
}

// GetFromFCtime 通过f_ctime获取内容 操作时间
func (obj *_TOperateLogMgr) GetFromFCtime(fCtime int) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_ctime = ?", fCtime).Find(&results).Error

	return
}

// GetBatchFromFCtime 批量唯一主键查找 操作时间
func (obj *_TOperateLogMgr) GetBatchFromFCtime(fCtimes []int) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_ctime IN (?)", fCtimes).Find(&results).Error

	return
}

// GetFromFMethod 通过f_method获取内容 操作方式；get，post
func (obj *_TOperateLogMgr) GetFromFMethod(fMethod string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_method = ?", fMethod).Find(&results).Error

	return
}

// GetBatchFromFMethod 批量唯一主键查找 操作方式；get，post
func (obj *_TOperateLogMgr) GetBatchFromFMethod(fMethods []string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_method IN (?)", fMethods).Find(&results).Error

	return
}

// GetFromFURL 通过f_url获取内容 操作URL
func (obj *_TOperateLogMgr) GetFromFURL(fURL string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_url = ?", fURL).Find(&results).Error

	return
}

// GetBatchFromFURL 批量唯一主键查找 操作URL
func (obj *_TOperateLogMgr) GetBatchFromFURL(fURLs []string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_url IN (?)", fURLs).Find(&results).Error

	return
}

// GetFromFData 通过f_data获取内容 操作数据
func (obj *_TOperateLogMgr) GetFromFData(fData string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_data = ?", fData).Find(&results).Error

	return
}

// GetBatchFromFData 批量唯一主键查找 操作数据
func (obj *_TOperateLogMgr) GetBatchFromFData(fDatas []string) (results []*TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_data IN (?)", fDatas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TOperateLogMgr) FetchByPrimaryKey(fID int) (result TOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id = ?", fID).Find(&result).Error

	return
}
