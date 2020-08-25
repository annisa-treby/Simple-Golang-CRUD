package usecases

import (
	"car-mysql/models"
	"car-mysql/repositories"
)

type StoreUseCaseImpl struct {
	storeRepo repositories.StoreRepository
}

func CreateStoreUseCase(storeRepo repositories.StoreRepository) StoreUseCase {
	return &StoreUseCaseImpl{storeRepo}
}

func (s StoreUseCaseImpl) GetStoreById(id int) (*models.Store, error) {
	return s.storeRepo.GetStoreById(id)
}

func (s StoreUseCaseImpl) UpdateStore(id int, store *models.Store) (*models.Store, error) {
	return s.storeRepo.UpdateStore(id,store)
}

func (s StoreUseCaseImpl) DeleteStore(id int) error {
	return s.storeRepo.DeleteStore(id)
}

func (s StoreUseCaseImpl) GetAllStore() (*[]models.Store, error) {
	return s.storeRepo.GetAllStore()
}

func (s StoreUseCaseImpl) InsertNewStore(store *models.Store) (*models.Store, error) {
	return s.storeRepo.InsertNewStore(store)
}
