package usecases

import (
	"car-mysql/models"
	"car-mysql/repositories"
)

type TransactionUseCaseImpl struct {
	transactionRepo repositories.TransactionRepository
}

func (t TransactionUseCaseImpl) GetAllTransaction() (*[]models.Transaction, error) {
	return t.transactionRepo.GetAllTransaction()
}

func (t TransactionUseCaseImpl) GetTransactionByID(id int) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionUseCaseImpl) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionUseCaseImpl) UpdateTransaction(id int, transaction *models.Transaction) (*models.Transaction, error) {
	panic("implement me")
}

func (t TransactionUseCaseImpl) DeleteTransaction(id int) error {
	panic("implement me")
}

func CreateTransactionUseCase(transactionRepo repositories.TransactionRepository) TransactionUseCase {
	return &TransactionUseCaseImpl{transactionRepo}
}
