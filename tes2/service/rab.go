package service

import (
	"fmt"
	"tes2/domain"

)

type RabServiceInterface interface{
	GetRabList()*rab.RestResponse
	CreateRab(rab.RabList) rab.RabList	
}

var (
	RabService RabServiceInterface = &rabService{}
)

type rabService struct{
	rabs []rab.RabList
}

func New() RabServiceInterface{
	return &rabService{
		rabs :[]rab.RabList{},
	}
}

func (r *rabService) GetRabList()*rab.RestResponse{
	fmt.Println("sudah masuk GetRABList Service")
	dao := &rab.RabList{}
	result, err := dao.GetRabList()
	if err != nil{
		return err
	}

	return result
}

func (r *rabService) CreateRab(data rab.RabList) rab.RabList{
	borrower :=BorrowerService.GetBorrower()
	if borrower !=nil{
		data.Borrower = borrower[0]
	}
	
	r.rabs = append(r.rabs,data)
	return data
}