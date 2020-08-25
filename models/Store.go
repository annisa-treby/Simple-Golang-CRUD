package models

import "github.com/jinzhu/gorm"

type Store struct {
	gorm.Model
	StoreName string `json:"store_name"`
	Cars []Car
}

func (c Store) TableName() string {
	return "store"
}
