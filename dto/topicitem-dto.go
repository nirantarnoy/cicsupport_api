package dto

type TopicitemDTO struct {
	TopicID        uint64 `json:"topic_id" form:"topic_id" binding:"required"`
	TopicName      string `json:"topic_name"`
	TopicItemID    uint64 `json:"topic_item_id"`
	TopiceItemName string `json:"topic_item_name"`
	Status         uint64 `json:"status"`
}
