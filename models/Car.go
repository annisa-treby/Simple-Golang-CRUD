package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	CarType      string `json:"car_type"`
	Color        string `json:"color"`
	Price        int64  `json:"price"`
	CarModel     string `json:"car_model"`
	CarModelYear int    `json:"car_model_year"`
	Availability bool   `json:"availability"`
	StoreID      uint
	Transaction  []*Transaction `gorm:"many2many:transaction_cars;"`
}

func (c Car) TableName() string {
	return "car"
}
