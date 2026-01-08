FROM golang:1.23.3-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build binário estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/server/   
    
FROM gcr.io/distroless/base-debian12

# RUN go build -o search-weather cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates


WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./app"]