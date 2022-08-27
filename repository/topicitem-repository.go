package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type TopiceitemRepository interface {
	FindTopicByPlan(planId uint64) []entity.TopicItem
}

type topiceitemConnect struct {
	connect *gorm.DB
}

// FindTopicByPlan implements TopiceitemRepository
func (db *topiceitemConnect) FindTopicByPlan(planId uint64) []entity.TopicItem {
	var topicitems []entity.TopicItem
	db.connect.Table("qry_topic_item").Find(&topicitems, planId).Order("seq_sort")
	return topicitems
}

func NewTopicitemRepository(db *gorm.DB) TopiceitemRepository {
	return &topiceitemConnect{connect: db}
}
