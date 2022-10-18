package web

// 服务类目映射模型
type Category struct {
	Id       uint64 `gorm:"primaryKey"` // 类目编号
	Name     string `gorm:"name"`       // 类目名称
	ParentId uint64 `gorm:"parent_id"`  // 父级编号
	Level    uint   `gorm:"level"`      // 类目级别
	Sort     uint   `gorm:"sort"`       // 类目排序
	Status   int    `gorm:"status"`     // 类目状态
	Created  string `gorm:"created"`    // 创建时间
	Updated  string `gorm:"updated"`    // 更新时间
	Sid      uint64 `gorm:"sid"`        // 店铺编号
}

// 服务类目创建参数模型
type CategoryCreateParam struct {
	Name     string `json:"name" binding:"required"`
	ParentId uint64 `json:"parentId" binding:"required,gt=0"`
	Level    uint   `json:"level" binding:"required,gt=0"`
	Sort     uint   `json:"sort" binding:"required,gt=0"`
	Sid      uint64 `json:"sid" binding:"required,gt=0"`
}

// 服务类目删除参数模型
type CategoryDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 服务类目更新参数模型
type CategoryUpdateParam struct {
	Id   uint64 `json:"id" binding:"required,gt=0"`
	Name string `json:"name"`
	Sort uint   `json:"sort" binding:"required,gt=0"`
}

// 服务类目查询参数模型
type CategoryQueryParam struct {
	Sid uint64 `form:"sid""`
}

// 服务类目列表传输模型
type CategoryList struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// 服务类目列表传输模型
type HomemakingCategoryList struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// 服务类目选项传输模型
type CategoryOption struct {
	Value    uint64           `json:"value"`
	Label    string           `json:"label"`
	Children []CategoryOption `json:"children"`
}
