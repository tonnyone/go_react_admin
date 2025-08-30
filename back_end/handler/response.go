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

// Code 响应码类型（字符串枚举）
type Code int

// 统一响应 code 常量
const (
	OK          Code = 0   // 成功
	FAIL        Code = 500 // 业务/系统错误
	PARAM_ERROR Code = 400 // 参数错误
)

// Resp 统一响应结构体
type Resp[T any] struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageData[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

// JSON 统一响应输出
func ResponseSuccss[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Resp[T]{
		Code: OK,
		Msg:  "success",
		Data: data,
	})
}

func ResponseFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAIL,
		"msg":  msg,
	})
}

func ResponseParamError(c *gin.Context, err error) {
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
		"code": PARAM_ERROR,
		"msg":  msg,
	})
}

// itoa64 辅助函数，将 int64 转为字符串
func itoa64(i int64) string {
	return fmt.Sprintf("%d", i)
}
