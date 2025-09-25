package database

import "os"

// Database provider
type Provider string

// Database providers supported
const (
	Postgres Provider = "postgres"
	Mysql    Provider = "mysql"
)

// Database configuration
type Config struct {
	Provider Provider
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Timezone string
	SSLMode  bool
}

// Returns the default database configuration
func ConfigDefault() Config {
	return Config{
		Provider: Postgres,
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Timezone: os.Getenv("DB_TIMEZONE"),
		SSLMode:  os.Getenv("DB_SSL_MODE") == "true" || os.Getenv("DB_SSL_MODE") == "1" || os.Getenv("DB_SSL_MODE") == "enable",
	}
}
