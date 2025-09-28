package model

import "time"

type Organisation struct {
	ID            uint64    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	ContactNumber string    `json:"contact_number" db:"contact_number"`
	Email         string    `json:"email" db:"email"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`
}
