package application

import (
	"tes2/controller"
)




func mapUrls(){
	rab := router.Group("/rab")
	{
		rab.GET("/",controller.GetRabList)
		rab.POST("/",controller.CreateRab)
	}
	borrower := router.Group("/borrower")
	{
		borrower.GET("/",controller.GetBorrower)
		borrower.POST("/",controller.CreateBorrower)
		borrower.DELETE("/:id",controller.DeleteBorrower)
		borrower.PUT("/:id",controller.EditBorrower)
	}
}