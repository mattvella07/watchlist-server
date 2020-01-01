package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mattvella07/watchlist-server/server/conn"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Login validates the username and password and returns a JWT
func Login(rw http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
	if !ok {
		log.Println("Username and/or password missing")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username and/or password missing"))
		return
	}

	existingUser := user{}
	query := `SELECT id, username, password
		FROM users
		WHERE username = $1
		LIMIT 1`

	// Check if user exists
	err := conn.DB.QueryRow(query, username).Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password)
	if err != nil {
		log.Printf("DB error: %s\n", err)
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Username and/or password"))
		return
	}

	// User exists, validate password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		log.Printf("Incorrect password: %s\n", err)
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Username and/or password"))
		return
	}

	// Generate JWT 

	// Set token as cookie

	log.Printf("User %s logged in\n", username)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Logged in"))
}
