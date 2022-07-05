package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type AreagroupRepository interface {
	FindListAll() []entity.Areagroup
	FindAreaById(id uint64) entity.Areagroup
	InsertAreagroup(areagroup entity.Areagroup) entity.Areagroup
	UpdateAreagroup(areagroup entity.Areagroup) entity.Areagroup
	DeleteAreagroup(areagroup entity.Areagroup)
}
type AreagroupConnect struct {
	connect *gorm.DB
}

// DeleteAreagroup implements AreagroupRepository
func (db *AreagroupConnect) DeleteAreagroup(areagroup entity.Areagroup) {
	db.connect.Table("area_group").Delete(&areagroup)
}

// InsertAreagroup implements AreagroupRepository
func (db *AreagroupConnect) InsertAreagroup(areagroup entity.Areagroup) entity.Areagroup {
	db.connect.Table("area_group").Save(&areagroup)
	//db.connect.Save(areagroup)
	//db.connect.Preload("area_group").Find(&areagroup)
	return areagroup
}

// UpdateAreagroup implements AreagroupRepository
func (db *AreagroupConnect) UpdateAreagroup(areagroup entity.Areagroup) entity.Areagroup {
	//	db.connect.Table("area_group").Where("id = ?", areagroup.ID).Updates(&areagroup)
	db.connect.Table("area_group").Updates(&areagroup)
	//db.connect.Preload("area_group").Find(&areagroup)
	return areagroup
}

// FindAreaById implements AreagroupRepository
func (db *AreagroupConnect) FindAreaById(id uint64) entity.Areagroup {
	var areagroup entity.Areagroup
	db.connect.Table("area_group").Find(&areagroup, id)
	return areagroup
}

// FindListAll implements AreagroupRepository
func (db *AreagroupConnect) FindListAll() []entity.Areagroup {
	var areagroup []entity.Areagroup
	db.connect.Table("area_group").Find(&areagroup)
	return areagroup
}

func NewAreaGroupRepository(db *gorm.DB) AreagroupRepository {
	return &AreagroupConnect{connect: db}
}
