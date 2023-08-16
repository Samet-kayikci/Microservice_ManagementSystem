package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var UserName string
	fmt.Print("username: ")
	fmt.Scan(&UserName)
	var mail string
	fmt.Print("Email: ")
	fmt.Scan(&mail)

	user := User{
		Username: UserName,
		Email:    mail,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("User created successfully!")
	} else {
		fmt.Println("Error creating user. Status:", resp.Status)
	}
}
