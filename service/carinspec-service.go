package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type CarinspecSevice interface {
	NonconformList() []entity.NonConformTitle
}

type carinspecRepository struct {
	carinspecRepo repository.CarinspecRepository
}

// NonconformList implements CarinspecSevice
func (db *carinspecRepository) NonconformList() []entity.NonConformTitle {
	return db.carinspecRepo.NonconformList()
}

func NewCarinspecService(repo repository.CarinspecRepository) CarinspecSevice {
	return &carinspecRepository{carinspecRepo: repo}
}
