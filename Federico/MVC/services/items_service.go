package services

import (
	"net/http"

	"../domain"
	"../utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(userId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me!",
		StatusCode: http.StatusInternalServerError,
		Code:       "bad request",
	}
}
