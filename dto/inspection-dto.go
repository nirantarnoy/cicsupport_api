package dto

type InspectionCreateDTO struct {
	ModuleTypeId uint64 `json:"module_type_id" form:"module_type_id"`
	PlanId       uint64 `json:"plan_id" form:"plan_id"`
	TeamId       uint64 `json:"team_id" from:"team_id"`
	TransDate    string `json:"trans_date" form:"trans_date"`
	EmpId        uint64 `json:"emp_id" form:"emp_id"`
	AreaGroupId  uint64 `json:"area_group_id" form:"area_group_id"`
	AreaId       uint64 `json:"area_id" form:"area_id"`
	TopicId      uint64 `json:"topic_id" form:"topic_id"`
	TopicItemId  uint64 `json:"topic_item_id" form:"topic_item_id"`
	Score        uint64 `json:"score" form:"score"`
	Status       uint64 `json:"status" form:"status"`
	Note         string `json:"note" form:"note"`
	Created_at   uint64 `json:"created_at" form:"created_at"`
	Created_by   uint64 `json:"created_by" form:"created_by"`
}

// type InspectionUpdateDTO struct {
// 	ID         uint64 `json:"id"`
// 	Scored     uint64 `json:"scored" form:"scored" binding:"required"`
// 	Note       string `json:"note"`
// 	Created_at uint64 `json:"created_at"`
// 	Created_by uint64 `json:"created_by" binding:"required"`
// }
