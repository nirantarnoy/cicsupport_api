package entity

type TopicItem struct {
	TopicID       uint64 `json:"topic_id"`
	TopicName     string `json:"topic_name"`
	TopicItemID   uint64 `json:"topic_item_id"`
	TopicItemName string `json:"topic_item_name"`
	Status        uint64 `json:"status"`
}
