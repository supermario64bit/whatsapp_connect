package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/supermario64bit/whatsapp_connect/server/model"
	"github.com/supermario64bit/whatsapp_connect/server/repository"
	"github.com/supermario64bit/whatsapp_connect/types"
)

type organisationService struct {
	repo repository.OrganisationRepository
}

type OrganisationService interface {
	Create(org *model.Organisation) (*model.Organisation, *types.ApplicationError)
	Find(filter *model.Organisation) ([]*model.Organisation, *types.ApplicationError)
	FindByID(id uint64) (*model.Organisation, *types.ApplicationError)
	UpdateByID(updates *model.Organisation, id uint64) (*model.Organisation, *types.ApplicationError)
	DeleteByID(id uint64) *types.ApplicationError
}

func NewOrganisationService() OrganisationService {
	return &organisationService{
		repo: repository.NewOrganisationRepository(),
	}
}

func (svc *organisationService) Create(org *model.Organisation) (*model.Organisation, *types.ApplicationError) {
	validationErrors := org.ValidateFields()
	if len(validationErrors) > 0 {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusBadRequest,
			Message:    "Validation Failed",
			Err:        fmt.Errorf("", validationErrors),
		}
	}

	new, err := svc.repo.Create(org)
	if err != nil {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to create organisation",
			Err:        fmt.Errorf("Unable to create organisation. Error: " + err.Error()),
		}
	}

	return new, nil
}

func (svc *organisationService) Find(filter *model.Organisation) ([]*model.Organisation, *types.ApplicationError) {
	orgSet, err := svc.repo.Find(filter)
	if err != nil {
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to find organisations",
			Err:        err,
		}
	}
	return orgSet, nil
}

func (svc *organisationService) FindByID(id uint64) (*model.Organisation, *types.ApplicationError) {
	org, err := svc.repo.FindByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to find organisation by id",
			Err:        err,
		}
	}
	return org, nil
}

func (svc *organisationService) UpdateByID(updates *model.Organisation, id uint64) (*model.Organisation, *types.ApplicationError) {
	updatedOrg, err := svc.repo.UpdateByID(updates, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &types.ApplicationError{
				HttpStatus: http.StatusNoContent,
				Message:    "Unable to update organisation",
				Err:        fmt.Errorf("No organistion available for the given id"),
			}
		}
		return nil, &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to update organisation",
			Err:        err,
		}
	}
	return updatedOrg, nil
}

func (svc *organisationService) DeleteByID(id uint64) *types.ApplicationError {
	err := svc.repo.DeleteByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.ApplicationError{
				HttpStatus: http.StatusNoContent,
				Message:    "Unable to delete organisation",
				Err:        fmt.Errorf("No organistion available for the given id"),
			}
		}
		return &types.ApplicationError{
			HttpStatus: http.StatusInternalServerError,
			Message:    "Unable to delete organisation",
			Err:        err,
		}
	}
	return nil
}
