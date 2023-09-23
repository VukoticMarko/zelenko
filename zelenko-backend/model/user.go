package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Mail     string
	Password string
	Name     string
	Surname  string
	//Picture
	City    string
	Country string
	Sex     string
	//Birthday date
	Disabled bool
	UserRank UserRank
	Role     string
}

const (
	// Sex
	Female = "female"
	Male   = "male"

	// Role
	Normal = "normal"
	Admin  = "admin"
)
