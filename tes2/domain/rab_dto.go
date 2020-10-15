package rab


type RabList struct{
	Name string `json:"name"`
	Comodity string `json:"comodity"`
	Quantity string `json:"quantity"`
	Description string `json:"description"`
	Borrower Borrower `json:"borrower"`
}

type Borrower struct{
	BId int `json:"borrower_id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	Occupation string `json:"occupation"`
}