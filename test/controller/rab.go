package controller

import (
	"test/service"
	"net/http"
	"test/domain/rab"
	"github.com/gin-gonic/gin"
	
)

func CreateRab(c *gin.Context){
	var rab rab.RABList
	_ = c.ShouldBindJSON(&rab) // ini gunanya buat apa ya 
	result := service.RabServiceInterface.CreateRab(rab)
	c.JSON(http.StatusOK,result)
}

func GetRabList(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"ok!",
	})
}