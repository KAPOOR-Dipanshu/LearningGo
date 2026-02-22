package routes

import (
	"go-api-app/src/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employees", handlers.FetchAllEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", handlers.FetchEmployeeById).Methods("GET")
	router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handlers.PatchEmployee).Methods("PATCH")
	router.HandleFunc("/employees/{id}", handlers.RemoveEmployee).Methods("DELETE")
	return router
}
