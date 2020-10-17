package service

import (
	// "fmt"
	rab "tes2/domain"
	//"reflect"
)

type BorrowerServiceInterface interface {
	GetBorrower() []rab.Borrower
	CreateBorrower(rab.Borrower) rab.Borrower
	DeleteBorrower(int64) *rab.RestResponse
	EditBorrower(int64, rab.Borrower) rab.Borrower
}

type borrowerService struct {
	listBorrower []rab.Borrower
}

var (
	BorrowerService BorrowerServiceInterface = &borrowerService{}
)

func (b *borrowerService) GetBorrower() []rab.Borrower {
	return b.listBorrower
}

func (b *borrowerService) CreateBorrower(data rab.Borrower) rab.Borrower {
	b.listBorrower = append(b.listBorrower, data)
	return data
}

func (b *borrowerService) DeleteBorrower(id int64) *rab.RestResponse {
	var deleteStatus bool = false
	
	var newList []rab.Borrower
	if len(b.listBorrower) == 0 {
		result := &rab.RestResponse{
			Status:200,
			Data:nil,
			Message:"Data Kosong tidak bisa delete",
		}
		return result
	} else {
		for i, _ := range b.listBorrower {
			if int64(b.listBorrower[i].BorrowerId) == id {
				deleteStatus = true
			} else {
				newList = append(newList, b.listBorrower[i])
			}
		}
		if deleteStatus == true {
			b.listBorrower = newList
			result := &rab.RestResponse{
				Status:200,
				Data:nil,
				Message:"Berhasil Delete",
			}
			return result
		}
		result := &rab.RestResponse{
			Status:200,
			Data:nil,
			Message:"Data Not Found",
		}
		return result

	}
}

func (b *borrowerService) EditBorrower(id int64, data rab.Borrower) rab.Borrower {
	var editStatus bool = false
	for i, _ := range b.listBorrower {
		if int64(b.listBorrower[i].BorrowerId) == id{
			editStatus = true
			b.listBorrower[i].Age = data.Age
			b.listBorrower[i].Firstname = data.Firstname
			b.listBorrower[i].Lastname = data.Lastname
			b.listBorrower[i].Occupation = data.Occupation
		}
	}
	if editStatus == true {
		return data
	}
	// harus bisa di kasih JSON juga yang ngembaliin message
	return data
}
