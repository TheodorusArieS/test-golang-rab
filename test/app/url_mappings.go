package app

import (
	"test/controller"
)

func mapUrls(){
	router.GET("/",controller.GetRabList)
	router.POST("/rab",controller.CreateRab)
}