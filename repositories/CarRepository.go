package repositories

import "car-mysql/models"

type CarRepository interface {
	GetAllCars() (*[]models.Car,error)
	GetCarById(id int) (*models.Car,error)
	InsertNewCar( car *models.Car) (*models.Car,error)
	UpdateCar(id int,car *models.Car) (*models.Car,error)
	DeleteCar(id int) error
	SearchCarBasedOnStore(id int) (*[]models.Car,error)
}
