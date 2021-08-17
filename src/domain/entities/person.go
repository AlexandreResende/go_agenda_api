package entities

import uuid "github.com/satori/go.uuid"

type Person struct {
	Id uuid.UUID `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}
