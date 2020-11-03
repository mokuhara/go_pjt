package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/utils"
	"strings"
	"time"
)

type Auth struct {
	UserId int64
	UserType int64
	Iat int64
}

const (
	secret = "VgHIOzK076FnCHA3NYgrJ2fZdfr9y5RRV5XBgwqvgNzNNop/7jC7Bg=="
	userIDKey = "user_id"
	userType = "type"
	iatKey =    "iat"
	expKey =    "exp"
	lifetime =  30 * time.Minute
)

func Generate(user *model.User, now time.Time) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: user.Id,
		userType: user.Type,
		iatKey: now.Unix(),
		expKey: now.Add(lifetime).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func Verify(c *gin.Context) (*Auth){
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
			apiErr := utils.NewBadRequestError("failed token parse")
			c.JSON(apiErr.Status, apiErr)
			return nil
		}
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				apiErr := utils.NewBadRequestError(fmt.Sprintf("not found claims in %s", authToken))
				c.JSON(apiErr.Status, apiErr)
				return nil
			}
			userId, ok := claims[userIDKey].(float64)
			if !ok {
				apiErr := utils.NewBadRequestError(fmt.Sprintf("not found %s in %s", userIDKey, authToken))
				c.JSON(apiErr.Status, apiErr)
				return nil
			}
			userType, ok := claims[userType].(float64)
			if !ok {
				apiErr := utils.NewBadRequestError(fmt.Sprintf("not found %s in %s", userType, authToken))
				c.JSON(apiErr.Status, apiErr)
				return nil
			}
			iat, ok := claims[iatKey].(float64)
			if !ok {
				apiErr := utils.NewBadRequestError(fmt.Sprintf("not found %s in %s", iatKey, authToken))
				c.JSON(apiErr.Status, apiErr)
				return nil
			}
			return &Auth{
				UserId: int64(userId),
				UserType: int64(userType),
				Iat: int64(iat),
			}
		} else {
			apiErr := utils.NewBadRequestError("failed token valid")
			c.JSON(apiErr.Status, apiErr)
			return nil
		}
	} else {
		apiErr := utils.NewBadRequestError("invalid token")
		c.JSON(apiErr.Status, apiErr)
		return nil
	}
}