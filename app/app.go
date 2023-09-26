package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ivkoandrv/simple-rest-go/api"
	"github.com/ivkoandrv/simple-rest-go/config"
	"github.com/ivkoandrv/simple-rest-go/database"
	"log"
	"net/http"
	"os"
)

func Run() error {
	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("error loading environment configuration: %v", err)
	}

	db, err := database.ConnectToDb()
	if err != nil {
		return fmt.Errorf("error initializing the database: %v", err)
	}

	dbName := db.GetDBName()

	if err := startServer(dbName); err != nil {
		return fmt.Errorf("error starting the server: %v", err)
	}

	return nil
}

func startServer(dbName string) error {
	router := mux.NewRouter()
	api.SetupRoutes(router)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Server is running on port %s. Connected to database: %s\n", port, dbName)
	return http.ListenAndServe(port, router)
}
