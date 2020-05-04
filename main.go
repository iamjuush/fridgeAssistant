package main

import (
	//"fmt"
	"fridgeAssistant/database"
	"fridgeAssistant/handlers"
	"fridgeAssistant/migrations"
	"log"
	"net/http"
)


func main() {
	database.InitDB() // Initialise the sqlite database if doesnt exist already.
	migrations.Migrate() // Migrate models inside models package into the database.
	defer database.DBCon.Close() // Close connection to database

	// Start the server and add the routing.
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/add/fridge", handlers.AddFridgeHandler)
	http.HandleFunc("/add/grocery", handlers.AddGroceryHandler)
	log.Fatal(http.ListenAndServe(":8778", nil))
}