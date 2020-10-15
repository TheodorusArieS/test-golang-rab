package application

import (
	"tes2/controller"
)




func mapUrls(){
	rab := router.Group("/rab")
	borrower := router.Group("/borrower")
	{
		rab.GET("/",controller.GetRabList)
		rab.POST("/",controller.CreateRab)
		borrower.GET("/",controller.GetBorrower)
		borrower.POST("/",controller.CreateBorrower)
		borrower.DELETE("/:id",controller.DeleteBorrower)
		borrower.PUT("/:id",controller.EditBorrower)
	}
}