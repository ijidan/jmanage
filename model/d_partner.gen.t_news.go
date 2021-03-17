package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TNewsMgr struct {
	*_BaseMgr
}

// TNewsMgr open func
func TNewsMgr(db *gorm.DB) *_TNewsMgr {
	if db == nil {
		panic(fmt.Errorf("TNewsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TNewsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_news"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TNewsMgr) GetTableName() string {
	return "t_news"
}

// Get 获取
func (obj *_TNewsMgr) Get() (result TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TNewsMgr) Gets() (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithFID f_id获取 主键
func (obj *_TNewsMgr) WithFID(fID uint32) Option {
	return optionFunc(func(o *options) { o.query["f_id"] = fID })
}

// WithFTitle f_title获取 标题
func (obj *_TNewsMgr) WithFTitle(fTitle string) Option {
	return optionFunc(func(o *options) { o.query["f_title"] = fTitle })
}

// WithFThumbnail f_thumbnail获取 缩略图
func (obj *_TNewsMgr) WithFThumbnail(fThumbnail string) Option {
	return optionFunc(func(o *options) { o.query["f_thumbnail"] = fThumbnail })
}

// WithFAbstract f_abstract获取 摘要
func (obj *_TNewsMgr) WithFAbstract(fAbstract string) Option {
	return optionFunc(func(o *options) { o.query["f_abstract"] = fAbstract })
}

// WithFContent f_content获取 内容
func (obj *_TNewsMgr) WithFContent(fContent string) Option {
	return optionFunc(func(o *options) { o.query["f_content"] = fContent })
}

// WithFCat f_cat获取 类目
func (obj *_TNewsMgr) WithFCat(fCat string) Option {
	return optionFunc(func(o *options) { o.query["f_cat"] = fCat })
}

// WithFAuthor f_author获取 作者
func (obj *_TNewsMgr) WithFAuthor(fAuthor string) Option {
	return optionFunc(func(o *options) { o.query["f_author"] = fAuthor })
}

// WithFView f_view获取 阅读量
func (obj *_TNewsMgr) WithFView(fView uint32) Option {
	return optionFunc(func(o *options) { o.query["f_view"] = fView })
}

// WithFTag f_tag获取 Tag标签，英文逗号分隔
func (obj *_TNewsMgr) WithFTag(fTag string) Option {
	return optionFunc(func(o *options) { o.query["f_tag"] = fTag })
}

// WithFIsValid f_is_valid获取 1-有效，0-无效
func (obj *_TNewsMgr) WithFIsValid(fIsValid bool) Option {
	return optionFunc(func(o *options) { o.query["f_is_valid"] = fIsValid })
}

// WithFOrder f_order获取 同级中的顺序，1-n，从上到下
func (obj *_TNewsMgr) WithFOrder(fOrder uint8) Option {
	return optionFunc(func(o *options) { o.query["f_order"] = fOrder })
}

// WithFReviewStatus f_review_status获取 审核状态
func (obj *_TNewsMgr) WithFReviewStatus(fReviewStatus int8) Option {
	return optionFunc(func(o *options) { o.query["f_review_status"] = fReviewStatus })
}

// WithFReviewTip f_review_tip获取 审核提示
func (obj *_TNewsMgr) WithFReviewTip(fReviewTip string) Option {
	return optionFunc(func(o *options) { o.query["f_review_tip"] = fReviewTip })
}

// WithFSynStatus f_syn_status获取 0-未同步，1-已经同步
func (obj *_TNewsMgr) WithFSynStatus(fSynStatus uint8) Option {
	return optionFunc(func(o *options) { o.query["f_syn_status"] = fSynStatus })
}

// WithFSynID f_syn_id获取 同步资讯ID
func (obj *_TNewsMgr) WithFSynID(fSynID int) Option {
	return optionFunc(func(o *options) { o.query["f_syn_id"] = fSynID })
}

// WithFOperatorID f_operator_id获取 操作者ID
func (obj *_TNewsMgr) WithFOperatorID(fOperatorID int) Option {
	return optionFunc(func(o *options) { o.query["f_operator_id"] = fOperatorID })
}

// WithFUpdateTime f_update_time获取 更新时间
func (obj *_TNewsMgr) WithFUpdateTime(fUpdateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["f_update_time"] = fUpdateTime })
}

// WithFCreateTime f_create_time获取 创建时间
func (obj *_TNewsMgr) WithFCreateTime(fCreateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["f_create_time"] = fCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_TNewsMgr) GetByOption(opts ...Option) (result TNews, err error) {
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
func (obj *_TNewsMgr) GetByOptions(opts ...Option) (results []*TNews, err error) {
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

// GetFromFID 通过f_id获取内容 主键
func (obj *_TNewsMgr) GetFromFID(fID uint32) (result TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id = ?", fID).Find(&result).Error

	return
}

// GetBatchFromFID 批量唯一主键查找 主键
func (obj *_TNewsMgr) GetBatchFromFID(fIDs []uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id IN (?)", fIDs).Find(&results).Error

	return
}

// GetFromFTitle 通过f_title获取内容 标题
func (obj *_TNewsMgr) GetFromFTitle(fTitle string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_title = ?", fTitle).Find(&results).Error

	return
}

// GetBatchFromFTitle 批量唯一主键查找 标题
func (obj *_TNewsMgr) GetBatchFromFTitle(fTitles []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_title IN (?)", fTitles).Find(&results).Error

	return
}

// GetFromFThumbnail 通过f_thumbnail获取内容 缩略图
func (obj *_TNewsMgr) GetFromFThumbnail(fThumbnail string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_thumbnail = ?", fThumbnail).Find(&results).Error

	return
}

// GetBatchFromFThumbnail 批量唯一主键查找 缩略图
func (obj *_TNewsMgr) GetBatchFromFThumbnail(fThumbnails []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_thumbnail IN (?)", fThumbnails).Find(&results).Error

	return
}

// GetFromFAbstract 通过f_abstract获取内容 摘要
func (obj *_TNewsMgr) GetFromFAbstract(fAbstract string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_abstract = ?", fAbstract).Find(&results).Error

	return
}

// GetBatchFromFAbstract 批量唯一主键查找 摘要
func (obj *_TNewsMgr) GetBatchFromFAbstract(fAbstracts []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_abstract IN (?)", fAbstracts).Find(&results).Error

	return
}

// GetFromFContent 通过f_content获取内容 内容
func (obj *_TNewsMgr) GetFromFContent(fContent string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_content = ?", fContent).Find(&results).Error

	return
}

// GetBatchFromFContent 批量唯一主键查找 内容
func (obj *_TNewsMgr) GetBatchFromFContent(fContents []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_content IN (?)", fContents).Find(&results).Error

	return
}

// GetFromFCat 通过f_cat获取内容 类目
func (obj *_TNewsMgr) GetFromFCat(fCat string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_cat = ?", fCat).Find(&results).Error

	return
}

// GetBatchFromFCat 批量唯一主键查找 类目
func (obj *_TNewsMgr) GetBatchFromFCat(fCats []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_cat IN (?)", fCats).Find(&results).Error

	return
}

// GetFromFAuthor 通过f_author获取内容 作者
func (obj *_TNewsMgr) GetFromFAuthor(fAuthor string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_author = ?", fAuthor).Find(&results).Error

	return
}

// GetBatchFromFAuthor 批量唯一主键查找 作者
func (obj *_TNewsMgr) GetBatchFromFAuthor(fAuthors []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_author IN (?)", fAuthors).Find(&results).Error

	return
}

// GetFromFView 通过f_view获取内容 阅读量
func (obj *_TNewsMgr) GetFromFView(fView uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_view = ?", fView).Find(&results).Error

	return
}

// GetBatchFromFView 批量唯一主键查找 阅读量
func (obj *_TNewsMgr) GetBatchFromFView(fViews []uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_view IN (?)", fViews).Find(&results).Error

	return
}

// GetFromFTag 通过f_tag获取内容 Tag标签，英文逗号分隔
func (obj *_TNewsMgr) GetFromFTag(fTag string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_tag = ?", fTag).Find(&results).Error

	return
}

// GetBatchFromFTag 批量唯一主键查找 Tag标签，英文逗号分隔
func (obj *_TNewsMgr) GetBatchFromFTag(fTags []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_tag IN (?)", fTags).Find(&results).Error

	return
}

// GetFromFIsValid 通过f_is_valid获取内容 1-有效，0-无效
func (obj *_TNewsMgr) GetFromFIsValid(fIsValid bool) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_is_valid = ?", fIsValid).Find(&results).Error

	return
}

// GetBatchFromFIsValid 批量唯一主键查找 1-有效，0-无效
func (obj *_TNewsMgr) GetBatchFromFIsValid(fIsValids []bool) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_is_valid IN (?)", fIsValids).Find(&results).Error

	return
}

// GetFromFOrder 通过f_order获取内容 同级中的顺序，1-n，从上到下
func (obj *_TNewsMgr) GetFromFOrder(fOrder uint8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_order = ?", fOrder).Find(&results).Error

	return
}

// GetBatchFromFOrder 批量唯一主键查找 同级中的顺序，1-n，从上到下
func (obj *_TNewsMgr) GetBatchFromFOrder(fOrders []uint8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_order IN (?)", fOrders).Find(&results).Error

	return
}

// GetFromFReviewStatus 通过f_review_status获取内容 审核状态
func (obj *_TNewsMgr) GetFromFReviewStatus(fReviewStatus int8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_review_status = ?", fReviewStatus).Find(&results).Error

	return
}

// GetBatchFromFReviewStatus 批量唯一主键查找 审核状态
func (obj *_TNewsMgr) GetBatchFromFReviewStatus(fReviewStatuss []int8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_review_status IN (?)", fReviewStatuss).Find(&results).Error

	return
}

// GetFromFReviewTip 通过f_review_tip获取内容 审核提示
func (obj *_TNewsMgr) GetFromFReviewTip(fReviewTip string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_review_tip = ?", fReviewTip).Find(&results).Error

	return
}

// GetBatchFromFReviewTip 批量唯一主键查找 审核提示
func (obj *_TNewsMgr) GetBatchFromFReviewTip(fReviewTips []string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_review_tip IN (?)", fReviewTips).Find(&results).Error

	return
}

// GetFromFSynStatus 通过f_syn_status获取内容 0-未同步，1-已经同步
func (obj *_TNewsMgr) GetFromFSynStatus(fSynStatus uint8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_syn_status = ?", fSynStatus).Find(&results).Error

	return
}

// GetBatchFromFSynStatus 批量唯一主键查找 0-未同步，1-已经同步
func (obj *_TNewsMgr) GetBatchFromFSynStatus(fSynStatuss []uint8) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_syn_status IN (?)", fSynStatuss).Find(&results).Error

	return
}

// GetFromFSynID 通过f_syn_id获取内容 同步资讯ID
func (obj *_TNewsMgr) GetFromFSynID(fSynID int) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_syn_id = ?", fSynID).Find(&results).Error

	return
}

// GetBatchFromFSynID 批量唯一主键查找 同步资讯ID
func (obj *_TNewsMgr) GetBatchFromFSynID(fSynIDs []int) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_syn_id IN (?)", fSynIDs).Find(&results).Error

	return
}

// GetFromFOperatorID 通过f_operator_id获取内容 操作者ID
func (obj *_TNewsMgr) GetFromFOperatorID(fOperatorID int) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_operator_id = ?", fOperatorID).Find(&results).Error

	return
}

// GetBatchFromFOperatorID 批量唯一主键查找 操作者ID
func (obj *_TNewsMgr) GetBatchFromFOperatorID(fOperatorIDs []int) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_operator_id IN (?)", fOperatorIDs).Find(&results).Error

	return
}

// GetFromFUpdateTime 通过f_update_time获取内容 更新时间
func (obj *_TNewsMgr) GetFromFUpdateTime(fUpdateTime uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_update_time = ?", fUpdateTime).Find(&results).Error

	return
}

// GetBatchFromFUpdateTime 批量唯一主键查找 更新时间
func (obj *_TNewsMgr) GetBatchFromFUpdateTime(fUpdateTimes []uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_update_time IN (?)", fUpdateTimes).Find(&results).Error

	return
}

// GetFromFCreateTime 通过f_create_time获取内容 创建时间
func (obj *_TNewsMgr) GetFromFCreateTime(fCreateTime uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_create_time = ?", fCreateTime).Find(&results).Error

	return
}

// GetBatchFromFCreateTime 批量唯一主键查找 创建时间
func (obj *_TNewsMgr) GetBatchFromFCreateTime(fCreateTimes []uint32) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_create_time IN (?)", fCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TNewsMgr) FetchByPrimaryKey(fID uint32) (result TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_id = ?", fID).Find(&result).Error

	return
}

// FetchIndexByFCat  获取多个内容
func (obj *_TNewsMgr) FetchIndexByFCat(fCat string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_cat = ?", fCat).Find(&results).Error

	return
}

// FetchIndexByFTag  获取多个内容
func (obj *_TNewsMgr) FetchIndexByFTag(fTag string) (results []*TNews, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("f_tag = ?", fTag).Find(&results).Error

	return
}
