package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type TopicItemService interface {
	FindTopicByPlan(planId uint64) []entity.TopicItem
}

type topiceItemService struct {
	topiceItemRepo repository.TopiceitemRepository
}

// FindTopicByPlan implements TopicItemService
func (db *topiceItemService) FindTopicByPlan(planId uint64) []entity.TopicItem {
	return db.topiceItemRepo.FindTopicByPlan(planId)
}

func NewTopicItemService(topicItemRepo repository.TopiceitemRepository) TopicItemService {
	return &topiceItemService{topiceItemRepo: topicItemRepo}
}
