package app

// 数据库，订单数据映射模型
type Order struct {
	Id               uint64 `gorm:"id"`
	OrderNo          string `gorm:"order_no"`
	PublisherOpenId  string `gorm:"publisher_open_id"`
	ConsumerOpenId   string `gorm:"consumer_open_id"`
	ServiceID        uint64 `gorm:"service_id"`
	NationalCodeFull string `gorm:"national_code_full"`
	TelNumber        string `gorm:"tel_number"`
	UserName         string `gorm:"user_name"`
	NationalCode     string `gorm:"national_code"`
	PostalCode       string `gorm:"postal_code"`
	ProvinceName     string `gorm:"province_name"`
	CityName         string `gorm:"city_name"`
	CountyName       string `gorm:"county_name"`
	StreetName       string `gorm:"street_name"`
	DetailInfoNew    string `gorm:"detail_info_new"`
	DetailInfo       string `gorm:"detail_info"`
	BeginDate        string `gorm:"begin_date"`
	EndDate          string `gorm:"end_date"`
	Description      string `gorm:"description"`
	ServiceSnap      string `gorm:"service_snap"`
	Sid              uint64 `gorm:"sid"`
	Status           int    `gorm:"status"`
	Created          string `gorm:"created"`
	Updated          string `gorm:"updated"`
}

// 订单更新参数模型
type OrderUpdateParam struct {
	Id     uint64 `form:"id"`
	Action int    `form:"action" binding:"required,gt=0"`
}

// 订单提交参数模型
type OrderSubmitParam struct {
	OpenId string `form:"openId" json:"openId" binding:"required"`
	Sid    uint64 `form:"sid" json:"sid" binding:"required,gt=0"`
}

// 订单查询参数模型
type OrderQueryParam struct {
	Type     int    `form:"type" json:"type"`
	OpenId   string `form:"openId" json:"openId"`
	Sid      uint64 `form:"sid" binding:"required,gt=0"`
	UserName uint64 `form:"user_name"`
}

// 订单列表传输模型
type OrderList struct {
	Id         uint64      `json:"id"`
	Status     int         `json:"status"`
	TotalPrice float64     `json:"totalPrice"`
	GoodsCount uint        `json:"goodsCount"`
	GoodsItem  []GoodsItem `json:"goodsItem"`
	Created    string      `json:"created"`
}

// 订单服务项传输模型
type GoodsItem struct {
	Id       uint64  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
	Count    int     `json:"count"`
}

// 服务状态更新参数模型
type GoodsStatusUpdateParam struct {
	Id     uint64 `json:"id" from:"id"`
	Action int    `json:"action" from:"action" binding:"required"`
}

type OrderCreateParam struct {
	ServiceID uint64 `json:"service_id"`
	Address   struct {
		NationalCodeFull string `json:"nationalCodeFull"`
		TelNumber        string `json:"telNumber"`
		UserName         string `json:"userName"`
		NationalCode     string `json:"nationalCode"`
		PostalCode       string `json:"postalCode"`
		ProvinceName     string `json:"provinceName"`
		CityName         string `json:"cityName"`
		CountyName       string `json:"countyName"`
		StreetName       string `json:"streetName"`
		DetailInfoNew    string `json:"detailInfoNew"`
		DetailInfo       string `json:"detailInfo"`
	} `json:"address"`
	BeginDate   string `json:"begin_date"`
	Description string `json:"description"`
	EndDate     string `json:"end_date"`
}

type OrderInfo struct {
	ID           uint64       `json:"id"`
	OrderNo      string       `json:"order_no"`
	Price        string       `json:"price"`
	Role         int          `json:"role"`
	ServiceSnap  ServicesInfo `json:"service_snap"`
	AddressSnap  AddressInfo  `json:"address_snap"`
	Status       int          `json:"status"`
	CreateTime   string       `json:"create_time"`
	PublisherTel string       `json:"publisher_tel"`
	BeginDate    string       `json:"begin_date"`
	EndDate      string       `json:"end_date"`
	Tel          string       `json:"tel"`
	Description  string       `json:"description"`
	Publisher    struct {
		ID       uint64      `json:"id"`
		Nickname string      `json:"nickname"`
		Avatar   string      `json:"avatar"`
		Tel      interface{} `json:"tel"`
	} `json:"publisher"`
	Consumer struct {
		ID       uint64      `json:"id"`
		Nickname string      `json:"nickname"`
		Avatar   string      `json:"avatar"`
		Tel      interface{} `json:"tel"`
	} `json:"consumer"`
}

type AddressInfo struct {
	ErrMsg       string `json:"errMsg"`
	CityName     string `json:"cityName"`
	UserName     string `json:"userName"`
	TelNumber    string `json:"telNumber"`
	CountyName   string `json:"countyName"`
	DetailInfo   string `json:"detailInfo"`
	PostalCode   string `json:"postalCode"`
	NationalCode string `json:"nationalCode"`
	ProvinceName string `json:"provinceName"`
}

// 订单列表查询参数模型
type OrderDetailParam struct {
	Id     uint64 `form:"id"`
	Sid    uint64 `form:"sid"`
	OpenId string `form:"openId"`
}
