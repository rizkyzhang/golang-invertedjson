package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type PayloadBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Payload struct {
	ID   string       `json:"id"`
	Body *PayloadBody `json:"body"`
}

type Result struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	_payload := &Payload{
		ID: uuid.NewString(),
		Body: &PayloadBody{
			Username: "test",
			Email:    "test@email.com",
		},
	}
	payload, err := json.Marshal(_payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/users", bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	fmt.Println("Waiting for result...")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var resBody Result
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return
	}

	fmt.Printf("Result: %+v", resBody)
}
