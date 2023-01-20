package dto

type PersonTrans struct {
	TeamId string `json:"teamid" form:"teamid" binding:"required"`
	EmpId  string `json:"empid" form:"empid" binding:"required"`
}
