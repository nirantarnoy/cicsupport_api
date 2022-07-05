package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type InspectionRepository interface {
	FindById(id uint64) entity.InspectionRecord
	CreateInspection(inspection entity.InspectionRecord) entity.InspectionRecord
	UpdateInspection(inspection entity.InspectionRecord) entity.InspectionRecord
}

type inspectionConnect struct {
	connect *gorm.DB
}

// CreateInspection implements InspectionRepository
func (ins *inspectionConnect) CreateInspection(inspection entity.InspectionRecord) entity.InspectionRecord {
	ins.connect.Table("").Save(&inspection)
	return inspection
}

// FindById implements InspectionRepository
func (*inspectionConnect) FindById(id uint64) entity.InspectionRecord {
	panic("unimplemented")
}

// UpdateInspection implements InspectionRepository
func (ins *inspectionConnect) UpdateInspection(inspection entity.InspectionRecord) entity.InspectionRecord {
	ins.connect.Table("").Updates(&inspection)
	return inspection
}

func NewInspectionRepository(db *gorm.DB) InspectionRepository {
	return &inspectionConnect{connect: db}
}
