package main

import (
	"context"
	"go-api-app/src/constants"
	"go-api-app/src/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := routes.SetupRoutes()
	PORT := constants.GetConstant("PORT")

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	go func() {
		log.Println("Starting server on port " + PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
