package services

import (
    "errors"
    "github.com/google/uuid"
    "seems/domain/dto"
    "seems/domain/models"
)

type Users struct {
    Storage map[uuid.UUID]*models.User
}

func (u *Users) GetUserById(id uuid.UUID) (*models.User, error) {
    return u.Storage[id], nil
}

func (u *Users) GetUserByCredentials(creds *dto.Credentials) (*models.User, error) {
    if creds == nil {
        return nil, errors.New("credentials cannot be nil")
    }
    for _, user := range u.Storage {
        if user.Name == creds.Login && user.Password == creds.Password {
            return user, nil
        }
    }
    return nil, errors.New("user not found")
}
func (u *Users) AddUser(user *models.User) error {
    if user != nil {
        u.Storage[user.Id] = user
        return nil
    }
    return errors.New("user is nil")
}
