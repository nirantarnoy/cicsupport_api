package dto

type CarCreateDto struct {
	Id             uint64   `json:"id"`
	AreaId         uint64   `json:"area_id" form:"area_id" binding:"required"`
	CarDate        string   `json:"car_date" form:"car_date"`
	CarDescription string   `json:"car_description" form:"car_description"`
	CarType        uint64   `json:"car_type" form:"car_type"`
	CarNonConform  string   `json:"car_non_conform" form:"car_non_conform"`
	Status         uint64   `json:"status" form:"status"`
	CarPhoto       []string `json:"car_photo"`
	CreatedBy      uint64   `json:"created_by" form:"created_by"`
	EmpId          uint64   `json:"emp_id" form:"emp_id"`
}
