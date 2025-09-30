package model

import "time"

type User struct {
	ID        uint64     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name" validate:"required,min=2,max=100"`
	Handle    string     `json:"handle" db:"handle" validate:"required,min=2,max=100"`
	Mobile    string     `json:"mobile_number" db:"mobile_number" validate:"required,len=10,numeric"`
	Email     string     `json:"email" db:"email" validate:"required,email"`
	Status    string     `json:"status" db:"status" validate:"required,oneof=active inactive"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
