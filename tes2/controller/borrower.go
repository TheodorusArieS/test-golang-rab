package controller

import (
	"tes2/domain"
	"github.com/gin-gonic/gin"
	//"fmt"
	"net/http"
	"tes2/service"
	"strconv"
)

func GetBorrower(c *gin.Context){
	result := service.BorrowerService.GetBorrower()
	c.JSON(http.StatusOK,result)
}


func CreateBorrower(c *gin.Context){
	var borrower rab.Borrower
	_=c.ShouldBindJSON(&borrower)
	result := service.BorrowerService.CreateBorrower(borrower)
	c.JSON(http.StatusOK,result)
}

func DeleteBorrower(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("id"),10,64)
	result := service.BorrowerService.DeleteBorrower(id)
	c.JSON(http.StatusOK,result)
}

func EditBorrower(c *gin.Context){
	var borrower rab.Borrower
	_ = c.ShouldBindJSON(&borrower)
	id,_ :=strconv.ParseInt(c.Param("id"),10,64)
	result :=service.BorrowerService.EditBorrower(id,borrower)
	c.JSON(http.StatusOK,result)
}