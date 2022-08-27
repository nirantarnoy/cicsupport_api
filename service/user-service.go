package service

import (
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type UserService interface {
	Profile(userID string, userName string) entity.User
	FindUserTeam(team_id uint64) []entity.TeamMember
}

type userService struct {
	userRepo repository.UserRepository
}

// FindUserTeam implements UserService
func (db *userService) FindUserTeam(team_id uint64) []entity.TeamMember {
	return db.userRepo.FindUserTeam(team_id)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Profile(userID string, username string) entity.User {
	return u.userRepo.ProfileUser(userID, username)
}
