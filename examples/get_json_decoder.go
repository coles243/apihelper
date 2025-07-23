package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/coles243/apihelper/client"
)

type Comments struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	post := Comments{}

	c := client.ClientServer{
		URL:            "https://jsonplaceholder.typicode.com/posts/1",
		TimeoutsClient: 10,
	}
	data, _ := c.GetData()

	json.NewDecoder(bytes.NewReader(data)).Decode(&post)

	fmt.Println(post.Body)
}
