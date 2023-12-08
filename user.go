// user.go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// User represents a user in the system
type User struct {
	ID         string     `json:"id"`
	SecretCode string     `json:"secretCode"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Playlists  []Playlist `json:"playlists"`
}

// Mutex for concurrency safety
var mu sync.Mutex

// Users map to store user data
var users = make(map[string]User)
// user.go



// findUserBySecretCode searches for a user by secret code
func findUserBySecretCode(secretCode string) (User, bool) {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.SecretCode == secretCode {
			return user, true
		}
	}
	return User{}, false
}


// generateUniqueID generates a unique ID for users and playlists
func generateUniqueID() string {
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		panic(err) // Handle error appropriately in a production environment
	}
	return hex.EncodeToString(id)
}

// generateUniqueSecretCode generates a unique secret code for users
func generateUniqueSecretCode() string {
	secretCode := make([]byte, 8)
	_, err := rand.Read(secretCode)
	if err != nil {
		panic(err) // Handle error appropriately in a production environment
	}
	return hex.EncodeToString(secretCode)
}

func login(w http.ResponseWriter, r *http.Request) {
	// Extract secret code from request
	secretCode := r.FormValue("secretCode")
	fmt.Println(secretCode)
	// Check if the user with the provided secret code exists

	user, found := findUserBySecretCode(secretCode)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	fmt.Println(user)
	// Return user details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func register(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body to get user details
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate unique ID and Secret Code for the new user
	newUser.ID = generateUniqueID()
	newUser.SecretCode = generateUniqueSecretCode()

	// Add the new user to the users map
	mu.Lock()
	defer mu.Unlock()
	users[newUser.ID] = newUser

	// Return the details of the newly created user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func viewProfile(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request
	userID := r.FormValue("userID")

	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
