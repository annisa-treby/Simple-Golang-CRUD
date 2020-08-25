package repositories

import "car-mysql/models"

type StoreRepository interface {

	GetAllStore() (*[]models.Store,error)
	InsertNewStore(store *models.Store) (*models.Store,error)
	GetStoreById(id int)(*models.Store,error)
	UpdateStore(id int, store *models.Store)(*models.Store,error)
	DeleteStore(id int)error
}
