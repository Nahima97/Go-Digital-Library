package services

import (
	"library/models"
	"library/repository"
	"library/utils"

	"github.com/google/uuid"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {

	//check if they exist
	_, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil {
		return err
	}

	//hashing the password
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPass

	myuuid := uuid.NewString()
	req.ID = myuuid
	

	//put into the database
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}
