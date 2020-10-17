package rab_db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

func init(){
	username := "root"
	password := "admin"
	host := "127.0.0.1:3306"
	dbSchema := "rab_service"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,host,dbSchema)

	Client,err := sql.Open("mysql",dataSourceName)
	if err != nil{
		panic(err.Error())
	}
	log.Println("Database successfully openned")
	if err = Client.Ping();err !=  nil{
		panic(err.Error())
	}

}