package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PayloadBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ClientPayload struct {
	ID   string       `json:"id"`
	Body *PayloadBody `json:"body"`
}

type Result struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func processTask(taskID string) {
	client := &http.Client{}

	_result := &Result{
		ID:     taskID,
		Status: "success",
	}

	result, err := json.Marshal(_result)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8001", bytes.NewBuffer(result))
	if err != nil {
		return
	}
	req.Header.Add("Type", "result")

	time.Sleep(time.Second * 6)

	_, err = client.Do(req)
	if err != nil {
		return
	}

	fmt.Println("Task completed")
}

func main() {
	for {
		fmt.Println("Waiting for task...")

		// Get client payload
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8001/users", nil)
		if err != nil {
			return
		}
		req.Header.Add("Type", "get")

		res, err := client.Do(req)
		if err != nil {
			return
		}
		defer res.Body.Close()

		var clientPayload *ClientPayload
		err = json.NewDecoder(res.Body).Decode(&clientPayload)
		if err != nil {
			return
		}

		clientPayloadJson, err := json.MarshalIndent(clientPayload, "", "  ")
		if err != nil {
			return
		}
		fmt.Printf("Client payload: %s\n", string(clientPayloadJson))

		// Send result to client
		if clientPayload != nil {
			go processTask(clientPayload.ID)
		}
	}
}
