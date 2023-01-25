package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type TeamNotifyService interface {
	FindEmpNotify(emp_id uint64) []entity.TeamNotify
	FindTeamNotify(emp_id uint64) []entity.TeamNotify
}

type teamNotify struct {
	teamNotifyRepo repository.TeamNotifyRepository
}

// FindEmpNotify implements TeamNotifyService
func (db *teamNotify) FindEmpNotify(emp_id uint64) []entity.TeamNotify {
	return db.teamNotifyRepo.FindEmpNotify(emp_id)
}
func (db *teamNotify) FindTeamNotify(emp_id uint64) []entity.TeamNotify {
	return db.teamNotifyRepo.FindTeamNotify(emp_id)
}

func NewTeamnotifyService(repo repository.TeamNotifyRepository) TeamNotifyService {
	return &teamNotify{teamNotifyRepo: repo}
}
