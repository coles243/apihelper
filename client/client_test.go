package client

import (
	"bytes"
	"encoding/json"
	"testing"
)

type Comments struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func TestGetRequest(t *testing.T) {
	post := Comments{}

	c := ClientServer{
		URL:            "https://jsonplaceholder.typicode.com/posts/1",
		TimeoutsClient: 10,
	}

	data, _ := c.GetData()

	json.NewDecoder(bytes.NewReader(data)).Decode(&post)

	actual := post.UserId
	expect := 1

	if actual != expect {
		t.Fatal("Did not get expected result")

	}

}
