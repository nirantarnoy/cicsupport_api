package dto

type TeamInspectionItemDto struct {
	TeamId               uint64 `json:"id"`
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
	IsEnable             uint64 `json:"is_enable"`
	SeqSort              uint64 `json:"seq_sort"`
	SeqSortItem          uint64 `json:"seq_sort_item"`
}
