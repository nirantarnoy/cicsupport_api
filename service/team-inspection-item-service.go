package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type TeaminspectionitemService interface {
	FindInspectionItem(team_id uint64) []entity.TeamInspectionItem
}
type teaminspectionitemService struct {
	teaminspectionitemRepo repository.TeaminspectionitemRepository
}

// FindInspectionItem implements TeaminspectionitemService
func (db *teaminspectionitemService) FindInspectionItem(team_id uint64) []entity.TeamInspectionItem {
	return db.teaminspectionitemRepo.FindInspectionItem(team_id)
}

func NewTeaminspectionitemService(repo repository.TeaminspectionitemRepository) TeaminspectionitemService {
	return &teaminspectionitemService{teaminspectionitemRepo: repo}
}
