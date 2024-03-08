package backend

import (
	"github.com/gin-gonic/gin"
	"myblog/app/models"
	"myblog/pkg/database"
	"net/http"
	"strconv"
)

type CategoryController struct{}

func (cc *CategoryController) Create(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 验证成功，创建数据
	_category := models.Category{
		Title: category.Title,
		Pid:   category.Pid,
		Type:  category.Type,
		Link:  category.Link,
		Sort:  category.Sort,
	}
	database.DB.Create(&_category)
	if _category.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": "创建成功",
			"id":    _category.ID,
		})
		return
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error": "创建失败",
	})
}

func (cc *CategoryController) List(c *gin.Context) {
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

func (cc *CategoryController) Delete(c *gin.Context) {
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
