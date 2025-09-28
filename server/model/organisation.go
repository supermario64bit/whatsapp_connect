package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Organisation struct {
	ID            uint64    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name" validate:"required,min=2,max=100"`
	ContactNumber string    `json:"contact_number" db:"contact_number" validate:"required,len=10,numeric"`
	Email         string    `json:"email" db:"email" validate:"required,email"`
	Status        string    `json:"status" db:"status" validate:"required,oneof=active inactive"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`
}

func (org Organisation) ValidateFields() []error {
	validate := validator.New()
	err := validate.Struct(org)

	var errors []error
	if err == nil {
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err)
	}

	return errors
}
