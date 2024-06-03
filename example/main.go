package main

import (
	"github.com/devkaare/db2"
	"log"
	//"time"

	"github.com/google/uuid"
)

var userKey = "Users"

func main() {
	db.LoadCache("users.json") // Loading the users.json files contents into memory and creating a users.json file if none exists
	defer db.SaveCache() // Ensuring that the program saves the user cache to the json file on shutdown
	// Important: The two functions above are CRUCIAL for this package to function!

	userId := uuid.New().String()
	newUser := CreateUser(userId, "Emma A. Davis", "password2", "emma@example.com2")

	// Inserting a new user
	db.AddToCache(userKey, newUser)
	log.Println("Added user:", newUser)

	userIdToDelete := "ee3bebf0-421d-4a9f-9a01-fd12ccb8805f" // Important: Replace this with an UserId from db.json to test
	user := db.SearchCache(userKey, "UserId", userIdToDelete)

	if user != nil {
		log.Println("User found:", user)
		log.Println("User Id:", user["UserId"])
		db.DeleteFromCache(userKey, "UserId", userIdToDelete)
	} else {
		log.Println("User with Id", userIdToDelete, "does not exist.")
	}
}

func CreateUser(userId, username, password, email string) map[string]interface{} {
	return map[string]interface{}{
		"UserId":   userId,
		"Username": username,
		"Email":    email,
		"Password": password,
	}
}
