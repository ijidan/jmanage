package model

// BlogArticle 文章管理
type BlogArticle struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	TagID      int    `json:"tag_id"` // 标签ID
	Title      string `json:"title"`  // 文章标题
	Desc       string `json:"desc"`   // 简述
	Content    string `json:"content"`
	CreatedOn  int    `json:"created_on"`
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedOn uint32 `json:"modified_on"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	DeletedOn  uint32 `json:"deleted_on"`
	State      int8   `json:"state"` // 状态 0为禁用1为启用
}

// BlogAuth [...]
type BlogAuth struct {
	ID       uint32 `gorm:"primaryKey" json:"id"`
	Username string `json:"username"` // 账号
	Password string `json:"password"` // 密码
}

// BlogTag 文章标签管理
type BlogTag struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`        // 标签名称
	CreatedOn  int    `json:"created_on"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedOn uint32 `json:"modified_on"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	DeletedOn  int    `json:"deleted_on"`
	State      uint8  `json:"state"` // 状态 0为禁用、1为启用
}
