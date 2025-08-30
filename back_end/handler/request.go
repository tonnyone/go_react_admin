package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/tonnyone/go_react_admin/internal/dto"
	"github.com/tonnyone/go_react_admin/internal/logger"
)

// ParsePageReq 从 gin.Context 解析分页参数，带默认值
func ParsePageReq(c *gin.Context) (dto.Pager, error) {
	var req dto.Pager
	logger.Infof("parse page request")
	if err := c.ShouldBindQuery(&req); err != nil {
		return dto.Pager{}, err
	}
	logger.Infof("parse pager: %+v", req)
	filters := map[string]string{}
	sorter := map[string]dto.Order{}
	for key, values := range c.Request.URL.Query() {
		if key == "sort_by" {
			pair := strings.Split(values[0], ":")
			if len(pair) == 2 {
				if strings.ToLower(pair[1]) == "asc" || strings.ToLower(pair[1]) == "desc" {
					sorter[strcase.ToSnake(pair[0])] = dto.Order(pair[1])
				}
			}
		}
		if key != "current" && key != "page_size" && key != "sort_by" && len(values) > 0 {
			filters[strcase.ToSnake(key)] = values[0]
		}
	}
	logger.Infof("pager: %+v, filters: %+v, sorter: %+v", req, filters, sorter)
	req.Sorter = sorter
	req.Filter = filters
	return req, nil
}
