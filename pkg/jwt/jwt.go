package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"strings"
	"time"
	"unicode"
)

type UserClaims struct {
	*jwt.RegisteredClaims
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

const (
	userTokenHeader = "Authorization"
	userTokenParam  = "_user_token"

	expireSeconds = 3600
	issuer        = "blog-go"
	secret        = "blog-go"
)

func CreateJWT(Id int64, username string) (string, error) {
	var (
		expiredAt = time.Now().Add(time.Duration(expireSeconds) * time.Second)
	)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        cast.ToString(Id),
		},
		Username: username,
	})
	return claims.SignedString([]byte(secret))
}

func parseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("couldn't handle this token")
	} else {
		return nil, errors.New("couldn't handle this token")
	}
}

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func GetUser(c *gin.Context) (user *UserClaims) {
	token := c.Request.Header.Get(userTokenHeader)
	if !IsBlank(token) {
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
	}
	user, _ = parseToken(token)
	return
}
