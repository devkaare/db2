package main

import (
	"log"

	"github.com/google/uuid"
    "github.com/devkaare/db2"
)

var userKey = "User"
var addressKey = "Address"

func main() {
	db.LoadCache()       // Loading the users.json files contents into memory and creating a users.json file if none exists
	defer db.SaveCache() // Ensuring that the program saves the user cache to the json file on shutdown
	// Important: The two functions above are CRUCIAL for this package to function!

	userId := uuid.New().String()
	newUser := CreateUser(userId, "emma_davis", "password", "emma@example.com")
	newAddress := CreateAddress(userId, "123 Main St", "New York", "NY", "10001")

	// Inserting a new user
	db.AddToCache(newUser)
	log.Println("Added user:", newUser)

	// Inserting a new address
	db.AddToCache(newAddress)
	log.Println("Added address:", newAddress)

	userIdToDelete := "e43cb71c-726d-4414-9433-a3cbfcfefb4d" // Important: Replace this with an UserId from db.json to test
	user := db.SearchCache(userKey, "UserId", userIdToDelete)

	if user != nil {
		log.Println("User found:", user)
		log.Println("Deleting user with ID:", userIdToDelete)
		db.DeleteFromCache(userKey, "UserId", userIdToDelete)
	} else {
		log.Println("User with ID", userIdToDelete, "does not exist.")
	}

	addressIdToDelete := "9a4a4582-2af0-465f-bded-ea32c4c30d61" // Important: Replace this with an UserId from db.json to test
	address := db.SearchCache(addressKey, "UserId", addressIdToDelete)

	if address != nil {
		log.Println("Address found:", address)
		log.Println("Deleting address with UserID:", addressIdToDelete)
		db.DeleteFromCache(addressKey, "UserId", addressIdToDelete)
	} else {
		log.Println("Address with UserID", addressIdToDelete, "does not exist.")
	}
}

func CreateUser(userId, username, password, email string) map[string]interface{} {
	return map[string]interface{}{
		userKey: map[string]interface{}{
			"UserId":   userId,
			"Username": username,
			"Email":    email,
			"Password": password,
		},
	}
}

func CreateAddress(userId, street, city, state, zip string) map[string]interface{} {
	return map[string]interface{}{
		addressKey: map[string]interface{}{
			"UserId":  userId,
			"Street":  "123 Main St",
			"City":    "New York",
			"State":   "NY",
			"ZipCode": "10001",
		},
	}
}
