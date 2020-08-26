package controllers

import (
	"car-mysql/usecases"
	"car-mysql/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionController struct {
	transactionUseCase usecases.TransactionUseCase
}

func CreateTransactionController(r *mux.Router, transactionUseCase usecases.TransactionUseCase) {

	transactionController := TransactionController{transactionUseCase}

	r.HandleFunc("/transactions", transactionController.getAllTransaction).Methods(http.MethodGet)
}

func (t TransactionController) getAllTransaction(resp http.ResponseWriter, r *http.Request) {
	transactions, err := t.transactionUseCase.GetAllTransaction()
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, err.Error())
	}
	utils.HandleSuccess(resp, http.StatusOK, transactions)
}
