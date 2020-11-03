package mypage

import (
	"github.com/gin-gonic/gin"
	"goPjt/model"
	"goPjt/repository"
	"goPjt/utils"
	"net/http"
	"strconv"
)

func LoginTest(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "this page neet to login!",
	})
}

func GetProfile(c *gin.Context){
	paramUserId, _ := strconv.ParseInt(c.Param("id"),10,64)
	profileRepository := repository.ProfileRepository{}
	profile, err := profileRepository.Get(paramUserId)
	if err != nil{
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status": "ok",
		"data": profile,
	})
}


func CreateProfile(c *gin.Context){
	profile := model.Profile{}
	err := c.BindJSON(&profile)
	if err != nil {
		c.String(http.StatusInternalServerError, "bind error")
		return
	}
	profileRepository := repository.ProfileRepository{}
	err = profileRepository.Create(&profile)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed create profile")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UpdateProfile(c *gin.Context){
	profile := model.Profile{}
	err := c.BindJSON(&profile)
	if err != nil {
		apiErr := utils.NewBadRequestError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	profileRepository := repository.ProfileRepository{}
	err = profileRepository.Update(&profile)
	if err != nil {
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})

}

func DeleteProfile(c *gin.Context){
	paramUserId, _ := strconv.ParseInt(c.Param("id"),10,64)
	profileRepository := repository.ProfileRepository{}
	profile, err := profileRepository.Get(paramUserId)
	if err != nil{
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	err = profileRepository.Delete(profile.Id)
	if err != nil {
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}