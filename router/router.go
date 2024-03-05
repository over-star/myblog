package router

import (
	"github.com/gin-gonic/gin"
	"myblog/app/http/controllers/backend"
	"net/http"
)

func NewRouter() *gin.Engine {
	gin.SetMode("debug")

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)
	group := server.Group("/backend")
	{
		uc := new(backend.UserController)
		group.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, world!",
				"port":    "8080",
			})
		})
		group.POST("/user", uc.Create)
		group.GET("/users", uc.List)
		group.DELETE("/user", uc.Delete)

		ac := new(backend.AuthController)
		group.POST("/login", ac.Login)

	}
	return server
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				//log.
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, "系统内部错误")
		}
	}()
	c.Next()
}
