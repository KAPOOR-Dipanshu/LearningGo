package main

import (
	"go-api-app/src/constants"
	"go-api-app/src/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()
	PORT := constants.GetConstant("PORT")
	log.Println("Starting server on port " + PORT)
	if err := http.ListenAndServe(":"+PORT, router); err != nil {
		log.Fatal(err)
	}
}
