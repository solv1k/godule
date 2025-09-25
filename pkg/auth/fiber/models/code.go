package models

import (
	"time"
)

// Auth code model
type AuthCode struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Type       string    `json:"type" gorm:"uniqueIndex:auth_code_index"`
	Identifier string    `json:"identifier" gorm:"uniqueIndex:auth_code_index"`
	Code       string    `json:"code"`
	ExpiresAt  time.Time `json:"expires_at" gorm:"type:timestamp"`
}
