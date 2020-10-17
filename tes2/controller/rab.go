package controller

import (
	"tes2/domain"
	"net/http"
	"tes2/service"
	"fmt"
	"github.com/gin-gonic/gin"
)




func GetRabList(c *gin.Context){
	fmt.Println("sudah ada di get RAB list")
	result :=service.RabService.GetRabList()
	c.JSON(http.StatusOK,result)
}

func CreateRab(c *gin.Context){
	var rab rab.RabList
	if err := c.ShouldBindJSON(&rab) ; err != nil{

		fmt.Println(c,"Ada di create RAb")
		result := service.RabService.CreateRab(rab)
		c.JSON(http.StatusOK,result)

	} else {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid type format"})
	}

}