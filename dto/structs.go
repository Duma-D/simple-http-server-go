package dto

import (
	"github.com/google/uuid"
)

type PersonDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Sex       string    `json:"sex"`
}

type Person struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Sex       byte    `json:"sex"`
}
