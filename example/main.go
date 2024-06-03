package main

import (
	"db-beta/db"
	"log"
	"time"

	"github.com/google/uuid"
)

var userKey = "User"

func main() {
	db.LoadCache()       // Loading the users.json files contents into memory and creating a users.json file if none exists
	defer db.SaveCache() // Ensuring that the program saves the user cache to the json file on shutdown
	// Important: The two functions above are CRUCIAL for this package to function!

	userId := uuid.New().String()
	newUser := CreateUser(userId, "emma_davis2", "password2", "emma@example.com2")

	// Inserting a new user
	db.AddToCache(userKey, newUser)
	log.Println("Added user:", newUser)

	userIdToDelete := "1205f43c-a04a-42fe-890e-83399c587e69" // Important: Replace this with an UserId from db.json to test
	user := db.SearchCache(userKey, "UserId", userIdToDelete)
    //log.Println("User:", user[0]["UserId"])

	if user != nil {
		log.Println("User found:", user)
		log.Println("Deleting user with ID:", userIdToDelete)
		db.DeleteFromCache(userKey, "UserId", userIdToDelete)
	} else {
		log.Println("User with ID", userIdToDelete, "does not exist.")
	}

    time.Sleep(10 * time.Minute)
}

func CreateUser(userId, username, password, email string) map[string]interface{} {
	return map[string]interface{}{
			"UserId":   userId,
			"Username": username,
			"Email":    email,
			"Password": password,
	}
}
