package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUserAD(userAD string) entity.User
	ProfileUser(userID string) entity.User
}

type UserConnect struct {
	connect *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnect{connect: db}
}

func (db *UserConnect) FindByUserAD(userAD string) entity.User {
	var user entity.User
	db.connect.Table("person").Where("ad_user = ?", userAD).Take(&user)
	return user
}

func (db *UserConnect) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connect.Table("person").Find(&user, userID)
	return user
}
