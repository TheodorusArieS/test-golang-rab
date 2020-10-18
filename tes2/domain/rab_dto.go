package rab


type RabList struct{
	Rab_Id string `json:"rab_id"`
	Name string `json:"name" binding:"required"`
	Comodity string `json:"comodity" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
	Description string `json:"description"`
	Borrower Borrower `json:"borrower_id"`
}

type Borrower struct{
	BorrowerId int `json:"borrower_id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	Occupation string `json:"occupation"`
}

type RestResponse struct {
	Status int64 `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}