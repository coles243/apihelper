// Example: Simple GET request
// Run from examples directory with: go run get_example.go
package main

import (
	"fmt"
	"log"

	"github.com/coles243/apihelper/client"
)

func main() {
	// Example: Get GitHub user information
	c := &client.ClientServer{
		URL:            "https://api.github.com/users/octocat",
		TimeoutsClient: 30,
	}

	fmt.Println("Fetching GitHub user data...")
	data, err := c.GetData()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(string(data))
}
