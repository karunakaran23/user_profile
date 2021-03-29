package main

import (
	"log"
	"net/http"
	"os"
	"user_profile/database"
	"user_profile/handler"
	"user_profile/server"
)

var (
	port         = "8082"
	databaseFile = "./sqliteDB/userprofile.db"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DATABASE_FILE_NAME"); env != "" {
		databaseFile = env
	}
}

func main() {
	db, err := database.InitDB(databaseFile)
	if err != nil {
		panic(err)
	}
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{
		DB: db,
	}
	server.InitRoute(&h)
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}
