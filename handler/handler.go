package handler

import (
	"github.com/HadryanSilva/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)
