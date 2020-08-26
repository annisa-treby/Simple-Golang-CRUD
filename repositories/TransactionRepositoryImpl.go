package repositories

import (
	"car-mysql/models"
	"github.com/jinzhu/gorm"
	"log"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func CreateTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db}
}

func (t TransactionRepositoryImpl) GetAllTransaction() (*[]models.Transaction, error) {
	var transactions []models.Transaction
	err := t.db.Preload("Cars").Find(&transactions).Error
	if err != nil {
		log.Fatal(err)
	}
	return &transactions, err
}

func (t TransactionRepositoryImpl) GetTransactionByID(id int) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionRepositoryImpl) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionRepositoryImpl) UpdateTransaction(id int, transaction *models.Transaction) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionRepositoryImpl) DeleteTransaction(id int) error {
	panic("implement me")
}
