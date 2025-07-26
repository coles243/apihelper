package main

import (
	"fmt"

	"github.com/coles243/apihelper/client"
)

func main() {

	poster()

}

func poster() {

	c := &client.ClientServer{
		URL:            "https://httpbin.org/anything",
		TimeoutsClient: 10,
	}

	data, err := c.PostData()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
