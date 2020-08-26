package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Cars []*Car `gorm:"many2many:transaction_cars;"`
}
