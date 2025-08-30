package dto

// 可扩展排序、过滤等字段
type Pager struct {
	// antd ProTable, antd Table 的分页参数
	Current  int `form:"current,default=1"`
	PageSize int `form:"page_size,default=10"`
	// antd ProTable, antd Table 的排序参数
	Sorter map[string]Order
	// antd ProTable, antd Table 的过滤参数
	Filter map[string]string
}

type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)
