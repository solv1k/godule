package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primarykey"`
}

type BaseIntModel struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type BaseModelWithTimestamps struct {
	BaseModel
	CreatedAt time.Time `json:"created_at" faker:"-"`
	UpdatedAt time.Time `json:"updated_at" faker:"-"`
}

type BaseIntModelWithTimestamps struct {
	BaseIntModel
	CreatedAt time.Time `json:"created_at" faker:"-"`
	UpdatedAt time.Time `json:"updated_at" faker:"-"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}
