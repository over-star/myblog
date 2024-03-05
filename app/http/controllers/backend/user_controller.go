package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog/app/models"
	"myblog/pkg/database"
	"net/http"
	"strconv"
)

type UserController struct{}

func (uc *UserController) Create(c *gin.Context) {
	// 1. 验证表单
	//request := requests.SignupUsingPhoneRequest{}
	//if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
	//	return
	//}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientIP := c.ClientIP()
	// 2. 验证成功，创建数据
	_user := models.User{
		Username:    user.Username,
		Password:    user.Password,
		LastLoginIp: clientIP,
	}
	fmt.Println(_user)
	_user.Create()
	if _user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": "创建成功",
			"id":    _user.ID,
		})
		return
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error": "创建失败",
	})
}

func (uc *UserController) List(c *gin.Context) {
	// 1. 验证表单
	//request := requests.SignupUsingPhoneRequest{}
	//if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
	//	return
	//}
	p, _ := strconv.Atoi(c.Query("page"))
	l, _ := strconv.Atoi(c.Query("limit"))
	offset := (p - 1) * l
	//var result []interface{}
	var users []models.User
	result := database.DB.Limit(l).Offset(offset).Select("id", "username", "last_login_ip").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
	})
	return
}

func (uc *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	_user := models.User{
		ID: uint64(id),
	}
	var users []models.User
	result := database.DB.Delete(&_user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
	})
	return
}
