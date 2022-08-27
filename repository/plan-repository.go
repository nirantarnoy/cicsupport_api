package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type PlanRepository interface {
	FindPlan(module_type_id uint64) []entity.Plan
	UpdatePlan(plan entity.Plan) entity.Plan
	FindPlanByTeam(team_id uint64) entity.Plan

	AddInspection(inspection entity.InspectionRecord) entity.InspectionRecord
}

type planConnect struct {
	connect *gorm.DB
}

// AddInspection implements PlanRepository
func (db *planConnect) AddInspection(inspection entity.InspectionRecord) entity.InspectionRecord {
	db.connect.Table("inspection_trans").Create(&inspection)
	return inspection
}

// FindPlanByTeam implements PlanRepository
func (db *planConnect) FindPlanByTeam(team_id uint64) entity.Plan {
	var plans entity.Plan
	db.connect.Table("person").Find(&plans, team_id)
	return plans
}

// UpdatePlan implements PlanRepository
func (pl *planConnect) UpdatePlan(plan entity.Plan) entity.Plan {
	pl.connect.Table("inspection_plan").Updates(&plan)
	return plan
}

// FindPlan implements PlanRepository
func (pl *planConnect) FindPlan(module_type_id uint64) []entity.Plan {
	panic("unimplemented")
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planConnect{connect: db}
}
