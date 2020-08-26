package config

import (
	"car-mysql/models"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&models.Car{},
		&models.Store{},
		&models.Transaction{})

	db.Model(&models.Car{})
	db.Model(&models.Store{})
	db.Model(&models.Transaction{})
}
