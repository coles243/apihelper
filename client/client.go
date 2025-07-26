package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client init
type ClientServer struct {
	URL            string
	TimeoutsClient int
	Body           []byte
	AuthToken      string
	AuthType       string
}

// Get Request
func (c *ClientServer) GetData() ([]byte, error) {

	ctx := context.Background()

	server := &http.Client{
		Timeout: time.Duration(c.TimeoutsClient) * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL, nil)

	if c.AuthToken != "" {

		switch c.AuthType {
		case "Bearer":
			req.Header.Set("Authorization", "Bearer "+c.AuthToken)

		case "x-api-key":
			req.Header.Set("x-api-key", c.AuthToken)

		default:
			return nil, errors.New("unable to resolve header type")

		}

	}

	if err != nil {
		msg := fmt.Sprintf("Error Processing inital request: %v", err)
		return nil, errors.New(msg)
	}

	response, err := server.Do(req)

	if err != nil {
		msg := fmt.Sprintf("Error Processing inital request: %v", err)
		return nil, errors.New(msg)
	}

	if response.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Error, Expected 200 status but returned status was: %s ", response.Status)
		return nil, errors.New(msg)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		msg := fmt.Sprintf("unable to parse data:  %v", err)
		return nil, errors.New(msg)
	}

	return data, nil

}

// PostRequest
func (c *ClientServer) PostData() (string, error) {
	ctx := context.Background()

	Server := &http.Client{
		Timeout: time.Duration(c.TimeoutsClient) * time.Second,
	}

	var bodyReader io.Reader
	if len(c.Body) == 0 {
		bodyReader = nil
	} else {
		bodyReader = bytes.NewReader(c.Body)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bodyReader)

	// Set Content-Type for JSON data
	if len(c.Body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.AuthToken != "" {

		switch c.AuthType {
		case "Bearer":
			req.Header.Set("Authorization", "Bearer "+c.AuthToken)

		case "x-api-key":
			req.Header.Set("x-api-key", c.AuthToken)

		default:
			return "", errors.New("unable to resolve header type")

		}
	}

	if err != nil {
		msg := fmt.Sprintf("Unable to process requests, Error: %v", err)
		return "", errors.New(msg)
	}

	resp, err := Server.Do(req)

	if err != nil {
		msg := fmt.Sprintf("Unable to process requests, Error: %v", err)
		return "", errors.New(msg)
	}

	defer resp.Body.Close()

	// Read the response body properly
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("unable to read response body: %v", err)
		return "", errors.New(msg)
	}

	return string(respData), nil
}
