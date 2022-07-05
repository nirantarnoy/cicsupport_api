package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type AreazoneService interface {
	FindListAll() []entity.Areazone
	FindById(id uint64) entity.Areazone
	CreateAreazone(areazone dto.AreazoneCreateDto) entity.Areazone
	UpdateAreazone(areazone dto.AreazoneUpdateDto) entity.Areazone
	DeleteAreazone(areazone entity.Areazone)
	CheckAllowed(userId string, areazoneId uint64) bool
}
type areazoneService struct {
	areagzoneRepo repository.AreazoneRepository
}

// FindById implements AreazoneService
func (az *areazoneService) FindById(id uint64) entity.Areazone {
	return az.areagzoneRepo.FindById(id)
}

// CheckAllowed implements AreazoneService
func (az *areazoneService) CheckAllowed(userId string, areazoneId uint64) bool {
	res := az.areagzoneRepo.FindById(areazoneId)
	id := fmt.Sprintf("%v", res.Created_by)
	return userId == id
	//return true
}

// CreateAreazone implements AreazoneService
func (az *areazoneService) CreateAreazone(areazoneDto dto.AreazoneCreateDto) entity.Areazone {
	areazone := entity.Areazone{}
	err := smapping.FillStruct(&areazone, smapping.MapFields(&areazoneDto))
	if err != nil {
		log.Fatalf("Fail to map %v", err)
	}

	return az.areagzoneRepo.CreateAreazone(areazone)
}

// DeleteAreazone implements AreazoneService
func (az *areazoneService) DeleteAreazone(areazone entity.Areazone) {
	az.areagzoneRepo.DeleteAreazone(areazone)
}

// FindListAll implements AreazoneService
func (az *areazoneService) FindListAll() []entity.Areazone {
	return az.areagzoneRepo.Findlistall()
}

// UpdateAreazone implements AreazoneService
func (az *areazoneService) UpdateAreazone(areazoneDto dto.AreazoneUpdateDto) entity.Areazone {
	areazone := entity.Areazone{}
	err := smapping.FillStruct(&areazone, smapping.MapFields(&areazoneDto))
	if err != nil {
		log.Fatalf("Fail to map %v", err)
	}
	log.Print(areazone)
	res := az.areagzoneRepo.UpdateAreazone(areazone)
	return res
}

func NewAreazoneService(repo repository.AreazoneRepository) AreazoneService {
	return &areazoneService{areagzoneRepo: repo}
}
