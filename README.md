# API Helper

A simple and efficient HTTP client library for Go with built-in authentication support.

## Features

- Simple HTTP GET and POST requests
- Built-in authentication support (Bearer tokens and API keys)
- Configurable timeouts
- Clean error handling

## Installation

```bash
go get github.com/coles243/apihelper
```

## Usage as a Library

### Basic GET Request

```go
package main

import (
    "fmt"
    "log"

    "github.com/coles243/apihelper/client"
)

func main() {
    c := &client.ClientServer{
        URL:            "https://api.github.com/users/octocat",
        TimeoutsClient: 30, // 30 seconds
    }

    data, err := c.GetData()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(data))
}
```

### Authenticated GET Request

```go
c := &client.ClientServer{
    URL:            "https://api.example.com/protected",
    TimeoutsClient: 30,
    AuthToken:      "your_bearer_token",
    AuthType:       "Bearer", // or "x-api-key"
}

data, err := c.GetData()
// handle response...
```

### POST Request

```go
import "encoding/json"

payload := map[string]interface{}{
    "name":  "test",
    "value": 123,
}

jsonData, _ := json.Marshal(payload)

c := &client.ClientServer{
    URL:            "https://jsonplaceholder.typicode.com/posts",
    TimeoutsClient: 30,
    Body:           jsonData,
}

status, err := c.PostData()
// handle response...
```

## Examples

The `examples/` directory contains working examples:

- `get_example.go` - Simple GET request to GitHub API
- `post_example.go` - POST request to JSONPlaceholder API

To run examples:

```bash
cd examples
go run get_example.go
go run post_example.go
```

## Examples

The `examples/` directory contains working examples:

- `get_example.go` - Simple GET request to GitHub API
- `post_example.go` - POST request to JSONPlaceholder API

To run examples:

```bash
cd examples
go run get_example.go
go run post_example.go
```

## API Reference

### ClientServer

```go
type ClientServer struct {
    URL            string // API endpoint URL
    TimeoutsClient int    // Timeout in seconds
    Body           []byte // Request body for POST requests
    AuthToken      string // Authentication token
    AuthType       string // "Bearer" or "x-api-key"
}
```

### Methods

#### GetData() ([]byte, error)

Performs an HTTP GET request and returns the response body.

#### PostData() (string, error)

Performs an HTTP POST request and returns the HTTP status.

## Authentication Types

- **Bearer**: Adds `Authorization: Bearer <token>` header
- **x-api-key**: Adds `x-api-key: <token>` header

## Error Handling

The library provides detailed error messages for common scenarios:

- Network connectivity issues
- Non-200 HTTP status codes
- JSON parsing errors
- Authentication failures

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
