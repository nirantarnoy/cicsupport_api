package entity

type CarListEmp struct {
	Id               uint64 `json:"id"`
	CarNo            string `json:"car_no"`
	AreaId           uint64 `json:"area_id"`
	CarDate          string `json:"car_date"`
	CarDescription   string `json:"car_description"`
	CarType          uint64 `json:"car_type"`
	Status           uint64 `json:"status"`
	CreatedBy        uint64 `json:"created_by"`
	EmpId            uint64 `json:"emp_id"`
	AreaName         string `json:"area_name"`
	IsNew            string `json:"is_new"`
	TargetFinishDate string `json:"target_finish_date"`
	Responsibility   string `json:"responsibility"`
	CarNonConform    string `json:"car_non_conform"`
}
