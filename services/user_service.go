package services

import (
	"errors"
	"fmt"
	"library/middleware"
	"library/models"
	"library/repository"

	"library/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}


func (s *UserService) Login(req *models.User) (string, error) {
		user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID.String(), user.UserRole)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (s *UserService) RegisterUser(req *models.User) error {
existingUser, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil && existingUser != nil {
    return fmt.Errorf("user already exists")
}

if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
    return err 
} 

	//hashing the password
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPass

	myuuid := uuid.New()
	req.ID = myuuid

	//put the users into the database
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil

}
