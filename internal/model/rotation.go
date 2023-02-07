package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation-id"`
}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationGetListInput 获取轮播图列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// RotationSearchInput 搜索列表
type RotationSearchInput struct {
	Key        string // 关键字
	Type       string // 轮播图模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationSearchOutput 搜索列表结果
type RotationSearchOutput struct {
	List  []RotationSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type RotationGetListOutputItem struct {
	// TODO: 在使用 ScanList 时作为绑定的内容
	//Rotation *RotationListItem `json:"rotation"`
	Id        uint        `json:"id"`   // 自增ID
	Sort      uint        `json:"sort"` // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
	// TODO：之后用到
	//Category *RotationListCategoryItem `json:"category"`
	//User     *RotationListUserItem     `json:"user"`
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}

// RotationListItem 主要用于列表展示
type RotationListItem struct {
	Id        uint        `json:"id"`   // 自增ID
	Sort      uint        `json:"sort"` // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
