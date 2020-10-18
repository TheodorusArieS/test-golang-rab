package service

import (
	"fmt"
	"tes2/domain"
	// "net/http"

)

type RabServiceInterface interface{
	GetRabList()*rab.RestResponse
	CreateRab(rab.RabList) *rab.RestResponse
	EditRab(int64,rab.RabList) *rab.RestResponse
	DeleteRab(int64) *rab.RestResponse
}

var (
	RabService RabServiceInterface = &rabService{}
)

type rabService struct{
	// rabs []rab.RabList
}

// func New() RabServiceInterface{
// 	return &rabService{
// 		rabs :[]rab.RabList{},
// 	}
// }

func (r *rabService) GetRabList()*rab.RestResponse{
	fmt.Println("sudah masuk GetRABList Service")
	dao := &rab.RabList{}
	result, err := dao.GetRabList()
	
	if err != nil{
		return err
	}


	return result
}

func (r *rabService) CreateRab(data rab.RabList) *rab.RestResponse{
	dao :=&rab.RabList{}
	
	result := dao.CreateRabList(data)
	return result
}

func (r *rabService) EditRab(id int64,data rab.RabList) *rab.RestResponse{
	dao :=&rab.RabList{}
	result := dao.EditRab(id,data)
	// result1:=&rab.RestResponse{
	// 	Status:200,
	// 	Data :nil,
	// 	Message:"ada di dalam edit service",
	// }
	return result
}
func (r *rabService) DeleteRab(id int64) *rab.RestResponse{
	fmt.Println(id)
	dao := &rab.RabList{}
	result :=dao.DeleteRab(id)
	return result
}