package main

import (
	"db-beta/db" // Replace this with the URL for this repo when testing!
	"log"
	//"time"

	"github.com/google/uuid"
)

var userKey = "Users"

func main() {
	db.LoadCache("users.json") // Loading the users.json files contents into memory and creating a users.json file if none exists
	defer db.SaveCache() // Ensuring that the program saves the user cache to the json file on shutdown, This can also be done after every save operation aswell

	userId := uuid.New().String()
	newUser := CreateUser(userId, "Emma A. Davis", "password2", "emma@example.com2")

	// Inserting a new user
	db.AddToCache(userKey, newUser)
	log.Println("Added user:", newUser)

	userIdToDelete := userId // Important: Replace this with an UserId from db.json to test
	user := db.SearchCache(userKey, "UserId", userIdToDelete)

	if user != nil {
		log.Println("User found:", user)
        // Get user Id using the key "UserId"
		log.Println("User Id:", user["UserId"])
        // Delete from cache
		db.DeleteFromCache(userKey, "UserId", userIdToDelete)
	} else {
		log.Println("User with Id", userIdToDelete, "does not exist.")
	}
    
    // Get full user cache
    cache := db.GetCache(userKey)
    for _, users := range cache {
        log.Println("User:", users)
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
