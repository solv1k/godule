package seller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/seller/handlers"
	"github.com/solv1k/croco-api/pkg/auth/fiber/middleware"
	"gorm.io/gorm"
)

// Seller module
type Module struct {
	DB            *gorm.DB
	AdvertHandler *handlers.AdvertHandler
	Middleware    struct {
		Jwt fiber.Handler
	}
}

// Seller module constructor
func NewModule(db *gorm.DB) *Module {
	module := &Module{
		DB:            db,
		AdvertHandler: handlers.NewAdvertHandler(db),
	}

	module.Middleware.Jwt = middleware.NewJwtMiddleware()

	return module
}

// Runs a database migration for the current module
func (m *Module) Migrate() error {
	return nil
}

// Returns module name
func (m *Module) Name() string {
	return "seller"
}

// Returns module description
func (m *Module) Description() string {
	return "Seller module"
}

// Returns module version
func (m *Module) Version() string {
	return "1.0.0"
}

// Registers module routes
func (m *Module) Routes(router fiber.Router) {
	userGroup := router.Group("/seller").Use(m.Middleware.Jwt)
	{
		advertsGroup := userGroup.Group("/adverts")
		advertsGroup.Get("/", m.AdvertHandler.GetAdverts)
		advertsGroup.Post("/", m.AdvertHandler.CreateAdvert)
		advertsGroup.Put("/:id", m.AdvertHandler.UpdateAdvert)
		advertsGroup.Delete("/:id", m.AdvertHandler.DeleteAdvert)
	}
}
