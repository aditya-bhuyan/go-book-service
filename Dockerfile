# Stage 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copy binary and config
COPY --from=builder /app/app .
COPY config ./config
COPY docs ./docs

# Expose the app port
EXPOSE 8080

# Run the app
CMD ["./app"]
