# Stage 1: Build the Go binary
FROM golang:1.23.4 AS builder

WORKDIR /go-book-service

# Enable Go modules explicitly
ENV GO111MODULE=on

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Install swag CLI to generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy source code
COPY . .

# Generate Swagger documentation
RUN /go/bin/swag init

# Build the binary from module path
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the app binary and runtime resources
COPY --from=builder /app/app .
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs

# Expose service port
EXPOSE 8080

# Run the binary
CMD ["./app"]
