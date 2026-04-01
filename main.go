package main

import "go-api/internal/db"

func main() {
	db.ConnectDB()
	db.Migrate()
}
