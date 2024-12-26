package interfaces

import (
	"github.com/google/uuid"
	"seems/domain/dto"
	"seems/domain/models"
)

type UserService interface {
	GetUserById(id uuid.UUID) (*models.User, error)
	GetUserByCredentials(creds *dto.Credentials) (*models.User, error)
	AddUser(user models.User) error
}
