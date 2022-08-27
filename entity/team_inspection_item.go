package entity

type TeamInspectionItem struct {
	TeamId               uint64 `json:"id"`
	Plan_no              string `json:"plan_no"`
	Plan_target_date     string `json:"plan_target_date"`
	Module_type_id       uint64 `json:"module_type_id"`
	AreaInspectionId     uint64 `json:"area_inspection_id"`
	AreaInspectionName   string `json:"area_inspection_name"`
	AreaId               uint64 `json:"area_id"`
	AreaName             string `json:"area_name"`
	AreagroupId          uint64 `json:"area_group_id"`
	AreagroupName        string `json:"area_group_name"`
	AreaInspectionLineId uint64 `json:"area_inspection_line_id"`
	TopicId              uint64 `json:"topic_id"`
	TopicName            string `json:"topic_name"`
	TopicItemId          uint64 `json:"topic_item_id"`
	TopicItemName        string `json:"topic_item_name"`
	Is_enable            uint64 `json:"is_enable"`
}
