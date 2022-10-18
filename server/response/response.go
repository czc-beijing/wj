package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"error_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PageResult 分页结果返回
type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// PageResult 分页结果返回
type AppPageResult struct {
	Total       int64       `json:"total"`
	PerPage     int64       `json:"per_page"`
	CurrentPage int64       `json:"current_page"`
	LastPage    int64       `json:"last_page"`
	List        interface{} `json:"data"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{0, message, data})
}

// FailedData 请求失败返回
func FailedData(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, data})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}

// SuccessPage 请求成功返回分页结果
func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
	page := &PageResult{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{0, message, page})
}

// SuccessAppPage 请求成功返回分页结果
func SuccessAppPage(message string, data interface{}, rows int64, currPage int, pageSize int, c *gin.Context) {
	lastPage := rows/int64(pageSize) + 1
	page := &AppPageResult{Total: rows, List: data, CurrentPage: int64(currPage), PerPage: int64(currPage), LastPage: lastPage}
	c.JSON(http.StatusOK, Response{0, message, page})
}

// AppFailed 请求失败返回
func AppFailed(httpCode int, code int, message string, c *gin.Context) {
	c.JSON(httpCode, Response{code, message, 0})
}
