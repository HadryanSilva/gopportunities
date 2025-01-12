package config

import (
	"github.com/HadryanSilva/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := "host=localhost user=postgres password=postgres dbname=postgres search_path=gopportunities"

	logger.Info("Connecting to Postgres")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errf("Failed to connect to Postgres: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("Failed to migrate schemas: %v", err)
		return nil, err
	}

	return db, nil
}
