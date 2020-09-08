package controllers

import (
	"car-mysql/models"
	"car-mysql/usecases"
	"car-mysql/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TransactionController struct {
	transactionUseCase usecases.TransactionUseCase
}

func CreateTransactionController(r *mux.Router, transactionUseCase usecases.TransactionUseCase) {

	transactionController := TransactionController{transactionUseCase}

	r.HandleFunc("/transactions", transactionController.getAllTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transaction", transactionController.CreateTransaction).Methods(http.MethodPost)
	s := r.PathPrefix("/transaction").Subrouter()
	s.HandleFunc("/{id}", transactionController.deleteTransaction).Methods(http.MethodDelete)
	s.HandleFunc("/{id}", transactionController.getTransactionByID).Methods(http.MethodGet)
	s.HandleFunc("/{id}", transactionController.updateTransaction).Methods(http.MethodPut)
}

func (t TransactionController) updateTransaction(resp http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest, err.Error())
	}
	var body models.Transaction
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest, err.Error())
	}
	transaction, err := t.transactionUseCase.UpdateTransaction(id, &body)
	if err != nil {
		utils.HandleError(resp, http.StatusBadGateway, err.Error())
	}
	utils.HandleSuccess(resp, http.StatusOK, transaction)
}

func (t TransactionController) getTransactionByID(resp http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest, err.Error())
	}
	transaction, err := t.transactionUseCase.GetTransactionByID(id)
	if err != nil {
		utils.HandleError(resp, http.StatusBadGateway, err.Error())
	}
	utils.HandleSuccess(resp, http.StatusOK, transaction)
}

func (t TransactionController) deleteTransaction(resp http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest, err.Error())
	}
	err = t.transactionUseCase.DeleteTransaction(id)
	if err != nil {
		utils.HandleError(resp, http.StatusBadGateway, err.Error())
	}
	utils.HandleSuccess(resp, http.StatusOK, nil)
}
func (t TransactionController) getAllTransaction(resp http.ResponseWriter, r *http.Request) {
	transactions, err := t.transactionUseCase.GetAllTransaction()
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, err.Error())
	}
	utils.HandleSuccess(resp, http.StatusOK, transactions)
}

func (t TransactionController) CreateTransaction(resp http.ResponseWriter, r *http.Request) {
	var body models.Transaction
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest, err.Error())
	}
	transaction, err := t.transactionUseCase.CreateTransaction(&body)
	if err != nil {
		utils.HandleError(resp, http.StatusBadGateway, err.Error())
	}

	utils.HandleSuccess(resp, http.StatusOK, transaction)
}
