package main

import (
	"GravitumTest/internal/api"
	"GravitumTest/internal/repository/postgres"
	"os"
)

func main() {
	connStr := os.Getenv("DB_URL")
	postgres.Connect(connStr)
	port := os.Getenv("port")
	api.StartRouter(port)
}
