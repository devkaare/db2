<h1>Not for production use and not recommended for large datasets!</h1>
<div>
  <p>This <code>db2</code> package provides basic functions for managing key-value pairs stored in a JSON file.</p>
</div>
<div>
  <p><strong>Important:</strong><br>
    Call <code>LoadCache(dbFilePath string)</code> at the start of your main function to initialize the cache with the JSON file located at <code>dbFilePath</code>, and <code>defer SaveCache()</code> to save changes when the program exits. Additionally, a background process automatically saves the cache every 5 minutes.</p>
</div>
<div>
  <h3>Functions:</h3>
  <p><strong>1. LoadCache(dbFilePath string):</strong><br>
    Loads the cache from the specified JSON file. Call this at the start of your main function.</p>
  <p><strong>2. SaveCache():</strong><br>
    Saves the cache to the JSON file. Defer this at the start of your main function.</p>
  <p><strong>3. AddToCache(key string, value map[string]interface{}):</strong><br>
    Adds a value to the cache under the specified key.</p>
  <p><strong>4. SearchCache[T any](key string, field string, search T) map[string]interface{}:</strong><br>
    Searches the cache for an item matching the specified key and field. Returns the first matching item.</p>
  <p><strong>5. DeleteFromCache[T any](key string, field string, search T):</strong><br>
    Deletes an item from the cache based on the specified key and field.</p>
</div>
<div>
  <h3>Example Usage:</h3>
  <pre><code>
package main
import (
"github.com/yourusername/db2"
"log"

arduino
Kopier kode
"github.com/google/uuid"
)

var userKey = "Users"

func main() {
db.LoadCache("users.json") // Loading the users.json file's contents into memory and creating a users.json file if none exists
defer db.SaveCache() // Ensuring that the program saves the user cache to the json file on shutdown
// Important: The two functions above are CRUCIAL for this package to function!

go
Kopier kode
userId := uuid.New().String()
newUser := CreateUser(userId, "Emma A. Davis", "password2", "emma@example.com2")

// Inserting a new user
db.AddToCache(userKey, newUser)
log.Println("Added user:", newUser)

userIdToDelete := "ee3bebf0-421d-4a9f-9a01-fd12ccb8805f" // Important: Replace this with a UserId from users.json to test
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
"UserId": userId,
"Username": username,
"Email": email,
"Password": password,
}
}
</code></pre>

</div>
