package model

import "time"

type Organisation struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	ContactNumber string    `json:"contact_number"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
