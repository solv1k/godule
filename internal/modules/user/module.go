package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/user/handlers"
	"github.com/solv1k/croco-api/internal/modules/user/models"
	"github.com/solv1k/croco-api/pkg/auth/fiber/middleware"
	"gorm.io/gorm"
)

// User module
type Module struct {
	DB          *gorm.DB
	AuthHandler *handlers.AuthHandler
	Middleware  struct {
		LoginLimiter    fiber.Handler
		SendCodeLimiter fiber.Handler
		Jwt             fiber.Handler
	}
}

// Runs a database migration for the current module
func (m *Module) Migrate() error {
	return m.DB.AutoMigrate(&models.User{})
}

// User module constructor
func NewModule(db *gorm.DB) *Module {
	module := &Module{
		DB:          db,
		AuthHandler: handlers.NewAuthHandler(db),
	}
	module.initMiddleware()
	return module
}

// Initializes middleware
func (m *Module) initMiddleware() {
	m.Middleware.LoginLimiter = middleware.NewLoginLimiter(10, time.Minute)
	m.Middleware.SendCodeLimiter = middleware.NewSendcodeLimiter(5, time.Minute)
	m.Middleware.Jwt = middleware.NewJwtMiddleware()
}

// Returns module name
func (m *Module) Name() string {
	return "user"
}

// Returns module description
func (m *Module) Description() string {
	return "User module"
}

// Returns module version
func (m *Module) Version() string {
	return "1.0.0"
}

// Registers module routes
func (m *Module) Routes(router fiber.Router) {
	userGroup := router.Group("/user")
	{
		authGroup := userGroup.Group("/auth")
		authGroup.Post("/send-code", m.Middleware.SendCodeLimiter, m.AuthHandler.SendCode)
		authGroup.Post("/login", m.Middleware.LoginLimiter, m.AuthHandler.Login)
		authGroup.Post("/logout", m.AuthHandler.Logout)
		authGroup.Get("/me", m.Middleware.Jwt, m.AuthHandler.Me)
	}
}
