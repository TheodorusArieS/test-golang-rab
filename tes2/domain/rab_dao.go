package rab

import (
	//"log"
	"fmt"
	//"database/sql"
	"net/http"
	"tes2/datasource/mysql/rab_db"
	_ "github.com/go-sql-driver/mysql"
)

func (data *RabList) GetRabList()(*RestResponse, *RestResponse){
	stmt, err := rab_db.Client.Query("SELECT * from rab_service.rab")
	if err != nil {
		result1 :=&RestResponse{
			Status: http.StatusBadRequest,
			Data : nil,
			Message : "Error",
		}
		return nil,result1
	}
	result := &RestResponse{
		Status: http.StatusOK,
		Data : nil,
		Message : "Berhasil masuk database",
	}
	for stmt.Next(){
		var name string
		if err :=stmt.Scan(&name); err !=nil {
			fmt.Println(err)
		}
		
		fmt.Println(name)
	}
	return result,nil
	
}