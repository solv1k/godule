package seeders

import (
	"errors"

	catalog "github.com/solv1k/croco-api/internal/modules/catalog/seeders"
	user "github.com/solv1k/croco-api/internal/modules/user/seeders"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, key string, count int) error {
	switch key {
	case "user":
		return user.Run(db, count)
	case "catalog":
		return catalog.Run(db, count)
	default:
		return errors.New("seeder with \"" + key + "\" key has not been implemented")
	}
}
