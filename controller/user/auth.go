package user

import (
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/repository"
	"goPjt/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Signup(c *gin.Context){
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "bind error")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed create password hash")
		return
	}
	user.Password = string(hash)
	userRepository := repository.UserRepository{}
	err = userRepository.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func Login(c *gin.Context){
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil{
		c.String(http.StatusInternalServerError, "bind error")
		return
	}
	userRepository := repository.UserRepository{}
	matchUser, err := userRepository.Get(user.Email)
	if err != nil{
		c.String(http.StatusInternalServerError, "don't get user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(matchUser.Password), []byte(user.Password))
	if err != nil {
		c.String(http.StatusUnauthorized, "invalid password")
		return
	}
	token, err := service.Generate(matchUser.Id, time.Now())
	if err != nil{
		c.String(http.StatusInternalServerError, "failed create token")
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"token": token,
	})
}


//func Logout(c *gin.Context){}
