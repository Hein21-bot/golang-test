package handlers

import (
	"log"

	"github.com/surrealdb/surrealdb.go" // Make sure this import matches the actual library
)

var DB *surrealdb.DB

func InitDB() {
	// Connect to the SurrealDB instance using WebSocket
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		log.Fatal("Failed to connect to SurrealDB:", err)
	}

	// Sign in with the specified user credentials
	if _, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		log.Fatal("Failed to sign in:", err)
	}
	db.Use("root", "root")

	// Assign the authenticated DB client to the global DB variable
	DB = db
}
