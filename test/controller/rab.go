package controller

import (
	"test/service"
	"net/http"
	"test/domain/rab"
	"github.com/gin-gonic/gin"
	
)


func CreateRab(c *gin.Context){
	var rab rab.RABList
	 _ = c.ShouldBindJSON(&rab)
	result := service.RabService.CreateRab(rab)
	c.JSON(http.StatusOK,result)
}

func GetRabList(c *gin.Context){
	result := service.RabService.GetRabList()
	c.JSON(http.StatusOK,result)
}