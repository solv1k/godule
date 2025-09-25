package catalog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/catalog/handlers"
	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	"gorm.io/gorm"
)

// Catalog module
type Module struct {
	DB            *gorm.DB
	AdvertHandler *handlers.AdvertHandler
}

// Runs a database migration for the current module
func (m *Module) Migrate() error {
	if err := m.DB.AutoMigrate(&models.AdvertType{}, &models.Advert{}); err != nil {
		return err
	}

	// Initialize advert types (for seeding database)
	m.AdvertHandler.InitAdvertTypes()
	return nil
}

// Catalog module constructor
func NewModule(db *gorm.DB) *Module {
	return &Module{
		DB:            db,
		AdvertHandler: handlers.NewAdvertHandler(db),
	}
}

// Returns module name
func (m *Module) Name() string {
	return "catalog"
}

// Returns module description
func (m *Module) Description() string {
	return "Catalog module"
}

// Returns module version
func (m *Module) Version() string {
	return "1.0.0"
}

// Registers module routes
func (m *Module) Routes(router fiber.Router) {
	authGroup := router.Group("/catalog")
	authGroup.Get("/adverts", m.AdvertHandler.GetAdverts)
}
