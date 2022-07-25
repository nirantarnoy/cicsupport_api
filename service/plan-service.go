package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type PlanService interface {
	FindPlanByTeam(team_id uint64) entity.Plan
	AddInspection(inspectionData []dto.InspectionCreateDTO) entity.InspectionRecord
}
type planService struct {
	planRepo repository.PlanRepository
}

// AddInspection implements PlanService
func (db *planService) AddInspection(inspectionDto []dto.InspectionCreateDTO) entity.InspectionRecord {
	inspection := entity.InspectionRecord{}
	for i := 0; i < len(inspectionDto); i++ {
		err := smapping.FillStruct(&inspection, smapping.MapFields(&inspectionDto[i]))
		if err != nil {
			log.Fatalf("fail to map %v", err)
		}
		db.planRepo.AddInspection(inspection)
	}
	return entity.InspectionRecord{}
}

// FindPlanByTeam implements PlanService
func (db *planService) FindPlanByTeam(team_id uint64) entity.Plan {
	return db.planRepo.FindPlanByTeam(team_id)
}

func NewPlanService(repo repository.PlanRepository) PlanService {
	return &planService{planRepo: repo}
}
