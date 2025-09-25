package models

import (
	"github.com/google/uuid"
	mm "github.com/solv1k/croco-api/internal/modules/media/models"
	um "github.com/solv1k/croco-api/internal/modules/user/models"
	"github.com/solv1k/croco-api/internal/shared/models"
	"gorm.io/gorm"
)

type AdvertType struct {
	models.DictModel
}

type Advert struct {
	models.BaseModelWithTimestamps
	Title       string `json:"title" validate:"required,min=2,max=100" gorm:"size:255"`
	Description string `json:"description" validate:"required,min=2,max=1000"`
	Price       int    `json:"price" validate:"required,min=1,max=9999999"`
	// Foreign keys
	SellerID uuid.UUID `json:"seller_id" gorm:"type:uuid" validate:"required"`
	TypeID   uint      `json:"type_id" validate:"required" faker:"oneof:1,2,3"`
	// Relations
	Seller      um.User     `json:"seller"`
	Type        AdvertType  `json:"type"`
	MainImage   *mm.Media   `json:"main_image" gorm:"polymorphic:Model"`
	Screenshots []*mm.Media `json:"screenshots" gorm:"polymorphic:Model"`
}

// Hook to set images collection before create advert
func (a *Advert) BeforeCreate(tx *gorm.DB) error {
	a.BaseModelWithTimestamps.BeforeCreate(tx)
	if a.MainImage != nil {
		a.MainImage.Collection = "main_image"
	}
	for _, screenshot := range a.Screenshots {
		if screenshot == nil {
			continue
		}
		screenshot.Collection = "screenshots"
	}
	return nil
}
