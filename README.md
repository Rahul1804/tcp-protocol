# TCP Protocol Project

This project demonstrates how to create a custom protocol based on TCP using Go. The protocol involves a client sending a string "HELLO" and the server responding with "WORLD". Any other message results in a response of "UNKNOWN COMMAND".

## Getting Started

### Prerequisites

- Go (https://golang.org/dl/)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/rahul1804/tcp-protocol.git
    cd tcp-protocol
    ```

2. Initialize the Go module:
    ```sh
    go mod tidy
    ```

### Running the Server

1. Navigate to the `server` directory:
    ```sh
    cd server
    ```

2. Run the server:
    ```sh
    go run main.go
    ```

   The server will start and listen on port 8080.

### Running the Client

1. Open another terminal and navigate to the `client` directory:
    ```sh
    cd client
    ```

2. Run the client:
    ```sh
    go run main.go
    ```

3. Enter messages when prompted. The client will send the message to the server and display the server's response.

### Protocol Description

- **Message:** `HELLO`
  - **Response:** `WORLD`

- **Message:** Any other string
  - **Response:** `UNKNOWN COMMAND`

### Testing

To run tests for the protocol logic, navigate to the root directory of the project and run:
```sh
go test ./protocol
