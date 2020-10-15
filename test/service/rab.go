package service

import (
	"net/http"
	"test/domain/rab"
)

type RabServiceInterface interface{
	CreateRab(rab.RABList) (*rab.ResResponse)
	GetRabList() []rab.RABList
}

type RabService struct {
	rabs []rab.RABList
}

func New() RabServiceInterface{
	return &RabService{
		rabs : []rab.RABList{},
	}
}

func (r *RabService) CreateRab(data rab.RABList) (*rab.ResResponse){
	r.rabs = append(r.rabs,data)
	return &rab.ResResponse{
		Status :http.StatusOK,
		Message :"Berhasil Create RAB",
		Data : data,
	}
}

func (r *RabService) GetRabList()[]rab.RABList{
	return r.rabs
}