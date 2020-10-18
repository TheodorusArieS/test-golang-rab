package rab

import (
	//"log"
	"database/sql"
	"fmt"
	"strconv"
	"net/http"
	// "tes2/datasource/mysql/rab_db"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func (data *RabList) GetRabList() (*RestResponse, *RestResponse) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rab_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM rab")
	var results []RabList
	for result.Next(){
		var res RabList
		if err := result.Scan(&res.Name);err !=nil{
			results = append(results,res)
			return nil, &RestResponse{
				Status:400,
				Data:results,
				Message :"Gagal ambil data",
			}
		}
		fmt.Println(result)
		results = append(results,res)
	}
	result1:=&RestResponse{
		Status:200,
		Data:results,
		Message:"Berhasil ambil data",
	}
	fmt.Println("ada di dalam dao")
	return result1,nil

}

func (data *RabList) CreateRabList(r RabList) *RestResponse{
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rab_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO rab(name,comodity,quantity,description) VALUES(?,?,?,?)")
	if err != nil{
		fmt.Println("Ada salah di stmt")
		return nil
	}
	defer stmt.Close()
	newQty,_ := strconv.ParseInt(r.Quantity,10,64)
	_, saveErr :=stmt.Exec(r.Name,r.Comodity,newQty,r.Description)
	if saveErr != nil{
		fmt.Println("Gagal di result")
		return nil
	}
	result1 :=&RestResponse{
		Status:http.StatusOK,
		Data:r,
		Message:"Create New RAB",
	}
	return result1
}

func (data *RabList) EditRab(id int64,rab RabList)*RestResponse{
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rab_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(rab.Name,id)
	stmt,err :=db.Prepare("UPDATE rab SET name = ? WHERE rab_id = ?")
	defer stmt.Close()
	if err != nil{
		fmt.Println("ada salah di smtm")
		return nil
	}
	update,err :=stmt.Exec(rab.Name,id)
	if err !=nil{
		return &RestResponse{
			Status:http.StatusBadRequest,
			Data:nil,
			Message:"gagal update",
		}
	}
	return &RestResponse{
		Status:http.StatusOK,
		Data:update,
		Message:"Berhasil update",
	}
}

func (data *RabList) DeleteRab(id int64) *RestResponse{
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/rab_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt,err := db.Prepare("DELETE FROM rab WHERE rab_id = ?")
	if err !=nil{
		fmt.Println("Error di stmt")
		return nil
	}
	defer stmt.Close()
	delete,err := stmt.Exec(id)
	if err !=nil{
		fmt.Println("error di delete")
		return nil
	}
	return &RestResponse{
		Status:http.StatusOK,
		Data:delete,
		Message:"Berhasil Delete",
	}
}