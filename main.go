package main

import (
	"github.com/gin-gonic/gin"
	"goPjt/controller/mypage"
	"goPjt/controller/user"
	"goPjt/middleware"
)

func main (){
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Use(middleware.RecordUaAndTime)
	APIEngine := engine.Group("/v1")
	{
		auth := APIEngine.Group("/auth")
		{
			auth.POST("/signup", user.Signup)
			auth.POST("/login", user.Login)
			//frontでtoken保存しているcookie消すでlogoutできるのでコメントアウト
			//auth.GET("/logout", user.Logout)
		}
		myPage := APIEngine.Group("/mypage")
		myPage.Use(middleware.IsLogin())
		{
			myPage.GET("/", mypage.LoginTest)
		}

	}
	engine.Run(":3000")
}