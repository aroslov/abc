package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func sendRegisterUser() {
	// Register user (POST http://167.114.247.67:8080/api/users)

	json := []byte(`{"email": "roslov.anton@gmail.com"}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://167.114.247.67:8080/api/users", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	sendRegisterUser()
}
