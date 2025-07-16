package services

import (
	"library/models"
	"library/repository"
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

	token, err := middleware.GenerateJWT(user.ID.String(), user.userRole)
	if err != nil {
		return "", err
	}
	return token, nil
}
