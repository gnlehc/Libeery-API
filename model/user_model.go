package model

import (
	"time"

	"github.com/google/uuid"
)

type MsUser struct {
	UserID    uuid.UUID `gorm:"primaryKey;unique;default:gen_random_uuid()"`
	NIM       *string   `gorm:"size:10" json:"NIM,omitempty"`
	NIS       *string   `gorm:"size:5" json:"NIS,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt,omitempty"`
	Stsrc     string    `gorm:"size:1" json:"Stsrc,omitempty"`
}
