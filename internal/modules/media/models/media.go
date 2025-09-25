package models

import (
	"github.com/google/uuid"
	shared "github.com/solv1k/croco-api/internal/shared/models"
)

// Media model
type Media struct {
	shared.BaseIntModel
	// Fields
	Url        string    `json:"url" gorm:"size:255" faker:"url"`
	Collection string    `json:"collection" gorm:"size:255" faker:"-"`
	ModelID    uuid.UUID `json:"model_id" gorm:"index" faker:"-"`
	ModelType  string    `json:"model_type" gorm:"index" faker:"-"`
}
