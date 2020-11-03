package specialist

import (
	"github.com/gin-gonic/gin"
	"goPjt/repository"
	"goPjt/utils"
	"net/http"
)

func GetAllProfile(c *gin.Context){
	profileRepository := repository.ProfileRepository{}
	profiles, err := profileRepository.GetAll()
	if err != nil{
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status": "ok",
		"data": profiles,
	})
}