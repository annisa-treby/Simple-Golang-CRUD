package repositories

import (
	"car-mysql/models"
	"github.com/jinzhu/gorm"
	"log"
)

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func CreateStoreRepository(db *gorm.DB) StoreRepository {
	return &StoreRepositoryImpl{db}
}

func (s StoreRepositoryImpl) GetStoreById(id int) (*models.Store, error) {
	var store models.Store
	err := s.db.Preload("Cars").First(&store, id).Error
	if err != nil {
		log.Fatal(err)
	}
	return &store, err
}

func (s StoreRepositoryImpl) UpdateStore(id int, store *models.Store) (*models.Store, error) {
	err := s.db.Model(&store).Where("id = ?", id).Update(store).Error
	if err != nil {
		log.Fatal(err)
	}
	return store, err
}

func (s StoreRepositoryImpl) DeleteStore(id int) error {
	store := models.Store{}
	err := s.db.Table("store").Where("id = ?", id).First(&store).Delete(&store).Error
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s StoreRepositoryImpl) GetAllStore() (*[]models.Store, error) {
	var stores []models.Store
	err := s.db.Preload("Cars").Find(&stores).Error
	if err != nil {
		log.Fatal(err)
	}
	return &stores, err
}

func (s StoreRepositoryImpl) InsertNewStore(store *models.Store) (*models.Store, error) {
	err := s.db.Save(&store).Error
	if err != nil {
		log.Fatal(err)
	}
	return store, err
}
