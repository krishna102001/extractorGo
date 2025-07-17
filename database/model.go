package database

import (
	"github.com/google/uuid"
)

type ExtractsTable struct {
	ExtractId   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DocName     string    `gorm:"type:varchar(255);not null"`
	ResponseUrl string    `gorm:"not null"`
}
