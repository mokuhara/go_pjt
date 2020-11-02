package mypage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginTest(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "this page neet to login!",
	})
}