package middleware

import (
	"github.com/gin-gonic/gin"
	"goPjt/service"
	"net/http"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenService := service.TokenService{}
		res := tokenService.Verify(c)
		if res == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "invalid token",
			})
			c.Abort()
		}
	}
}

func IsAdmin() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenService := service.TokenService{}
		res := tokenService.Verify(c)

		if res == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "invalid token",
			})
			c.Abort()
		} else if res.UserType != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "insufficient authority",
			})
			c.Abort()
		}
	}
}