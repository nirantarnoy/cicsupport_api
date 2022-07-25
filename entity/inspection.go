package entity

type InspectionRecord struct {
	ModuleTypeId uint64 `json:"module_type_id"`
	PlanId       uint64 `json:"plan_id"`
	TransDate    string `json:"trans_date"`
	EmpId        uint64 `json:"emp_id"`
	AreaGroupId  uint64 `json:"area_group_id"`
	AreaId       uint64 `json:"area_id"`
	TeamId       uint64 `json:"team_id"`
	TopicId      uint64 `json:"topic_id"`
	TopicItemId  uint64 `json:"topic_item_id"`
	Score        uint64 `json:"score"`
	Status       uint64 `json:"status"`
	Note         string `json:"note"`
	Created_at   uint64 `json:"created_at"`
	Created_by   uint64 `json:"created_by"`
}
