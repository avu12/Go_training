package services

import (
	"../domain"
	"../utils"
)

type userService struct {
}

var (
	UsersService userService
)

// here it could be the business logic!
func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)

	if err != nil {
		return nil, err
	}
	return user, nil
}
