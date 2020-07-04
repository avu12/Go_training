package user

import (
	"errors"
	"strings"

	"../database"
	"../entity"
)

type UserService struct {
	userRepository     database.User
	badWordsRepository database.BadWords
}

func NewUserService(userRepository database.User, badWordsRepository database.BadWords) *UserService {
	return &UserService{
		userRepository:     userRepository,
		badWordsRepository: badWordsRepository,
	}
}
func (c *UserService) Register(user entity.User) error {

	BadWords, err := c.badWordsRepository.FindAll()
	if err != nil {
		return err
	}
	for _, BadWord := range BadWords {
		if strings.Contains(user.Description, BadWord) {
			return errors.New("Bad word found!")
		}
	}

	err = c.userRepository.Add(user)
	if err != nil {
		return err
	}

	return nil

}
