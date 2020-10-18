package controller

import (
	"fmt"
	"net/http"
	"strconv"
	rab "tes2/domain"
	"tes2/service"

	"github.com/gin-gonic/gin"
)

func GetRabList(c *gin.Context) {
	fmt.Println("sudah ada di get RAB list")
	result := service.RabService.GetRabList()
	c.JSON(http.StatusOK, result)
}

func CreateRab(c *gin.Context) {
	var rab rab.RabList

	if err := c.ShouldBindJSON(&rab); err != nil {
		fmt.Println(&rab)
		result := service.RabService.CreateRab(rab)
		c.JSON(http.StatusOK, result)

	} else {
		fmt.Println(err)
		fmt.Println(&rab)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type format"})
	}

}

func EditRab(c *gin.Context) {
	var rab rab.RabList
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := c.ShouldBindJSON(&rab)
	if err != nil {
		result := service.RabService.EditRab(id, rab)
		c.JSON(http.StatusOK, result)
	}
}

func DeleteRab(c *gin.Context) {
	id,_:= strconv.ParseInt(c.Param("id"),10,64)
	result := service.RabService.DeleteRab(id)
	c.JSON(http.StatusOK,result)
}
