package middleware

import (
	"github.com/gin-gonic/gin"
	"goPjt/service"
	"net/http"
	"strconv"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		paramUserId, _ := strconv.ParseInt(c.Param("id"),10,64)
		tokenService := service.TokenService{}
		res := tokenService.Verify(c)

		if res == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "invalid token",
			})
			c.Abort()
		} else if res.UserId != paramUserId {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "invalid userId",
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