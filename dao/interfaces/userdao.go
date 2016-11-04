package interfaces

import "github.com/alexyslozada/daopattern/models"

type UserDAO interface {
	Create(u *models.User) error
	GetAll() ([]models.User, error)
}
