package app

// 服务类目映射模型
type Category struct {
	Id       uint64 `gorm:"primaryKey" json:"id"`       // 类目编号
	Name     string `gorm:"name" json:"name"`           // 类目名称
	ParentId uint64 `gorm:"parent_id" json:"parent_id"` // 父级编号
	Level    uint   `gorm:"level" json:"level"`         // 类目级别
	Sort     uint   `gorm:"sort" json:"sort"`           // 类目排序
	Status   int    `gorm:"status" json:"status"`       // 类目排序
	Sid      uint64 `gorm:"sid" json:"sid"`             // 店铺编号
	Created  string `gorm:"created" json:"created"`     // 创建时间
	Updated  string `gorm:"updated" json:"updated"`     // 更新时间
}

// 类目选项查询参数模型
type CategoryQueryParam struct {
	Sid uint64 `form:"sid"`
}
