package main

import (
	"log"

	"github.com/gorilla/mux"

	"backend/database"
	"backend/server"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		panic(err.Error())
	}
	defer db.Close()

	srv := &server.Server{
		Db:     db,
		Router: mux.NewRouter(),
		Port:   9080,
	}
	srv.Init()

	log.Printf("Starting the HTTP server on port %d\n", srv.Port)
	srv.ServeHTTP()
}
