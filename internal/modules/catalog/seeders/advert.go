package seeders

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	mm "github.com/solv1k/croco-api/internal/modules/media/models"
	"gorm.io/gorm"
)

// Generate fake adverts
func Run(db *gorm.DB, count int) error {
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
		// Generate fake advert
		var advert models.Advert
		if err := faker.FakeData(&advert, options.WithFieldsToIgnore("Type", "Screenshots")); err != nil {
			tx.Rollback()
			return fmt.Errorf("error generating fake data: %w", err)
		}

		// Generate fake advert screenshots
		advert.Screenshots = generateScreenshots(5)

		// Create advert with screenshots
		if result := tx.Create(&advert); result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("error creating advert: %w", result.Error)
		}
	}

	return tx.Commit().Error
}

func generateScreenshots(count int) []*mm.Media {
	screenshots := make([]*mm.Media, count)
	for i := 0; i < count; i++ {
		screenshots[i] = &mm.Media{
			Url: fmt.Sprintf("https://images.storage/screenshots/%d", i+1),
		}
	}
	return screenshots
}
