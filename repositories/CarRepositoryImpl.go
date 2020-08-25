package repositories

import (
	"car-mysql/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type CarRepositoryImpl struct {
	db *gorm.DB
}

func CreateCarRepository(db *gorm.DB) CarRepository {
	return &CarRepositoryImpl{db }
}

func (c CarRepositoryImpl) SearchCarBasedOnStore(id int) (*[]models.Car, error) {
	var cars []models.Car
	err := c.db.Find(&cars, "store_id = ?", id).Error
	if err != nil {
		log.Fatal(err)
	}
	return &cars,err
}

func (c CarRepositoryImpl) GetAllCars() (*[]models.Car, error) {
	var cars []models.Car
	err := c.db.Find(&cars).Error
	if err != nil {
		log.Fatal(err)
	}
	return &cars,err
}

func (c CarRepositoryImpl) GetCarById(id int) (*models.Car, error) {
	var car models.Car
	err := c.db.First(&car, id).Error
	if err!=nil {
		log.Fatal(err)
	}
	return &car,err
}

func (c CarRepositoryImpl) InsertNewCar(car *models.Car) (*models.Car, error) {
	err := c.db.Save(&car).Error
	if err != nil {
		log.Fatal(err)
	}
	return car,err
}

func (c CarRepositoryImpl) UpdateCar(id int, car *models.Car) (*models.Car, error) {
	err := c.db.Model(&car).Where("id=?",id).Update(car).Error
	if err != nil {
		log.Fatal(err)
	}
	return car,err
}

func (c CarRepositoryImpl) DeleteCar(id int) error {
	car := models.Car{}
	err := c.db.Table("car").Where("id = ?", id).First(&car).Delete(&car).Error
	fmt.Println(err)
	if err != nil{
		log.Fatal(err,"repo")
	}
	return nil
}
