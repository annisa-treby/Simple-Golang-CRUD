package usecases

import (
	"car-mysql/models"
	"car-mysql/repositories"
	"log"
)

type TransactionUseCaseImpl struct {
	transactionRepo repositories.TransactionRepository
	carUseCase      CarUseCase
}

func (t TransactionUseCaseImpl) GetAllTransaction() (*[]models.Transaction, error) {
	return t.transactionRepo.GetAllTransaction()
}

func (t TransactionUseCaseImpl) GetTransactionByID(id int) (*models.Transaction, error) {
	return t.transactionRepo.GetTransactionByID(id)
}

func (t TransactionUseCaseImpl) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	var total int64
	for _, car := range transaction.Cars {
		car, err := t.carUseCase.GetCarById(int(car.ID))
		if err != nil {
			log.Fatal(err)
		}
		total += car.Price
	}
	transaction.Total = total
	return t.transactionRepo.CreateTransaction(transaction)
}

func (t TransactionUseCaseImpl) UpdateTransaction(id int, transaction *models.Transaction) (*models.Transaction, error) {
	return t.transactionRepo.UpdateTransaction(id, transaction)
}

func (t TransactionUseCaseImpl) DeleteTransaction(id int) error {
	return t.transactionRepo.DeleteTransaction(id)
}

func CreateTransactionUseCase(transactionRepo repositories.TransactionRepository, carUseCase CarUseCase) TransactionUseCase {
	return &TransactionUseCaseImpl{transactionRepo, carUseCase}
}
