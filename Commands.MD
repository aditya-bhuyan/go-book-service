# Initial Commands

go mod init go-book-service

go get github.com/gin-gonic/gin


# Swagger Commands

go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/swaggo/gin-swagger

go get -u github.com/swaggo/files

swag init [To generate documentation]


# Adding Validations and Unit Test Cases

go get github.com/go-playground/validator/v10

go get github.com/stretchr/testify/assert

go test ./server -v [To test the application with server running]

# Coverage Details

## Run test with coverage output

go test ./server -coverprofile=coverage.out -covermode=atomic

## Generate HTML report

go tool cover -html=coverage.out -o coverage.html

## Open it in browser

open coverage.html    # On macOS

## OR

xdg-open coverage.html  # On Linux

# Docker Commands

Build - docker build -t go-book-service .

Run - docker run -d -p 8080:8080 --name go-books go-book-service

Verify - http://localhost:8080/books