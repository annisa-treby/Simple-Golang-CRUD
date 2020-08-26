package main

import (
	"car-mysql/config"
	"car-mysql/controllers"
	"car-mysql/repositories"
	"car-mysql/usecases"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	db := config.Config()
	defer db.Close()
	config.Migrate(db)

	router := mux.NewRouter().StrictSlash(true)

	carRepo := repositories.CreateCarRepository(db)
	carUseCase := usecases.CreateCarUseCase(carRepo)
	controllers.CreateCarController(router, carUseCase)

	storeRepo := repositories.CreateStoreRepository(db)
	storeUseCase := usecases.CreateStoreUseCase(storeRepo)
	controllers.CreateStoreController(router, storeUseCase)

	transactionRepo := repositories.CreateTransactionRepository(db)
	transactionUseCase := usecases.CreateTransactionUseCase(transactionRepo)
	controllers.CreateTransactionController(router, transactionUseCase)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
