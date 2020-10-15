package app

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


var (
	router = gin.Default()
)
func StartApp(){
	mapUrls()
	fmt.Println("Connect to app.go")
	router.Run()
}