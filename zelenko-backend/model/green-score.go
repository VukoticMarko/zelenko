package model

import "github.com/google/uuid"

type GreenScore struct {
	ID           uuid.UUID
	Verification int
	Report       int
	TrashRank    string
}

const (
	New = "new"
)
