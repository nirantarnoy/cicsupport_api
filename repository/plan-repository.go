package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type PlanRepository interface {
	FindPlan(module_type_id uint64) []entity.Plan
	UpdatePlan(plan entity.Plan) entity.Plan
}

type planConnect struct {
	connect *gorm.DB
}

// UpdatePlan implements PlanRepository
func (pl *planConnect) UpdatePlan(plan entity.Plan) entity.Plan {
	pl.connect.Table("inspection_plan").Updates(&plan)
	return plan
}

// FindPlan implements PlanRepository
func (*planConnect) FindPlan(module_type_id uint64) []entity.Plan {
	panic("unimplemented")
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planConnect{connect: db}
}
