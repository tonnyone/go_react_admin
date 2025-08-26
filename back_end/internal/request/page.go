package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// PageReq 通用分页请求参数
// 可扩展排序、过滤等字段
type PageReq struct {
	Page      int               `form:"page" json:"page"`
	PageSize  int               `form:"page_size" json:"page_size"`
	SortField string            `form:"sort_field" json:"sort_field"`
	SortOrder string            `form:"sort_order" json:"sort_order"`
	Filters   map[string]string `form:"-" json:"filters"` // 过滤条件，适配 antd 表单
}

// ParsePageReq 从 gin.Context 解析分页参数，带默认值
func ParsePageReq(c *gin.Context) PageReq {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	sortField := c.DefaultQuery("sort_field", "")
	sortOrder := c.DefaultQuery("sort_order", "asc")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	filters := map[string]string{}
	for key, values := range c.Request.URL.Query() {
		if key != "page" && key != "page_size" && key != "sort_field" && key != "sort_order" && len(values) > 0 {
			filters[key] = values[0]
		}
	}
	return PageReq{
		Page:      page,
		PageSize:  pageSize,
		SortField: sortField,
		SortOrder: sortOrder,
		Filters:   filters,
	}
}
