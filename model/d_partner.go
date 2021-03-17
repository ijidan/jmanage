package model

// TNews [...]
type TNews struct {
	FID           uint32 `gorm:"primaryKey" json:"f_id"`   // 主键
	FTitle        string `json:"f_title"`                  // 标题
	FThumbnail    string `json:"f_thumbnail"`              // 缩略图
	FAbstract     string `json:"f_abstract"`               // 摘要
	FContent      string `json:"f_content"`                // 内容
	FCat          string `gorm:"index:f_cat" json:"f_cat"` // 类目
	FAuthor       string `json:"f_author"`                 // 作者
	FView         int    `json:"f_view"`                   // 阅读量
	FTag          string `gorm:"index:f_tag" json:"f_tag"` // Tag标签，英文逗号分隔
	FIsValid      bool   `json:"f_is_valid"`               // 1-有效，0-无效
	FOrder        uint8  `json:"f_order"`                  // 同级中的顺序，1-n，从上到下
	FReviewStatus uint8  `json:"f_review_status"`          // 审核状态
	FReviewTip    string `json:"f_review_tip"`             // 审核提示
	FSynStatus    uint8  `json:"f_syn_status"`             // 0-未同步，1-已经同步
	FSynID        uint32 `json:"f_syn_id"`                 // 同步资讯ID
	FOperatorID   int    `json:"f_operator_id"`            // 操作者ID
	FUpdateTime   uint32 `json:"f_update_time"`            // 更新时间
	FCreateTime   uint32 `json:"f_create_time"`            // 创建时间
}

// TOperateLog 系统操作记录表
type TOperateLog struct {
	FID     int    `gorm:"primaryKey" json:"f_id"`
	FUId    int    `json:"f_uid"`    // 操作用户ID
	FName   string `json:"f_name"`   // 操作用户名
	FCtime  int    `json:"f_ctime"`  // 操作时间
	FMethod string `json:"f_method"` // 操作方式；get，post
	FURL    string `json:"f_url"`    // 操作URL
	FData   string `json:"f_data"`   // 操作数据
}

// TOperateMenu 系统菜单表
type TOperateMenu struct {
	ID          int    `gorm:"primaryKey;index:id" json:"id"`
	Name        string `json:"name"`
	Module      string `gorm:"index:module" json:"module"`         // 模块
	Controller  string `gorm:"index:controller" json:"controller"` // 控制器
	Function    string `gorm:"index:function" json:"function"`     // 方法
	Parameter   string `json:"parameter"`                          // 参数
	Description string `json:"description"`                        // 描述
	IsDisplay   int    `gorm:"index:is_display" json:"is_display"` // 1显示在左侧菜单2只作为节点
	Type        int    `gorm:"index:type" json:"type"`             // 1权限节点2普通节点
	Pid         int    `json:"pid"`                                // 上级菜单0为顶级菜单
	CreateTime  int    `json:"create_time"`
	UpdateTime  int    `json:"update_time"`
	Icon        string `json:"icon"`    // 图标
	IsOpen      int    `json:"is_open"` // 0默认闭合1默认展开
	Orders      int    `json:"orders"`  // 排序值，越小越靠前
}

// TOperateRole [...]
type TOperateRole struct {
	ID          int    `gorm:"primaryKey;index:id" json:"id"`
	Name        string `gorm:"index:name" json:"name"`
	Permissions string `json:"permissions"` // 权限菜单
	CreateTime  int    `gorm:"index:create_time" json:"create_time"`
	UpdateTime  int    `json:"update_time"`
	Desc        string `json:"desc"` // 备注
}

// TOperateUser [...]
type TOperateUser struct {
	ID          int    `gorm:"primaryKey" json:"id"`  // 用户id
	Passwd      string `json:"passwd"`                // 账户密码
	Email       string `json:"email"`                 // 工作邮箱
	NameEn      string `gorm:"unique" json:"name_en"` // 英文名
	NameCn      string `json:"name_cn"`               // 中文名
	Permissions string `json:"permissions"`           // 权限ID集合，逗号分隔
	RoleID      int    `json:"role_id"`               // 角色ID
	CreateTime  int    `json:"create_time"`           // 创建时间
	UpdateTime  int    `json:"update_time"`           // 更新时间
	LowerUserID string `json:"lower_user_id"`         // 下级用户ID，逗号分隔
	IsDisabled  int8   `json:"is_disabled"`           // 是否禁用
}
