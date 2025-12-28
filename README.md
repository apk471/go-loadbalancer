# Go Load Balancer

A simple HTTP Load Balancer implemented in Go (Golang) using the Round Robin algorithm. This project demonstrates how to create a reverse proxy and distribute incoming traffic across multiple backend servers.

## Features

- **Round Robin Balancing**: Distributes requests sequentially across the available servers.
- **Reverse Proxy**: Uses `net/http/httputil` to forward requests to backend servers.
- **Server Health Check Interface**: Includes an interface for server health checks (currently stubbed to always return true).

## Project Structure

- `main.go`: Contains the core logic for the load balancer, server interface, and the main entry point.
- `makefile`: Helper commands for building and running the project.
- `go.mod`: Go module definition.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later recommended)

### Build and Run

The project includes a `makefile` for easy management.

1.  **Run the Load Balancer**:

    ```bash
    make run
    ```

    This will start the load balancer on port `8080`.

2.  **Build the Binary**:

    ```bash
    make build
    ```

    This compiles the application and places the binary in `bin/loadbalancer`.

3.  **Clean Build Artifacts**:
    ```bash
    make clean
    ```

## Usage

Once running, the load balancer listens on `http://localhost:8080`.

By default, it is configured to forward requests to the following external services for demonstration:

- https://google.com
- https://github.com
- https://youtube.com
- https://facebook.com
- https://twitter.com

You can test it by opening a browser or using `curl`:

```bash
curl http://localhost:8080
```

Repeated requests will cycle through the configured servers.
