package services

import (
	"../domain"
	"../utils"
)

// here it could be the business logic!
func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
