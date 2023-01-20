package dto

type CloseCarDto struct {
	Id               uint64   `json:"id"`
	CloseDate        string   `json:"close_date" form:"close_date"`
	CloseDescription string   `json:"close_description" form:"close_description"`
	Status           uint64   `json:"status" form:"status"`
	CarPhoto         []string `json:"car_photo"`
	CloseBy          uint64   `json:"close_by" form:"close_by"`
}
