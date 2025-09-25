package media

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/media/models"
	"gorm.io/gorm"
)

// Media module
type Module struct {
	DB *gorm.DB
}

// Runs a database migration for the current module
func (m *Module) Migrate() error {
	return m.DB.AutoMigrate(&models.Media{})
}

// Media module constructor
func NewModule(db *gorm.DB) *Module {
	return &Module{DB: db}
}

// Returns module name
func (m *Module) Name() string {
	return "media"
}

// Returns module description
func (m *Module) Description() string {
	return "Media module"
}

// Returns module version
func (m *Module) Version() string {
	return "1.0.0"
}

// Registers module routes
func (m *Module) Routes(router fiber.Router) {
	//
}
