package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"policyAuth/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	InitKeycloak()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbService := database.New()
	database.InitSchema(database.DBInstance.DB) // Initialize the database schema

	NewServer := &Server{
		port: port,
		db:   dbService,
	}
	fmt.Println("Server is running on port: ", port)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
