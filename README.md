## Description

Here is a simple API I created based on your specifications. I just followed the description given in the README. I have also included some curl commands that I used for testing. You'll also notice that I have simplified the IDs to just be an incrementing integer, I used a UUID at first but it's just easier to test using an integer. You can revert that change if you'd like. I have also left in my comments so you can see my thought process. 

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/Kcin4/fetch-api.git
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the API

To run the API, execute the following command:
```
go run cmd/main.go
```

The API will start on `http://localhost:8080`.

Just send your requests through `/http://localhost:8080/receipts/process` and `/http://localhost:8080/receipts/{id}/points` by replacing `{id}` with the ID you'd like to get.