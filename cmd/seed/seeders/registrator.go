package seeders

import (
	"errors"

	"github.com/solv1k/croco-api/cmd/seed/types"
	"github.com/solv1k/croco-api/database"
	catalogSeeder "github.com/solv1k/croco-api/internal/modules/catalog/seeders"
	userSeeder "github.com/solv1k/croco-api/internal/modules/user/seeders"
	"gorm.io/gorm"
)

type Registrator struct {
	DB *gorm.DB
}

func NewRegistrator() *Registrator {
	db, err := database.Default()
	if err != nil {
		panic(err)
	}

	return &Registrator{
		DB: db,
	}
}

// Returning all seeders
func (r *Registrator) Seeders() map[string]func() types.Seeder {
	return map[string]func() types.Seeder{
		"catalog": func() types.Seeder {
			return catalogSeeder.New(r.DB)
		},
		"user": func() types.Seeder {
			return userSeeder.New(r.DB)
		},
	}
}

// Registering seeder by key and returning it
func (r *Registrator) Get(key string) (types.Seeder, error) {
	// Checking if seeder with key has been implemented
	seeders := r.Seeders()
	if _, ok := seeders[key]; !ok {
		return nil, errors.New("seeder with \"" + key + "\" key has not been implemented")
	}

	// Returning seeder by key
	return seeders[key](), nil
}
