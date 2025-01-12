package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Initialize Postgres
	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("failed to initialize Postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = newLogger(p)
	return logger
}
