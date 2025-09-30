package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/supermario64bit/whatsapp_connect/server/model"
	"github.com/supermario64bit/whatsapp_connect/server/repository"
	"github.com/supermario64bit/whatsapp_connect/types"
)

type userservice struct {
	repo repository.UserRepository
}

type UserService interface {
	Create(user *model.User) (*model.User, *types.ApplicationError)
	Find(filter *model.User) ([]*model.User, *types.ApplicationError)
	FindByID(id uint64) (*model.User, *types.ApplicationError)
	UpdateByID(updates *model.User, id uint64) (*model.User, *types.ApplicationError)
	DeleteByID(id uint64) *types.ApplicationError
}

func NewUserService() UserService {
	return &userservice{
		repo: repository.NewUserRepository(),
	}
}

func (svc *userservice) Create(user *model.User) (*model.User, *types.ApplicationError) {
	validationErrors := user.ValidateFields()
	if len(validationErrors) > 0 {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusBadRequest,
			Message:    "Validation Failed",
			Err:        fmt.Errorf("", validationErrors),
		}
	}

	new, err := svc.repo.Create(user)
	if err != nil {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to create user",
			Err:        fmt.Errorf("Unable to create user. Error: " + err.Error()),
		}
	}

	return new, nil
}

func (svc *userservice) Find(filter *model.User) ([]*model.User, *types.ApplicationError) {
	userSet, err := svc.repo.Find(filter)
	if err != nil {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to find users",
			Err:        err,
		}
	}
	return userSet, nil
}

func (svc *userservice) FindByID(id uint64) (*model.User, *types.ApplicationError) {
	user, err := svc.repo.FindByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to find user by id",
			Err:        err,
		}
	}
	return user, nil
}

func (svc *userservice) UpdateByID(updates *model.User, id uint64) (*model.User, *types.ApplicationError) {
	updatedUser, err := svc.repo.UpdateByID(updates, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &types.ApplicationError{
				HttpStatus: http.StatusNoContent,
				Message:    "Unable to update user",
				Err:        fmt.Errorf("No user available for the given id"),
			}
		}
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to update user",
			Err:        err,
		}
	}
	return updatedUser, nil
}

func (svc *userservice) DeleteByID(id uint64) *types.ApplicationError {
	err := svc.repo.DeleteByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.ApplicationError{
				HttpStatus: http.StatusNoContent,
				Message:    "Unable to delete user",
				Err:        fmt.Errorf("No user available for the given id"),
			}
		}
		return &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to delete user",
			Err:        err,
		}
	}
	return nil
}
