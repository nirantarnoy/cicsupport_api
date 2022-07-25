package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type TeaminspectionitemRepository interface {
	FindInspectionItem(team_id uint64) []entity.TeamInspectionItem
}
type teaminspectionitemConnect struct {
	connect *gorm.DB
}

// FindInspectionItem implements TeaminspectionitemRepository
func (db *teaminspectionitemConnect) FindInspectionItem(team_id uint64) []entity.TeamInspectionItem {
	var inspectionitems []entity.TeamInspectionItem
	db.connect.Table("query_inspection_item_with_plan").Find(&inspectionitems, team_id)
	return inspectionitems
}

func NewTeaminspectionitemRepository(db *gorm.DB) TeaminspectionitemRepository {
	return &teaminspectionitemConnect{connect: db}
}
