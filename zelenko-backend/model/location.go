package model

import "github.com/google/uuid"

type Location struct {
	ID        uuid.UUID
	Latitude  float32
	Longitude float32
	Street    string
	City      string
	Country   string
}
