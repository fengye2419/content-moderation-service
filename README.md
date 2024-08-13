# Content Moderation Service

This is a content moderation web service implemented in Go. It provides a RESTful API for validating keywords and checking if they match any sensitive words. The service returns the matched sensitive words and the count of matches.

## Project Structure

```
content-moderation-service
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   └── validate.go
│   ├── service
│   │   └── moderation.go
│   └── model
│       └── sensitive_words.go
├── pkg
│   └── middleware
│       └── logging.go
├── go.mod
└── README.md
```

## Files

### `cmd/main.go`

This file is the entry point of the application. It initializes the server and starts listening for incoming requests.

### `internal/api/validate.go`

This file defines the RESTful API endpoint `/validate` which receives a request with keywords and checks if they match any sensitive words. It returns the matched sensitive words and the count of matches.

### `internal/service/moderation.go`

This file contains the implementation of the content moderation service. It has a function `CheckSensitiveWords` that takes a list of keywords and checks if they match any sensitive words.

### `internal/model/sensitive_words.go`

This file defines the sensitive words model. It contains a list of sensitive words that will be used for matching against the keywords.

### `pkg/middleware/logging.go`

This file defines a middleware function for logging incoming requests and outgoing responses.

### `go.mod`

This file is the Go module file that specifies the project's dependencies.

## Usage

To use the content moderation service, you can send a POST request to the `/api/v1/sensitiveWords/validate` endpoint with a JSON payload containing the keywords to be validated. The service will respond with the matched sensitive words and the count of matches.

Example request:

```
POST /api/v1/sensitiveWords/validate
Content-Type: application/json

{
  "keyword": "keyword1"
}
```

Example response:

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "matched_words": ["keyword1", "keyword3"],
  "match_count": 2
}
```

## Dependencies

This project uses Go modules for dependency management. The required dependencies are listed in the `go.mod` file.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.