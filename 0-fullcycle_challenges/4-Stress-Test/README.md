# Stress Test Tool

A CLI tool written in Go for performing HTTP stress tests on web services. It allows you to send a configurable number of concurrent requests to a specified URL and generates a detailed report on the performance and response status codes.

## Features

- Concurrent HTTP requests using goroutines
- Configurable number of total requests and concurrency level
- Detailed report including:
  - Total execution time
  - Number of successful requests (HTTP 200)
  - Distribution of other HTTP status codes
- Performance profiling with pprof (CPU and memory profiles)
- Report saved to file (report.txt)
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
Relatório salvo em report.txt
Perfis salvos: cpu.prof e mem.prof
```

The report is saved to `report.txt` with the following content:

```
Tempo total gasto na execução: 5.234567891s
Quantidade total de requests realizados: 1000
Quantidade de requests com status HTTP 200: 950
Distribuição de outros códigos de status HTTP:
  404: 30
  500: 20
```

## Report Details

The tool generates a report saved to `report.txt` containing:

- **Total execution time**: Time taken to complete all requests
- **Total requests made**: Number of requests actually sent
- **Successful requests (200)**: Number of requests that returned HTTP 200
- **Status code distribution**: Breakdown of other HTTP status codes encountered

## Performance Profiling

The tool includes built-in performance profiling using Go's pprof package:

- **CPU Profile**: Captures CPU usage during the stress test and saves it to `cpu.prof`
- **Memory Profile**: Captures memory allocation and saves it to `mem.prof`

### Analyzing Profiles

To analyze the generated profiles, use the `go tool pprof` command:

```bash
# Analyze CPU profile
go tool pprof -text cpu.prof

# Analyze memory profile
go tool pprof -text mem.prof

# Interactive analysis (web interface)
go tool pprof -web cpu.prof
go tool pprof -web mem.prof
```

The profiles help identify performance bottlenecks, memory leaks, and optimization opportunities in the stress test execution.

## Notes

- The tool uses HTTP GET requests only
- Errors during requests (network issues, timeouts) are recorded with status code 0
- The concurrency level controls how many requests are made simultaneously
- Ensure the target service can handle the load you're testing