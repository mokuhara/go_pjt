package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	secret = "VgHIOzK076FnCHA3NYgrJ2fZdfr9y5RRV5XBgwqvgNzNNop/7jC7Bg=="
	userIDKey = "user_id"
	iatKey =    "iat"
	expKey =    "exp"
	lifetime =  30 * time.Minute
)

func Generate(userId int64, now time.Time) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: userId,
		iatKey: now.Unix(),
		expKey: now.Add(lifetime).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func Verify(c *gin.Context) (string, error){
	authHeader := c.Request.Header["Authorization"][0]
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2{
		authToken := bearerToken[1]
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token verify error")
			}
			return []byte(secret), nil
		})
		if err != nil {
			return "", errors.New("")
		}
		if token.Valid {
			return "ok", nil
		} else {
			return "",  errors.New("failed token valid")
		}
	} else {
		return "", errors.New("invalid token")
	}
}