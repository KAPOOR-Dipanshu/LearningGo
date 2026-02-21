package routes

import (
	"go-api-app/src/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/mongodb-data", handlers.GetDataFromMongoDB).Methods("GET")
	return router
}
