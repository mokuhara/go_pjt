package admin

import (
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/repository"
	"goPjt/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context){
	userRepository := repository.UserRepository{}
	users, err := userRepository.GetAll()
	if err != nil{
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status": "ok",
		"data": users,
	})
}

func UpdateUser(c *gin.Context){
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		apiErr := utils.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	userRepository := repository.UserRepository{}
	err = userRepository.Update(&user)
	if err != nil {
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func DeleteUser(c *gin.Context){
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10,0)
	if err != nil {
		apiErr := utils.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	userRepository := repository.UserRepository{}
	err = userRepository.Delete(intId)
	if err != nil {
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func CreateUser(c *gin.Context){
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
	if err != nil {
		c.String(http.StatusInternalServerError, "failed create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
