package service

import (
	"net/http"
	"test/domain/rab"
)

var (
	RabService RabServiceInterface = &rabService{}
)


type RabServiceInterface interface{
	CreateRab(rab.RABList) (*rab.ResResponse)
	GetRabList() []rab.RABList
}

type rabService struct {
	rabs []rab.RABList
}



func (r *rabService) CreateRab(data rab.RABList) (*rab.ResResponse){
	r.rabs = append(r.rabs,data)
	return &rab.ResResponse{
		Status :http.StatusOK,
		Message :"Berhasil Create RAB",
		Data : data,
	}
}

func (r *rabService) GetRabList()[]rab.RABList{
	return r.rabs
}