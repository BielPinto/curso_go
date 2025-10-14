# Rate Limiter

This is a rate limiter implementation in Go that limits the number of requests per second based on IP address or API token.

## Features

- Rate limiting by IP address
- Rate limiting by API token (overrides IP limits)
- Configurable limits and block times
- Redis-backed storage
- HTTP middleware for easy integration
- Docker support

## Configuration

Configuration is done via environment variables or a `.env` file in the root directory.

### Environment Variables

- `REDIS_ADDR`: Redis server address (default: localhost:6379)
- `REDIS_PASSWORD`: Redis password (default: "")
- `REDIS_DB`: Redis database number (default: 0)
- `RATE_LIMIT_IP`: Max requests per second per IP (default: 10)
- `RATE_LIMIT_TOKEN`: Max requests per second per token (default: 100)
- `BLOCK_TIME_IP`: Block time in seconds for IP when limit exceeded (default: 300)
- `BLOCK_TIME_TOKEN`: Block time in seconds for token when limit exceeded (default: 300)
- `PORT`: Server port (default: 8080)

## Running

1. Start Redis:
   ```bash
   docker-compose up -d
   ```

2. Run the server:
   ```bash
   go run cmd/main.go
   ```

The server will start on port 8080.

## Usage

### As Middleware

The rate limiter is implemented as HTTP middleware. It checks the `API_KEY` header for token-based limiting, otherwise falls back to IP-based limiting.

### API

- GET / : Test endpoint that returns "Request allowed" if not rate limited.

### Response

When rate limited, returns HTTP 429 with message: "you have reached the maximum number of requests or actions allowed within a certain time frame"

## Testing

Run tests:
```bash
go test ./...
```

## Architecture

- `internal/config`: Configuration loading
- `internal/storage`: Storage interface and Redis implementation
- `internal/limiter`: Rate limiting logic
- `internal/middleware`: HTTP middleware
- `cmd/main.go`: Main application
