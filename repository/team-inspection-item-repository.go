package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type TeaminspectionitemRepository interface {
	FindInspectionItem(team_id uint64) []entity.TeamInspectionItem
	FindTransByEmp(team_id int64, emp_id int64) int64
	FindTransHistoryByEmp(emp_id int64) []entity.InspectionTransHistory
}
type teaminspectionitemConnect struct {
	connect *gorm.DB
}

// FindInspectionItem implements TeaminspectionitemRepository
func (db *teaminspectionitemConnect) FindInspectionItem(team_id uint64) []entity.TeamInspectionItem {
	var inspectionitems []entity.TeamInspectionItem
	var plans entity.Plan

	db.connect.Table("inspection_plan").Where("team_id = ?", team_id).Where("status = 0").Scan(&plans)
	if (plans != entity.Plan{}) {
		db.connect.Table("query_inspection_item_with_plan").Where("area_group_id = ?", plans.Inspection_area_id).Where("is_enable = 1").Where("status=0").Find(&inspectionitems)
		return inspectionitems
	} else {
		return nil
	}

}
func (db *teaminspectionitemConnect) FindTransByEmp(team_id int64, emp_id int64) int64 {
	var counts int64

	db.connect.Table("inspection_trans").Where("team_id = ?", team_id).Where("emp_id = ?", emp_id).Count(&counts)

	return counts

}
func (db *teaminspectionitemConnect) FindTransHistoryByEmp(emp_id int64) []entity.InspectionTransHistory {
	var items []entity.InspectionTransHistory
	db.connect.Table("inspection_trans").Select("distinct(inspection_trans.plan_id),inspection_plan.plan_target_date as plan_date,inspection_plan.plan_no,inspection_trans.status,date(inspection_trans.trans_date) as plan_actual_date").Joins("inner join inspection_plan on inspection_trans.plan_id = inspection_plan.plan_num").Where("inspection_trans.emp_id = ?", emp_id).Where("inspection_trans.status = 1").Scan(&items)
	return items
}

func NewTeaminspectionitemRepository(db *gorm.DB) TeaminspectionitemRepository {
	return &teaminspectionitemConnect{connect: db}
}
