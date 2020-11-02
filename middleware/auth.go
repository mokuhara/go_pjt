package middleware

import (
	"github.com/gin-gonic/gin"
	"goPjt/service"
	"net/http"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		res, err := service.Verify(c)
		if err != nil{
			c.String(http.StatusUnauthorized, "verify token error")
		}

		if res != "ok" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "ng",
				"data": "invalid token",
			})
			c.Abort()
		}
	}
}