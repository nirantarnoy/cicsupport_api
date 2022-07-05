package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type AreazoneRepository interface {
	Findlistall() []entity.Areazone
	FindById(id uint64) entity.Areazone
	CreateAreazone(areazone entity.Areazone) entity.Areazone
	UpdateAreazone(areazone entity.Areazone) entity.Areazone
	DeleteAreazone(areazone entity.Areazone) entity.Areazone
}

type AreazoneConnnect struct {
	connect *gorm.DB
}

// CreateAreazone implements AreazoneRepository
func (az *AreazoneConnnect) CreateAreazone(areazone entity.Areazone) entity.Areazone {
	az.connect.Table("area_zone").Save(areazone)
	return areazone
}

// DeleteAreazone implements AreazoneRepository
func (az *AreazoneConnnect) DeleteAreazone(areazone entity.Areazone) entity.Areazone {
	az.connect.Table("area_zone").Delete(areazone)
	return areazone
}

// FindById implements AreazoneRepository
func (az *AreazoneConnnect) FindById(id uint64) entity.Areazone {
	var areazone entity.Areazone
	az.connect.Table("area_zone").Find(areazone, id)
	return areazone
}

// Findlistall implements AreazoneRepository
func (az *AreazoneConnnect) Findlistall() []entity.Areazone {
	var areazone []entity.Areazone
	az.connect.Table("area_zone").Find(areazone)
	return areazone

}

// UpdateAreazone implements AreazoneRepository
func (*AreazoneConnnect) UpdateAreazone(areazone entity.Areazone) entity.Areazone {
	panic("unimplemented")
}

func NewAreazoneRepository(db *gorm.DB) AreazoneRepository {
	return &AreazoneConnnect{connect: db}
}
