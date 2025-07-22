package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coles243/apihelper/client"
)

func main() {
	payload := map[string]interface{}{
		"message": "Hello from API Helper!",
		"user":    "test_user",
		"data":    "This is a test POST with auth",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	c := &client.ClientServer{
		URL:            "https://httpbin.org/bearer", // This endpoint requires Bearer auth
		TimeoutsClient: 30,
		Body:           jsonData,
		AuthToken:      "your-test-token-123",
		AuthType:       "Bearer",
	}

	fmt.Println("Sending POST request with Bearer auth to httpbin...")
	status, err := c.PostData()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response Status: %s\n", status)
}
