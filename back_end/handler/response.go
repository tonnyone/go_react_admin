package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

// 通用响应结构体
// code: 0=成功, 1=业务失败, 其它=自定义
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
	})
}

func ParamError(c *gin.Context, err error) {
	var msg string
	switch e := err.(type) {
	case *json.SyntaxError:
		body, _ := c.GetRawData()
		msg = "JSON 格式错误: " + e.Error() + ", body: " + string(body) + ", offset: " + itoa64(e.Offset)
	case *json.UnmarshalTypeError:
		msg = "参数类型错误: 字段 " + e.Field + " 期望类型 " + e.Type.String() + ", value: " + e.Value + ", offset: " + itoa64(e.Offset)
	default:
		switch {
		case errors.Is(err, io.EOF):
			msg = "请求体不能为空"
		case errors.As(err, new(validator.ValidationErrors)):
			msg = "参数校验失败: " + err.Error()
		default:
			msg = err.Error()
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"msg":  msg,
	})
}

// itoa64 辅助函数，将 int64 转为字符串
func itoa64(i int64) string {
	return fmt.Sprintf("%d", i)
}
