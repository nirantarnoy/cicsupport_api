package entity

type InspectionRecord struct {
	Module_type_id  uint64 `json:"module_type_id"`
	Inspector_id    uint64 `json:"emp_id"`
	Inspection_date string `json:"inspection_date"`
	Plan_id         uint64 `json:"plan_id"`
	Scored          uint64 `json:"scored"`
	Note            string `json:"note"`
	Created_at      uint64 `json:"created_at"`
	Created_by      uint64 `json:"created_by"`
}
