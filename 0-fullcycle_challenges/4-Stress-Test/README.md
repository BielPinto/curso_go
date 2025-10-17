# Stress Test Tool

A CLI tool written in Go for performing HTTP stress tests on web services. It allows you to send a configurable number of concurrent requests to a specified URL and generates a detailed report on the performance and response status codes.

## Features

- Concurrent HTTP requests using goroutines
- Configurable number of total requests and concurrency level
- Detailed report including:
  - Total execution time
  - Number of successful requests (HTTP 200)
  - Distribution of other HTTP status codes
- Docker containerization for easy deployment

## Prerequisites

- Go 1.22 or later (for local development)
- Docker (for containerized execution)

## Installation

### Local Build

1. Clone or download the project
2. Navigate to the project directory
3. Build the application:

```bash
go build -o stress-test .
```

### Docker Build

1. Clone or download the project
2. Navigate to the project directory
3. Build the Docker image:

```bash
docker build -t stress-test .
```

## Usage

### Command Line Arguments

- `--url`: The URL of the service to test (required)
- `--requests`: Total number of requests to send (required, must be > 0)
- `--concurrency`: Number of concurrent requests (required, must be > 0)

### Running Locally

```bash
./stress-test --url=http://example.com --requests=1000 --concurrency=10
```

### Running with Docker

```bash
docker run --rm stress-test --url=http://example.com --requests=1000 --concurrency=10
```

### Example Output

```
Tempo total gasto na execução: 5.234567891s
Quantidade total de requests realizados: 1000
Quantidade de requests com status HTTP 200: 950
Distribuição de outros códigos de status HTTP:
  404: 30
  500: 20
```

## Report Details

The tool generates a report containing:

- **Total execution time**: Time taken to complete all requests
- **Total requests made**: Number of requests actually sent
- **Successful requests (200)**: Number of requests that returned HTTP 200
- **Status code distribution**: Breakdown of other HTTP status codes encountered

## Notes

- The tool uses HTTP GET requests only
- Errors during requests (network issues, timeouts) are recorded with status code 0
- The concurrency level controls how many requests are made simultaneously
- Ensure the target service can handle the load you're testing