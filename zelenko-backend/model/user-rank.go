package model

import "github.com/google/uuid"

type UserRank struct {
	ID         uuid.UUID
	UserPoints int
	UserRank   string
	// Premium
	// Premium Length

}

const (
	// User Rank
	Baby = "baby"
)
