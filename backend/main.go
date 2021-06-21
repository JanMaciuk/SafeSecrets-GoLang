package main

import (
	"log"
	"os"
	"sb14project_backend/pkg/api"
	db2 "sb14project_backend/pkg/database"
)

func main() {
	db, err := db2.NewDB("app.db")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatalln(api.NewServer(":"+port, db))
}
