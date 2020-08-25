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

type CarController struct {
	carUseCase usecases.CarUseCase
}

func CreateCarController(r *mux.Router, carUseCase usecases.CarUseCase)  {
	carController := CarController{carUseCase}

	r.HandleFunc("/cars", carController.getAllCars).Methods(http.MethodGet)
	r.HandleFunc("/cars/store/{id}", carController.getAllCarsByStore).Methods(http.MethodGet)
	r.HandleFunc("/car",carController.insertNeCar).Methods(http.MethodPost)
	s := r.PathPrefix("/car").Subrouter()
	s.HandleFunc("/{id}", carController.getCarByID).Methods(http.MethodGet)
	s.HandleFunc("/{id}",carController.updateCar).Methods(http.MethodPut)
	s.HandleFunc("/{id}",carController.deleteCar).Methods(http.MethodDelete)
}

func (c CarController) getAllCars(resp http.ResponseWriter, req *http.Request) {
	cars,err:=c.carUseCase.GetAllCars()
	if err != nil {
		log.Fatal(err)
	}
	utils.HandleSuccess(resp, http.StatusOK,cars)
}
func (c CarController) getAllCarsByStore(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil{
		log.Fatal(err)
	}
	cars,err:=c.carUseCase.SearchCarBasedOnStore(id)
	if err != nil {
		log.Fatal(err)
	}
	utils.HandleSuccess(resp, http.StatusOK,cars)
}

func (c CarController) getCarByID(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	car,err := c.carUseCase.GetCarById(id)
	if err != nil {
		log.Fatal(err)
	}

	utils.HandleSuccess(resp,http.StatusOK,car)
}

func (c CarController) insertNeCar(resp http.ResponseWriter, req *http.Request) {
	var car models.Car

	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		log.Fatal(err)
	}

	response,err := c.carUseCase.InsertNewCar(&car)
	if err != nil {
		log.Fatal(err)
	}

	utils.HandleSuccess(resp,http.StatusOK,response)
}

func (c CarController) updateCar(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		log.Fatal(err)
	}

	var car models.Car
	err = json.NewDecoder(req.Body).Decode(&car)
	if err!=nil {
		log.Fatal(err)
	}

	response,err := c.carUseCase.UpdateCar(id,&car)
	if err != nil {
		log.Fatal(err)
	}

	utils.HandleSuccess(resp,http.StatusOK,response)
}

func (c CarController) deleteCar(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id,err := strconv.Atoi(params["id"])
	if err!=nil {
		log.Fatal(err,"convert id")
	}
	err = c.carUseCase.DeleteCar(id)
	if err!=nil {
		log.Fatal(err,"handler")
	}
	utils.HandleSuccess(resp,http.StatusOK,err)
}
