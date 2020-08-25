package usecases

import (
	"car-mysql/models"
	"car-mysql/repositories"
)

type CarUseCaseImpl struct {
	carRepo repositories.CarRepository
}

func (c CarUseCaseImpl) SearchCarBasedOnStore(id int) (*[]models.Car, error) {
	return c.carRepo.SearchCarBasedOnStore(id)
}

func CreateCarUseCase(carRepo repositories.CarRepository) CarUseCase {
	return &CarUseCaseImpl{carRepo}
}

func (c CarUseCaseImpl) GetAllCars() (*[]models.Car, error) {
	return c.carRepo.GetAllCars()
}

func (c CarUseCaseImpl) GetCarById(id int) (*models.Car, error) {
	return c.carRepo.GetCarById(id)
}

func (c CarUseCaseImpl) InsertNewCar(car *models.Car) (*models.Car, error) {
	return c.carRepo.InsertNewCar(car)
}

func (c CarUseCaseImpl) UpdateCar(id int, car *models.Car) (*models.Car, error) {
	return c.carRepo.UpdateCar(id, car)
}

func (c CarUseCaseImpl) DeleteCar(id int) error {
	return c.carRepo.DeleteCar(id)
}
