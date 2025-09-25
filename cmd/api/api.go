package api

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/shared/middleware"
)

// API
type Api struct {
	Fiber  *fiber.App
	Config Config
}

// API configuration
type Config struct {
	AppName     string
	BaseRoute   string
	AutoMigrate bool
	Modules     []Module
}

// API module interface
type Module interface {
	Name() string
	Description() string
	Version() string
	Migrate() error
	Routes(router fiber.Router)
}

// API constructor
func New(config Config) *Api {
	app := &Api{
		Fiber: fiber.New(fiber.Config{
			AppName: config.AppName,
		}),
		Config: config,
	}
	app.initRoutes()
	if config.AutoMigrate {
		if err := app.RunMigrations(); err != nil {
			log.Fatal(err)
		}
	}
	return app
}

// Run API server on specified address
func (a *Api) Run(addr string) error {
	return a.Fiber.Listen(addr)
}

// Run database migrations
func (a *Api) RunMigrations() error {
	for _, module := range a.Config.Modules {
		if err := module.Migrate(); err != nil {
			return err
		}
	}
	return nil
}

// Initialize API routes
func (a *Api) initRoutes() {
	apiLimiter := middleware.NewApiLimiter(100, time.Minute)
	v1 := a.Fiber.Group(a.Config.BaseRoute).Use(apiLimiter)

	for _, module := range a.Config.Modules {
		module.Routes(v1)
	}
}
