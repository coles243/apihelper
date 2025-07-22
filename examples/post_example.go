// Example: POST request with JSON data
// Run from examples directory with: go run post_example.go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coles243/apihelper/client"
)

func main() {
	// Example: POST request with JSON data
	payload := map[string]interface{}{
		"title":  "Sample Post",
		"body":   "This is a test post",
		"userId": 1,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	c := &client.ClientServer{
		URL:            "https://jsonplaceholder.typicode.com/posts",
		TimeoutsClient: 30,
		Body:           jsonData,
	}

	fmt.Println("Sending POST request...")
	status, err := c.PostData()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response Status: %s\n", status)
}
