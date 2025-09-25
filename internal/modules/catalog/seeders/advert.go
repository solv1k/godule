package seeders

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	"gorm.io/gorm"
)

// Generate fake adverts
func Seed(db *gorm.DB, count int) error {
	if count <= 0 {
		return nil
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	for i := 1; i <= count; i++ {
		var advert models.Advert
		if err := faker.FakeData(&advert, options.WithFieldsToIgnore("Type", "Screenshots")); err != nil {
			tx.Rollback()
			return fmt.Errorf("error generating fake data: %w", err)
		}

		if result := tx.Create(&advert); result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("error creating advert: %w", result.Error)
		}
	}

	return tx.Commit().Error
}
