package dto

type PlanUpdateDTO struct {
	ID               uint64 `json:"id" form:"id" binding:"required"`
	Action_plan_date string `json:"action_plan_date"`
	Status           uint64 `json:"status"`
}
type PlanFindDTO struct {
	Module_type_id uint64 `json:"module_type_id"`
	Status         uint64 `json:"status"`
}
