package main

import (
	"github.com/gin-gonic/gin"
	"goPjt/controller/mypage"
	"goPjt/controller/user"
	"goPjt/controller/admin"
	"goPjt/middleware"
)

func main (){
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Use(middleware.RecordUaAndTime)
	APIEngine := engine.Group("/v1")
	{
		authEngine := APIEngine.Group("/auth")
		{
			authEngine.POST("/signup", user.Signup)
			authEngine.POST("/login", user.Login)
			//frontでtoken保存しているcookie消すでlogoutできるのでコメントアウト
			//auth.GET("/logout", user.Logout)
		}
		myPageEngine := APIEngine.Group("/mypage")
		myPageEngine.Use(middleware.IsLogin())
		{
			myPageEngine.GET("/", mypage.LoginTest)
		}
		adminEngine := APIEngine.Group("/admin")
		adminEngine.Use(middleware.IsAdmin())
		{
			adminEngine.GET("/users", admin.GetUsers)
		}

	}
	engine.Run(":3000")
}
