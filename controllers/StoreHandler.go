package controllers

import (
	"car-mysql/models"
	"car-mysql/usecases"
	"car-mysql/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type StoreController struct {
	storeUseCase usecases.StoreUseCase
}

func CreateStoreController(r *mux.Router, storeUseCase usecases.StoreUseCase)  {

	storeController := StoreController{storeUseCase}

	r.HandleFunc("/stores", storeController.getAllStore).Methods(http.MethodGet)
	r.HandleFunc("/store", storeController.insertNewStore).Methods(http.MethodPost)
	s := r.PathPrefix("/store").Subrouter()
	s.HandleFunc("/{id}",storeController.UpdateStore).Methods(http.MethodPut)
	s.HandleFunc("/{id}",storeController.DeleteStore).Methods(http.MethodDelete)
	s.HandleFunc("/{id}",storeController.getStoreById).Methods(http.MethodGet)
}

func (s StoreController) DeleteStore(resp http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err!= nil{
		log.Fatal(err)
	}

	err = s.storeUseCase.DeleteStore(id)
	utils.HandleSuccess(resp,http.StatusOK,err)
}

func (s StoreController) UpdateStore(resp http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err!= nil{
		log.Fatal(err)
	}

	var store models.Store
	err = json.NewDecoder(r.Body).Decode(&store)
	if err!=nil {
		log.Fatal(err)
	}

	response,err := s.storeUseCase.UpdateStore(id,&store)
	if err!=nil {
		log.Fatal(err)
	}

	utils.HandleSuccess(resp,http.StatusOK,response)
}

func (s StoreController) getStoreById(resp http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		utils.HandleError(resp, http.StatusBadRequest,"Path error")
	}

	store,err := s.storeUseCase.GetStoreById(id)
	if err!=nil {
		utils.HandleError(resp, http.StatusBadGateway,err.Error())
	}
	utils.HandleSuccess(resp,http.StatusOK,store)
}

func (s StoreController) getAllStore(resp http.ResponseWriter, r *http.Request) {
	stores, err := s.storeUseCase.GetAllStore()
	if err != nil {
		log.Fatal(err)
	}
	utils.HandleSuccess(resp,http.StatusOK,stores)
}

func (s StoreController) insertNewStore(resp http.ResponseWriter, r *http.Request) {
	var store models.Store

	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	body, err := s.storeUseCase.InsertNewStore(&store)
	if err != nil {
		log.Fatal(err)
	}
	utils.HandleSuccess(resp,http.StatusOK,body)
}
