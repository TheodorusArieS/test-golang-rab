package rab

type RABList struct{
	RabID int64 `json:"id"`
	RabName string `json:"name" binding:"required"`
	Comodity string `json:"comodity"`
	Province string `json:"province"`
	City string `json:"city"`
	Quantity int64 `json:"quantity"`
}

type ResResponse struct{
	Status int64 `json:"status"`
	Message string `json:"message"`
	Data RABList `json:"data"`
}