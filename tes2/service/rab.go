package service

import (
	"tes2/domain"

)

type RabServiceInterface interface{
	GetRabList()[]rab.RabList
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

func (r *rabService) GetRabList()[]rab.RabList{
	return r.rabs
}

func (r *rabService) CreateRab(data rab.RabList) rab.RabList{
	borrower :=BorrowerService.GetBorrower()
	data.Borrower = borrower[0]
	r.rabs = append(r.rabs,data)
	return data
}