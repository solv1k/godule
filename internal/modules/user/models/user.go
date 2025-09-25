package models

import (
	"time"

	mm "github.com/solv1k/croco-api/internal/modules/media/models"
	"github.com/solv1k/croco-api/internal/shared/models"
	"gorm.io/gorm"
)

// User model
type User struct {
	models.BaseModelWithTimestamps
	// Fields
	Name     string     `json:"name" validate:"required,min=2,max=255" gorm:"size:255" faker:"first_name"`
	Email    string     `json:"email" validate:"required,email,max=255" gorm:"uniqueIndex;size:255" faker:"email,unique"`
	OnlineAt *time.Time `json:"online_at" gorm:"type:timestamp" faker:"-"`
	// Relations
	Avatar *mm.Media `json:"avatar" gorm:"polymorphic:Model"`
}

// Hook to set avatar collection before create user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Avatar != nil {
		u.Avatar.Collection = "avatar"
	}
	return nil
}

// ID for AuthEntity interface [pkg/auth]
func (u *User) GetAuthID() string {
	return u.ID.String()
}

// Code type for AuthEntity interface [pkg/auth]
func (u *User) GetAuthCodeType() string {
	return "email"
}

// Code identifier for AuthEntity interface [pkg/auth]
func (u *User) GetAuthCodeIdentifier() string {
	return u.Email
}

// Payload for AuthEntity interface [pkg/auth]
func (u *User) GetAuthPayload() map[string]interface{} {
	return map[string]interface{}{}
}
