package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type AreagroupService interface {
	ListAll() []entity.Areagroup
	FindById(id uint64) entity.Areagroup
	Create(areagroup dto.AreagroupDTO) entity.Areagroup
	Update(areagroup dto.AreagroupUpdateDTO) entity.Areagroup
	Delete(areagroup entity.Areagroup)
	CheckAllowed(userId string, areagroupId uint64) bool
}

type areagroupService struct {
	areagroupRepo repository.AreagroupRepository
}

// Delete implements AreagroupService
func (ag *areagroupService) Delete(areagroup entity.Areagroup) {
	ag.areagroupRepo.DeleteAreagroup(areagroup)
}

// Update implements AreagroupService
func (ag *areagroupService) Update(areagroupDto dto.AreagroupUpdateDTO) entity.Areagroup {
	aggroup := entity.Areagroup{}
	err := smapping.FillStruct(&aggroup, smapping.MapFields(&areagroupDto))
	if err != nil {
		log.Fatalf("Fail to map %v", err)
	}
	log.Print(aggroup)
	res := ag.areagroupRepo.UpdateAreagroup(aggroup)
	return res
}

// Create implements AreagroupService
func (ag *areagroupService) Create(areagroupDTO dto.AreagroupDTO) entity.Areagroup {
	aggroup := entity.Areagroup{}
	err := smapping.FillStruct(&aggroup, smapping.MapFields(&areagroupDTO))
	if err != nil {
		log.Fatalf("Fail to mapping %v", err)
	}
	res := ag.areagroupRepo.InsertAreagroup(aggroup)
	return res
}

// FindById implements AreagroupService
func (ag_service *areagroupService) FindById(id uint64) entity.Areagroup {
	return ag_service.areagroupRepo.FindAreaById(id)
}

func NewAreaGroupService(areagroupRepo repository.AreagroupRepository) AreagroupService {
	return &areagroupService{
		areagroupRepo: areagroupRepo,
	}
}

// FindListAll implements AreagroupService
func (aregroup *areagroupService) ListAll() []entity.Areagroup {
	return aregroup.areagroupRepo.FindListAll()
}

// isAllowedToEdit implements AreagroupService
func (ag *areagroupService) CheckAllowed(userId string, areagroupId uint64) bool {
	res := ag.areagroupRepo.FindAreaById(areagroupId)
	id := fmt.Sprintf("%v", res.Created_by)
	return userId == id
}
