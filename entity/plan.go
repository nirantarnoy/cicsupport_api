package entity

type Plan struct {
	ID               uint64 `json:"id"`
	Plan_no          string `json:"plan_no"`
	Module_type_id   uint64 `json:"module_type_id"`
	Plan_target_date string `json:"plan_target_date"`
	Status           uint64 `json:"status"`
}
