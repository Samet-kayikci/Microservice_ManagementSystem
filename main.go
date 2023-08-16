package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = make(map[int]User)
var lastID = 0

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastID++
	user.ID = lastID
	users[user.ID] = user

	w.WriteHeader(http.StatusCreated)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userID := parseInt(id)
	user, ok := users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userID := parseInt(id)
	_, ok := users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, userID)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/create", createUserHandler)
	http.HandleFunc("/get", getUserHandler)
	http.HandleFunc("/delete", deleteUserHandler)

	fmt.Println("Microservice listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
