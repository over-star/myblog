package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int   // 当前页
	PerPage     int   // 每页条数
	TotalPage   int   // 总页数
	TotalCount  int64 // 总条数
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL    string // 用以拼接 URL
	PerPage    int    // 每页条数
	Page       int    // 当前页
	Offset     int    // 数据库读取数据时 Offset 的值
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数 = TotalCount/PerPage
	Sort       string // 排序规则
	Order      string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context，方便调用
}
