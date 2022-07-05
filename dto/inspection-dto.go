package dto

type InspectionCreateDTO struct {
	Module_type_id  uint64 `json:"module_type_id" form:"module_type_id" binding:"required"`
	Inspector_id    uint64 `json:"emp_id" form:"emp_id" binding:"required"`
	Inspection_date string `json:"inspection_date" form:"inspection_date" binding:"required"`
	Plan_id         uint64 `json:"plan_id" form:"plan_id" binding:"required"`
	Scored          uint64 `json:"scored" form:"scored" binding:"required"`
	Note            string `json:"note"`
	Created_at      uint64 `json:"created_at"`
	Created_by      uint64 `json:"created_by" binding:"required"`
}

type InspectionUpdateDTO struct {
	ID         uint64 `json:"id"`
	Scored     uint64 `json:"scored" form:"scored" binding:"required"`
	Note       string `json:"note"`
	Created_at uint64 `json:"created_at"`
	Created_by uint64 `json:"created_by" binding:"required"`
}
