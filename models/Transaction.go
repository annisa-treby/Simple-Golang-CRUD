package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Total int64  `json:"total"`
	Cars  []*Car `gorm:"many2many:transaction_cars;association_autoupdate:false;association_autoucreate:false"`
}
