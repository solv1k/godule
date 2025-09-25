# Godule

## Overview

This package provides a modular API example built on top of [Fiber](https://github.com/gofiber/fiber) web framework for Go. The framework is designed to support modular architecture, allowing developers to easily create and manage separate API modules.

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
            ...
        },
    }
    
    // Initialize and run API
    app := api.New(config)
    app.Run(":3000")
```


## API Methods

-   **New()** — creates a new API server instance
    
-   **Run()** — starts the server on the specified address
    
-   **RunMigrations()** — executes migrations for all modules
    
-   **initRoutes()** — initializes API routes
    

## Rate Limiting

The framework includes a built-in rate limiter configured to allow 100 requests per minute. This can be adjusted as needed.

## Versioning

The base route can be customized to support different API versions. By default, it uses `/api/v1`.

## Error Handling

The framework provides basic error handling mechanisms. Custom error handling can be added to individual modules as required.

## Extending the Framework

To add new functionality, simply create a new module that implements the `Module` interface and register it in the API configuration.

## Environment Configuration

The application uses environment variables for configuration, loaded via the `godotenv` package. This allows for easy configuration management across different environments.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.