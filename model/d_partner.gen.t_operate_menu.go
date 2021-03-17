package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TOperateMenuMgr struct {
	*_BaseMgr
}

// TOperateMenuMgr open func
func TOperateMenuMgr(db *gorm.DB) *_TOperateMenuMgr {
	if db == nil {
		panic(fmt.Errorf("TOperateMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TOperateMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_operate_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TOperateMenuMgr) GetTableName() string {
	return "t_operate_menu"
}

// Get 获取
func (obj *_TOperateMenuMgr) Get() (result TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TOperateMenuMgr) Gets() (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TOperateMenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_TOperateMenuMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithModule module获取 模块
func (obj *_TOperateMenuMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithController controller获取 控制器
func (obj *_TOperateMenuMgr) WithController(controller string) Option {
	return optionFunc(func(o *options) { o.query["controller"] = controller })
}

// WithFunction function获取 方法
func (obj *_TOperateMenuMgr) WithFunction(function string) Option {
	return optionFunc(func(o *options) { o.query["function"] = function })
}

// WithParameter parameter获取 参数
func (obj *_TOperateMenuMgr) WithParameter(parameter string) Option {
	return optionFunc(func(o *options) { o.query["parameter"] = parameter })
}

// WithDescription description获取 描述
func (obj *_TOperateMenuMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithIsDisplay is_display获取 1显示在左侧菜单2只作为节点
func (obj *_TOperateMenuMgr) WithIsDisplay(isDisplay int) Option {
	return optionFunc(func(o *options) { o.query["is_display"] = isDisplay })
}

// WithType type获取 1权限节点2普通节点
func (obj *_TOperateMenuMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithPid pid获取 上级菜单0为顶级菜单
func (obj *_TOperateMenuMgr) WithPid(pid int) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithCreateTime create_time获取
func (obj *_TOperateMenuMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_TOperateMenuMgr) WithUpdateTime(updateTime int) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithIcon icon获取 图标
func (obj *_TOperateMenuMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithIsOpen is_open获取 0默认闭合1默认展开
func (obj *_TOperateMenuMgr) WithIsOpen(isOpen int) Option {
	return optionFunc(func(o *options) { o.query["is_open"] = isOpen })
}

// WithOrders orders获取 排序值，越小越靠前
func (obj *_TOperateMenuMgr) WithOrders(orders int) Option {
	return optionFunc(func(o *options) { o.query["orders"] = orders })
}

// GetByOption 功能选项模式获取
func (obj *_TOperateMenuMgr) GetByOption(opts ...Option) (result TOperateMenu, err error) {
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
func (obj *_TOperateMenuMgr) GetByOptions(opts ...Option) (results []*TOperateMenu, err error) {
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
func (obj *_TOperateMenuMgr) GetFromID(id int) (result TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TOperateMenuMgr) GetBatchFromID(ids []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TOperateMenuMgr) GetFromName(name string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TOperateMenuMgr) GetBatchFromName(names []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容 模块
func (obj *_TOperateMenuMgr) GetFromModule(module string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量唯一主键查找 模块
func (obj *_TOperateMenuMgr) GetBatchFromModule(modules []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

// GetFromController 通过controller获取内容 控制器
func (obj *_TOperateMenuMgr) GetFromController(controller string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller = ?", controller).Find(&results).Error

	return
}

// GetBatchFromController 批量唯一主键查找 控制器
func (obj *_TOperateMenuMgr) GetBatchFromController(controllers []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller IN (?)", controllers).Find(&results).Error

	return
}

// GetFromFunction 通过function获取内容 方法
func (obj *_TOperateMenuMgr) GetFromFunction(function string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("function = ?", function).Find(&results).Error

	return
}

// GetBatchFromFunction 批量唯一主键查找 方法
func (obj *_TOperateMenuMgr) GetBatchFromFunction(functions []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("function IN (?)", functions).Find(&results).Error

	return
}

// GetFromParameter 通过parameter获取内容 参数
func (obj *_TOperateMenuMgr) GetFromParameter(parameter string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parameter = ?", parameter).Find(&results).Error

	return
}

// GetBatchFromParameter 批量唯一主键查找 参数
func (obj *_TOperateMenuMgr) GetBatchFromParameter(parameters []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parameter IN (?)", parameters).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_TOperateMenuMgr) GetFromDescription(description string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找 描述
func (obj *_TOperateMenuMgr) GetBatchFromDescription(descriptions []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromIsDisplay 通过is_display获取内容 1显示在左侧菜单2只作为节点
func (obj *_TOperateMenuMgr) GetFromIsDisplay(isDisplay int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_display = ?", isDisplay).Find(&results).Error

	return
}

// GetBatchFromIsDisplay 批量唯一主键查找 1显示在左侧菜单2只作为节点
func (obj *_TOperateMenuMgr) GetBatchFromIsDisplay(isDisplays []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_display IN (?)", isDisplays).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 1权限节点2普通节点
func (obj *_TOperateMenuMgr) GetFromType(_type int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 1权限节点2普通节点
func (obj *_TOperateMenuMgr) GetBatchFromType(_types []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 上级菜单0为顶级菜单
func (obj *_TOperateMenuMgr) GetFromPid(pid int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量唯一主键查找 上级菜单0为顶级菜单
func (obj *_TOperateMenuMgr) GetBatchFromPid(pids []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid IN (?)", pids).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_TOperateMenuMgr) GetFromCreateTime(createTime int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_TOperateMenuMgr) GetBatchFromCreateTime(createTimes []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_TOperateMenuMgr) GetFromUpdateTime(updateTime int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找
func (obj *_TOperateMenuMgr) GetBatchFromUpdateTime(updateTimes []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_TOperateMenuMgr) GetFromIcon(icon string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_TOperateMenuMgr) GetBatchFromIcon(icons []string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromIsOpen 通过is_open获取内容 0默认闭合1默认展开
func (obj *_TOperateMenuMgr) GetFromIsOpen(isOpen int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_open = ?", isOpen).Find(&results).Error

	return
}

// GetBatchFromIsOpen 批量唯一主键查找 0默认闭合1默认展开
func (obj *_TOperateMenuMgr) GetBatchFromIsOpen(isOpens []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_open IN (?)", isOpens).Find(&results).Error

	return
}

// GetFromOrders 通过orders获取内容 排序值，越小越靠前
func (obj *_TOperateMenuMgr) GetFromOrders(orders int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("orders = ?", orders).Find(&results).Error

	return
}

// GetBatchFromOrders 批量唯一主键查找 排序值，越小越靠前
func (obj *_TOperateMenuMgr) GetBatchFromOrders(orderss []int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("orders IN (?)", orderss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TOperateMenuMgr) FetchByPrimaryKey(id int) (result TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByID  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByID(id int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// FetchIndexByModule  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByModule(module string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// FetchIndexByController  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByController(controller string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller = ?", controller).Find(&results).Error

	return
}

// FetchIndexByFunction  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByFunction(function string) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("function = ?", function).Find(&results).Error

	return
}

// FetchIndexByIsDisplay  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByIsDisplay(isDisplay int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_display = ?", isDisplay).Find(&results).Error

	return
}

// FetchIndexByType  获取多个内容
func (obj *_TOperateMenuMgr) FetchIndexByType(_type int) (results []*TOperateMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}
