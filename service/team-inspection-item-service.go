package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type TeaminspectionitemService interface {
	FindInspectionItem(team_id uint64) []entity.TeamInspectionItem
	FindTransByEmp(team_id int64, emp_id int64) int64
	FindTransHistoryByEmp(emp_id int64) []entity.InspectionTransHistory
}
type teaminspectionitemService struct {
	teaminspectionitemRepo repository.TeaminspectionitemRepository
}

// FindInspectionItem implements TeaminspectionitemService
func (db *teaminspectionitemService) FindInspectionItem(team_id uint64) []entity.TeamInspectionItem {
	return db.teaminspectionitemRepo.FindInspectionItem(team_id)
}

func (db *teaminspectionitemService) FindTransByEmp(team_id int64, emp_id int64) int64 {
	return db.teaminspectionitemRepo.FindTransByEmp(team_id, emp_id)
}

func (db *teaminspectionitemService) FindTransHistoryByEmp(emp_id int64) []entity.InspectionTransHistory {
	return db.teaminspectionitemRepo.FindTransHistoryByEmp(emp_id)
}

func NewTeaminspectionitemService(repo repository.TeaminspectionitemRepository) TeaminspectionitemService {
	return &teaminspectionitemService{teaminspectionitemRepo: repo}
}
