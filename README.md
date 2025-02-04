# My Go API

This project is a basic API built with Go. It provides endpoints to manage items, allowing users to retrieve and create items.

## Project Structure

```
my-go-api
├── cmd
│   └── main.go          # Entry point of the application
├── pkg
│   ├── handlers
│   │   └── handlers.go  # HTTP request handlers
│   ├── models
│   │   └── models.go    # Data structures
│   └── routes
│       └── routes.go    # API routes setup
├── go.mod                # Module definition
└── README.md             # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/my-go-api.git
   ```
2. Navigate to the project directory:
   ```
   cd my-go-api
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```

### Running the API

To run the API, execute the following command:
```
go run cmd/main.go
```

The API will start on `http://localhost:8080`.

### API Endpoints

- `GET /items` - Retrieve all items
- `POST /items` - Create a new item

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.