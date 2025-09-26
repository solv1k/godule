package seeders

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/solv1k/croco-api/internal/modules/user/models"
	"gorm.io/gorm"
)

// User seeder
type Seeder struct {
	DB *gorm.DB
}

// User seeder constructor
func New(db *gorm.DB) *Seeder {
	return &Seeder{
		DB: db,
	}
}

// Generate fake users
func (s *Seeder) Run(count int) error {
	if count <= 0 {
		return nil
	}

	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	for i := 1; i <= count; i++ {
		var user models.User
		if err := faker.FakeData(&user); err != nil {
			tx.Rollback()
			return fmt.Errorf("error generating fake data: %w", err)
		}

		if result := tx.Create(&user); result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("error creating advert: %w", result.Error)
		}
	}

	return tx.Commit().Error
}
