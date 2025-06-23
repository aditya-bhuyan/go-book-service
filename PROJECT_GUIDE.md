# 📚 Go Book Service - Developer Reference

A fully-documented Go web service for managing books. Includes in-memory storage, REST API with validation, Swagger/OpenAPI docs, test coverage, CI/CD with GitHub Actions, and Docker deployment.

---

## 📁 Project File Overview

| File / Folder            | Purpose |
|--------------------------|---------|
| `main.go`                | Entry point for the application. Loads config and registers REST API routes. |
| `go.mod`, `go.sum`       | Go modules definition and dependencies. |
| `model/book.go`          | Defines the `Book` struct with JSON and validation tags. |
| `server/handlers.go`     | REST API handlers for adding, retrieving, and deleting books. |
| `client/client.go`       | CLI client to interact with the REST API via prompts. |
| `config/config.json`     | Configuration file for server and client (ports, server address). |
| `docs/`                  | Contains Swagger/OpenAPI files, Swagger UI assets (for GitHub Pages). |
| `Dockerfile`             | Multi-stage Docker build for the Go server. |
| `.dockerignore`          | Files and directories ignored by Docker build context. |
| `.github/workflows/`     | GitHub Actions CI for testing, coverage, and Codecov upload. |
| `LICENSE`                | GNU General Public License v3.0 for open source use. |
| `README.md`              | Primary project documentation and badges. |

---

## 🛠️ Local Setup & Testing

### 📦 1. Clone the Repository

```bash
git clone https://github.com/aditya-bhuyan/go-book-service.git
cd go-book-service
```

### ⚙️ 2. Install Dependencies

```bash
go mod tidy
```

### 🚀 3. Run the Server

```bash
go run main.go
```

### 🧪 4. Run Unit Tests with Coverage

```bash
go test ./server -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html  # (or xdg-open on Linux)
```

### 💻 5. Use the CLI Client

```bash
cd client
go run client.go
```

Follow the interactive prompts to Add / View / Delete books.

---

## 🐳 Docker Build, Run & Test

### 🔨 1. Build the Docker Image

```bash
docker build -t go-book-service .
```

### ▶️ 2. Run the Docker Container

```bash
docker run -p 8080:8080 go-book-service
```

### 🌐 3. Test API in Browser or Postman

- List Books: `GET http://localhost:8080/books`
- Add Book: `POST http://localhost:8080/books` with JSON body
- Delete Book: `DELETE http://localhost:8080/books/{id}`

### 🔍 4. Swagger UI (Optional)

If served via GitHub Pages:  
🔗 `https://aditya-bhuyan.github.io/go-book-service/`

Or run Swagger locally via `/swagger/*` endpoint if included.

---

## ✅ CI/CD & DockerHub

| Feature       | Status  |
|---------------|---------|
| ✅ Build       | ![Build](https://github.com/aditya-bhuyan/go-book-service/actions/workflows/go-test.yml/badge.svg)
| ✅ Coverage    | [![codecov](https://codecov.io/gh/aditya-bhuyan/go-book-service/branch/main/graph/badge.svg)](https://codecov.io/gh/aditya-bhuyan/go-book-service)
| 🐳 DockerHub   | [![Docker Pulls](https://img.shields.io/docker/pulls/your-dockerhub-username/go-book-service)](https://hub.docker.com/r/your-dockerhub-username/go-book-service)

---

## 👤 Author

**Aditya Pratap Bhuyan**  
🔗 [LinkedIn](https://linkedin.com/in/adityabhuyan)  
📦 [GitHub](https://github.com/aditya-bhuyan)

---

## 📝 License

This project is licensed under the [GNU GPL v3.0](LICENSE).