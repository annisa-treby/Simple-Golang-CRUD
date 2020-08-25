package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	CarType string `json:"car_type"`
	Color string `json:"color"`
	Price string `json:"price"`
	CarModel string `json:"car_model"`
	CarModelYear int `json:"car_model_year"`
	Availability bool `json:"availability"`
	StoreID uint
}

func (c Car) TableName() string {
	return "car"
}


