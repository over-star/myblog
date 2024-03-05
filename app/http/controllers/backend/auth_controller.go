package backend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"myblog/pkg/jwt"
	"net/http"
)

type AuthController struct{}

type LoginMethod struct {
	Password bool `json:"password"`
	QQ       bool `json:"qq"`
	Github   bool `json:"github"`
	Osc      bool `json:"osc"`
}

func (ac *AuthController) Login(c *gin.Context) {
	// 1. 验证表单
	//request := requests.SignupUsingPhoneRequest{}
	//if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
	//	return
	//}
	data, _ := c.GetRawData()
	var post map[string]string
	_ = json.Unmarshal(data, &post)
	jwtToken, err := jwt.CreateJWT(1, post["username"])
	if err != nil {
		log.Fatalln(err)
	}
	//log.Fatalln(post)
	c.JSON(http.StatusOK, gin.H{
		"jwt": jwtToken,
	})
	return
}
