# Godule

## Overview

This package provides a modular API example built on top of [Fiber](https://github.com/gofiber/fiber) web framework for Go. The framework is designed to support modular architecture, allowing developers to easily create and manage separate API modules.

- [Key features](#key-features)

- [Core components](#core-components)

- [Api struct](#api-struct)

- [Config struct](#config-struct)

- [Module interface](#module-interface)

- [Implementing a new module](#implementing-a-new-module)

- [Usage example](#usage-example)

- [Database migrations](#database-migrations)

- [Database seeding](#database-seeding)

- [Configuration options](#configuration-options)

- [Rate limiting](#rate-limiting)

- [Versioning](#versioning)

- [Error handling](#error-handling)

- [License](#license)

## Key Features

-   **Modular Architecture** — supports dynamic registration of API modules
    
-   **Customizable Base Route** — allows setting a custom base route for API endpoints
    
-   **Rate Limiting** — built-in request limiter for API protection
    
-   **Migration Support** — automatic migration handling for all modules
    
-   **Extensibility** — easy to add new modules and features
    

## Core Components

### Api Struct

The main struct representing the API server:

-   **Fiber** — Fiber application instance
    
-   **Config** — API configuration settings
    

### Config Struct

Configuration structure for the API:

-   **AppName** — name of the application
    
-   **BaseRoute** — base route for API endpoints
    
-   **Modules** — list of registered modules
    

### Module Interface

Every module must implement the following methods:

-   **Name()** — returns the module name
    
-   **Description()** — provides a description of the module
    
-   **Version()** — returns the module version
    
-   **Migrate()** — performs database migrations
    
-   **Routes()** — registers API routes


## Implementing a new module

To create a new module, implement the `Module` interface and register it in the API configuration.

```go
    type Module struct {
        // ...
    }

    // Register the module with all dependencies
    func NewModule(db *gorm.DB) *MyModule {
        return &Module{
            // ...
        }
    }
    
    func (m *Module) Name() string {
        return "my_module"
    }

    func (m *Module) Description() string {
        return "My Module"
    }

    func (m *Module) Version() string {
        return "1.0.0"
    }

    func (m *Module) Migrate() error {
        return m.DB.AutoMigrate(&FirstModel{}, &SecondModel{}) // Replace with your models
    }

    func (m *Module) Routes(app *fiber.App) {
        moduleGroup := app.Group("/my_module")
        moduleGroup.Get("/", m.Handler) // Replace with your handler
    }
}
```


## Configuration Options

The API can be configured with the following parameters:

-   **Application Name** — sets the name of the application
    
-   **Base Route** — defines the base path for all API endpoints
    
-   **Modules** — list of modules to register
    

## Usage Example

```go
    // Create configuration
    config := api.Config{
        AppName: "My API",
        BaseRoute: "/api/v1",
        Modules: api.Module{
            mediaModule,
            userModule,
            catalogModule,
            anotherAwesomeModule,
            ...
        },
    }
    
    // Initialize and run API
    app := api.New(config)
    app.Run(":3000")
``` 

## Database Migrations

The framework supports automatic database migrations for all registered modules. To enable migrations, set the `AutoMigrate` field in the API configuration to `true`. This will automatically run migrations for all registered modules.

Dont forget implementing `Migrate()` method in your module.

## Database Seeding

For seeding database you can use seeding console command.

```bash
go run cmd/seed/main.go [key] [count]
```

The `key` parameter is the name of the seeder, and the `count` parameter is the number of records to create.

### Usage Example

For example:

```bash
go run cmd/seed/main.go user 10
```

This command will create 10 users in the database.

### Seeder Interface

```go
type Seeder interface {
    Run(count int) error
}
```

### Registering Seeders

You can register new seeders by implementing the `Seeder` interface and registering them in `cmd/seed/seeders/registrator.go` inside the `Seeders()` method.


## Rate Limiting

The framework includes a built-in rate limiter configured to allow 100 requests per minute. This can be adjusted as needed.

## Versioning

The base route can be customized to support different API versions. By default, it uses `/api/v1`.

## Error Handling

The framework provides basic error handling mechanisms. Custom error handling can be added to individual modules as required.

## Environment Configuration

The application uses environment variables for configuration, loaded via the `godotenv` package. This allows for easy configuration management across different environments.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.