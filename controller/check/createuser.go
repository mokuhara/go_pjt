package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goPjt/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func SignUp(c *gin.Context) {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		apiErr := utils.NewInternalServerError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	PORT := os.Getenv("PORT")
	url:= fmt.Sprintf("http://localhost:%s/v1/auth/signup", PORT)

	type SignUp struct {
		Email    string `json:email`
		Password string `json:password`
		Type int `json:int`
	}

	randEmail := fmt.Sprintf("%s@gmail.com", RandomString(10))
	signUp := SignUp{Email:randEmail,Password:"password",Type:0}
	body, _ := createJson(signUp)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		apiErr := utils.NewBadRequestError(fmt.Errorf("failed create request"))
		c.JSON(apiErr.Status, apiErr)
		return
	}
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr := utils.NewBadRequestError(fmt.Errorf("failed request post"))
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status": "ok",
		"data": string(byteArray),
	})
}

func createJson(anyStruct interface{}) (*bytes.Buffer, error){
	json, err := json.Marshal(anyStruct)
	if err != nil {
		return nil, fmt.Errorf("faild create json")
	}
	return bytes.NewBuffer(json), nil
}


func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}