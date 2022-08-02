package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type CarinspecRepository interface {
	NonconformList() []entity.NonConformTitle
}

type carinspecConnect struct {
	connect *gorm.DB
}

// NonconformList implements CarinspecRepository
func (db *carinspecConnect) NonconformList() []entity.NonConformTitle {
	var nonconform []entity.NonConformTitle
	db.connect.Table("nonconform").Find(&nonconform)
	return nonconform
}

func NewCarinspecRepository(db *gorm.DB) CarinspecRepository {
	return &carinspecConnect{connect: db}
}
