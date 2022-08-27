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
	var plans entity.Plan

	db.connect.Table("inspection_plan").Where("team_id = ?", team_id).Scan(&plans)
	if (plans != entity.Plan{}) {
		db.connect.Table("query_inspection_item_with_plan").Where("area_group_id = ?", plans.Inspection_area_id).Where("is_enable = 1").Find(&inspectionitems)
		return inspectionitems
	} else {
		return nil
	}

}

func NewTeaminspectionitemRepository(db *gorm.DB) TeaminspectionitemRepository {
	return &teaminspectionitemConnect{connect: db}
}
