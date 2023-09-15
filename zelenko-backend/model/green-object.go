package model

import "github.com/google/uuid"

type GreenObject struct {
	ID           uuid.UUID
	LocationName string
	// Location Location
	Shape      string
	TrashType  string
	GreenScore GreenScore
	Disabled   bool
}
