package services

import "gitlab.com/nisarg/7/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Updateuser(*models.User) error
	DeleteUser(*string) error
}
