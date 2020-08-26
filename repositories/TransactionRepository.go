package repositories

import "car-mysql/models"

type TransactionRepository interface {
	GetAllTransaction() (*[]models.Transaction, error)
	GetTransactionByID(id int) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransaction(id int, transaction *models.Transaction) (*models.Transaction, error)
	DeleteTransaction(id int) error
}
