package model

import "github.com/gogf/gf/v2/os/gtime"

// AdminCreateUpdateBase 创建/修改管理员基类
type AdminCreateUpdateBase struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleIds  string `json:"role_ids"`
	UserSalt string `json:"user_salt"`
	IsAdmin  int    `json:"is_admin"`
}

// AdminCreateInput 创建管理员
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// AdminCreateOutput 创建管理员返回结果
type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

// AdminUpdateInput 修改管理员
type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

// AdminGetListInput 获取管理员列表
type AdminGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// AdminGetListOutput 查询列表结果
type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// AdminSearchInput 搜索列表
type AdminSearchInput struct {
	Key        string // 关键字
	Type       string // 管理员模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// AdminSearchOutput 搜索列表结果
type AdminSearchOutput struct {
	List  []AdminSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type AdminGetListOutputItem struct {
	// TODO: 在使用 ScanList 时作为绑定的内容
	//Admin *AdminListItem `json:"admin"`
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	RoleIds   string      `json:"role_ids"`
	IsAdmin   string      `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
	// TODO：之后用到
	//Category *AdminListCategoryItem `json:"category"`
	//User     *AdminListUserItem     `json:"user"`
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}

//// AdminListItem 主要用于列表展示
//type AdminListItem struct {
//	Id        uint        `json:"id"`   // 自增ID
//	Sort      uint        `json:"sort"` // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
