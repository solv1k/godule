package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Returns the default database connection
func Default() (*gorm.DB, error) {
	return Connect(ConfigDefault(), nil)
}

// Connects to the database
func Connect(config Config, gormConfig *gorm.Config) (*gorm.DB, error) {
	switch config.Provider {
	case Postgres:
		return connectPostgres(config, gormConfig)
	case Mysql:
		return nil, errors.New("mysql is not supported yet")
	default:
		return nil, errors.New("unknown database provider")
	}
}

// Connects to the postgres database
func connectPostgres(config Config, gormConfig *gorm.Config) (*gorm.DB, error) {
	if gormConfig == nil {
		gormConfig = &gorm.Config{}
	}

	dsn := buildDsnString(config)
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Builds the database connection string
func buildDsnString(config Config) string {
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Name + " port=" + config.Port
	if config.Timezone != "" {
		dsn += " TimeZone=" + config.Timezone
	}
	if !config.SSLMode {
		dsn += " sslmode=disable"
	}

	return dsn
}
