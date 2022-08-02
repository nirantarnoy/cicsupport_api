package repository

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUserAD(userAD string) entity.User
	ProfileUser(userID string) entity.User
	FindUserTeam(team_id uint64) []entity.TeamMember
}

type UserConnect struct {
	connect *gorm.DB
}

// FindUserTeam implements UserRepository
func (db *UserConnect) FindUserTeam(team_id uint64) []entity.TeamMember {
	var member []entity.TeamMember
	db.connect.Table("person").Select("current_team_id,employee.fname,employee.lname").Joins("inner join employee on employee.id = person.emp_id").Where("person.current_team_id = ?", team_id).Scan(&member)
	return member

	// rows, err := db.connect.Table().Rows();
	// for rows.Next() {

	// }
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnect{connect: db}
}

func (db *UserConnect) FindByUserAD(userAD string) entity.User {
	var user entity.User
	db.connect.Table("user").Where("dns_user = ?", userAD).Take(&user)
	return user
}

func (db *UserConnect) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connect.Table("person").Find(&user, userID)
	return user
}
