# Content Sharing Service

### A lightweight content-sharing service that allows users to create, store, and share content using in-memory storage. This project serves as a simple example of how to run a service that utilizes Go's native features.

## Prerequisites

- **Go**: Make sure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).
- **Ports**: The service is configured to run on `localhost:3000` by default, but this can be changed as needed.

## How to Run

### 1. Configuration

- Open the `main.go` file.
- Inside `main.go`, locate the section where the server URL and port are configured. By default, the service runs on `localhost:3000`.

- Modify the host or port if you need to run the server on a different address.

### 2. Starting the Server

- Open your terminal and navigate to the project directory where `main.go` is located.

- Use the following command to start the server:

  ```bash
  go run main.go
  ```
