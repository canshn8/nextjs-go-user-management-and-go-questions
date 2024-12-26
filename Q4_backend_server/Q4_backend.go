package main

import (
	"database/sql"   // Package for interacting with databases
	"encoding/json"  // Package for encoding/decoding JSON
	"log"             // Package for logging errors
	"net/http"        // Package for handling HTTP requests
	"strconv"         // Package for converting strings to integers

	_ "github.com/mattn/go-sqlite3" // Import SQLite3 driver
)

type User struct {
	ID   int    `json:"id"`   // User ID (Primary Key)
	Name string `json:"name"` // User name
	Age  int    `json:"age"`  // User age
}

var db *sql.DB // Global variable for the database connection

// Function to initialize the SQLite database and create the users table if it doesn't exist
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db") // Open the SQLite database file
	if err != nil {
		log.Fatal(err) // Log an error and stop if database connection fails
	}

	// Create the users table if it doesn't already exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)`)
	if err != nil {
		log.Fatal(err) // Log an error if the table creation fails
	}
}

// Handler to retrieve all users from the database and return them as JSON
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all users
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an error if the query fails
		return
	}
	defer rows.Close() // Close the rows when done

	var users []User // Slice to store the list of users

	// Loop through the result set and scan each user into the users slice
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Handle scan errors
			return
		}
		users = append(users, user) // Add the user to the slice
	}

	// Set the response header to application/json and encode the users slice to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler to get a user by ID and return the user as JSON
func getUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest) // Return error if ID is not valid
		return
	}

	// Query the database to get the user with the given ID
	row := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound) // Return error if user is not found
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Handle other errors
		}
		return
	}

	// Set the response header and encode the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler to create a new user in the database
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode the request body into a user object
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // Return error if the body is not valid JSON
		return
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return error if insertion fails
		return
	}

	// Respond with the created user object
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler to update an existing user's information
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest) // Return error if ID is not valid
		return
	}

	var user User

	// Decode the request body into a user object
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // Return error if the body is not valid JSON
		return
	}

	// Update the user's details in the database
	_, err = db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return error if update fails
		return
	}

	// Respond with the updated user object
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler to delete a user from the database
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest) // Return error if ID is not valid
		return
	}

	// Delete the user from the database
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return error if deletion fails
		return
	}

	// Respond with HTTP 204 (No Content) to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB() // Initialize the database

	// Define HTTP routes and associate them with the respective handlers
	http.HandleFunc("/users", getAllUsers)         // Route to get all users
	http.HandleFunc("/users/", getUserByID)        // Route to get a user by ID
	http.HandleFunc("/users", createUser)          // Route to create a new user
	http.HandleFunc("/users/", updateUser)         // Route to update an existing user
	http.HandleFunc("/users/", deleteUser)         // Route to delete a user

	log.Fatal(http.ListenAndServe(":8080", nil))
}
