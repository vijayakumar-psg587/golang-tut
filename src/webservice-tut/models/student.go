package models

import (
	"github.com/google/uuid"
)

type Student struct {
	studentId   uuid.UUID
	studentAddr Address
}
