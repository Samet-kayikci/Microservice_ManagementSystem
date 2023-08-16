package main

import (
	"fmt"
	"net/http"
)

func main() {
	baseURL := "http://localhost:8080" // Sunucu adresinizi ve port numarasını ayarlayın

	var UserID int
	fmt.Print("User to be deleted: ")
	fmt.Scan(&UserID)

	deleteURL := fmt.Sprintf("%s/delete?id=%d", baseURL, UserID)
	DeleteRequest, err := http.NewRequest("delete", deleteURL, nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(DeleteRequest)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("user successfully deleted")
	} else if response.StatusCode == http.StatusNotFound {
		fmt.Println("user not found")
	} else {
		fmt.Println("error: ", response.Status)
	}
}
