package user

import (
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

}

func Logout(c *gin.Context){

}
