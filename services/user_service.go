package services

import (
	"errors"
	"fmt"
	"library/middleware"
	"library/models"
	"library/repository"
	"library/utils"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {
	existingUser, err := s.Repo.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPass

	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(req *models.User) (string, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
