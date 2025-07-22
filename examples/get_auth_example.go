package main

import (
	"fmt"

	"github.com/coles243/apihelper/client"
)

func main() {
	c := &client.ClientServer{
		URL:            "https://httpbin.org/bearer",
		TimeoutsClient: 10,
		AuthToken:      "JustABunchoftext",
		AuthType:       "Bearer",
	}

	data, err := c.GetData()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
