package user

import (
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/repository"
	"goPjt/service"
	"goPjt/utils"
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
	purePassword := user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed create password hash")
		return
	}
	user.Password = string(hash)
	userRepository := repository.UserRepository{}
	err = userRepository.Create(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed create user")
		return
	}
	//passwordをhash化前に戻す
	user.Password = purePassword
	token := createToken(&user, c)
	if err != nil{
		c.String(http.StatusInternalServerError, "failed create token")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"token": token,
	})

}

func Login(c *gin.Context){
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil{
		c.String(http.StatusInternalServerError, "bind error")
		return
	}
	token := createToken(&user, c)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"token": token,
	})
}

func createToken(user *model.User, c *gin.Context) string{
	userRepository := repository.UserRepository{}
	matchUser, err := userRepository.Get(user.Email)
	if err != nil{
		apiErr := utils.NewBadRequestError("don't get user")
		c.JSON(apiErr.Status, apiErr)
		return ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(matchUser.Password), []byte(user.Password))
	if err != nil {
		apiErr := utils.NewBadRequestError("invalid password")
		c.JSON(apiErr.Status, apiErr)
		return ""
	}
	token, err := service.Generate(matchUser, time.Now())
	if err != nil{
		apiErr := utils.NewBadRequestError("failed create token")
		c.JSON(apiErr.Status, apiErr)
		return ""
	}
	return token
}


//func Logout(c *gin.Context){}
