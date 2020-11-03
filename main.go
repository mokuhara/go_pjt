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
		specialistEngine := APIEngine.Group("/specialist")
		{
			specialistEngine.GET("/")
		}
		myPageEngine := APIEngine.Group("/mypage/:id")
		myPageEngine.Use(middleware.IsLogin())
		{
			myPageEngine.GET("/", mypage.LoginTest)
			myPageUserEngine := myPageEngine.Group("/user")
			{
				myPageUserEngine.GET("/", mypage.GetProfile)
				myPageUserEngine.POST("/create", mypage.CreateProfile)
				myPageUserEngine.PUT("/update", mypage.UpdateProfile)
				myPageUserEngine.DELETE("/delete", mypage.DeleteProfile)
			}

		}
		adminEngine := APIEngine.Group("/admin")
		adminEngine.Use(middleware.IsAdmin())
		{
			userEngine := adminEngine.Group("/user")
			{
				userEngine.GET("/index", admin.GetUsers)
				userEngine.PUT("/update", admin.UpdateUser)
				userEngine.DELETE("/delete", admin.DeleteUser)
				userEngine.POST("/create", admin.CreateUser)
			}
		}
	}
	engine.Run(":3000")
}
