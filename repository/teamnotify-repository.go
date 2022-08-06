package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type TeamNotifyRepository interface {
	FindEmpNotify(emp_id uint64) []entity.TeamNotify
}

type teamNotifyConnect struct {
	connect *gorm.DB
}

// EmpNotify implements TeamNotifyRepository
func (db *teamNotifyConnect) FindEmpNotify(emp_id uint64) []entity.TeamNotify {
	var teamnotidata []entity.TeamNotify
	db.connect.Table("team_notify").Where("emp_id = ?", emp_id).Scan(&teamnotidata)
	return teamnotidata
}

func NewTeamnotifyRepository(db *gorm.DB) TeamNotifyRepository {
	return &teamNotifyConnect{connect: db}
}
