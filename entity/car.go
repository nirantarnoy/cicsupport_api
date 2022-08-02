package entity

type CarCreate struct {
	Id             uint64   `json:"id"`
	AreaId         uint64   `json:"area_id"`
	CarDate        string   `json:"car_date"`
	CarDescription string   `json:"car_description"`
	CarType        uint64   `json:"car_type"`
	CarNonConform  string   `json:"car_non_conform"`
	Status         uint64   `json:"status"`
	CarPhoto       []string `json:"car_photo"`
	CreatedBy      uint64   `json:"created_by"`
	EmpId          uint64   `json:"emp_id"`
}
