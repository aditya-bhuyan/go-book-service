
# 📚 Go Book Service

A lightweight, containerized RESTful web service to manage books, written in Go using the Gin framework. This project includes Swagger/OpenAPI documentation, input validation, GitHub Actions CI with test coverage, Docker support, and a client CLI tool to interact with the API.

---

## ✨ Features

- 🧠 REST API for managing books
- 📦 In-memory database (map-based)
- ✅ Input validation using `go-playground/validator`
- 🧪 Unit tests with coverage (`testify`, `go test`)
- 🐳 Dockerized with multi-stage build
- 🚀 GitHub Actions CI pipeline
- 🌐 Swagger UI via GitHub Pages

---

## 📂 Project Structure

```

go-book-service/
├── client/                 # CLI client to interact with server
├── config/                 # Configuration files (JSON)
├── docs/                   # Swagger docs + GitHub Pages
├── model/                  # Book struct
├── server/                 # Handlers with validation
├── .github/workflows/      # CI config
├── Dockerfile              # Docker image build file
├── go.mod / go.sum         # Go module dependencies
└── main.go                 # Entry point

````

---

## 🚀 Run the Project

### 🔧 1. Build and Run Locally

```bash
go run main.go
````

### 🐳 2. Run via Docker

```bash
docker build -t go-book-service .
docker run -p 8080:8080 go-book-service
```

### ⚙️ 3. Run Tests with Coverage

```bash
go test ./server -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

---

## 🧪 API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/books`      | List all books    |
| POST   | `/books`      | Add a new book    |
| DELETE | `/books/{id}` | Delete book by ID |

---

## 🌐 Swagger API Docs

* Live docs: [GitHub Pages](https://aditya-bhuyan.github.io/go-book-service/)
* Local: run server and visit `http://localhost:8080/swagger/index.html`

---

## 🧑‍💻 CLI Client Tool

```bash
cd client/
go run client.go
```

Interactively choose options to add, view, or delete books.

---

## 🛠 Technologies Used

* [Go (Gin)](https://github.com/gin-gonic/gin)
* [Docker](https://www.docker.com/)
* [Swagger](https://swagger.io/tools/swagger-ui/)
* [Codecov](https://codecov.io/)
* [GitHub Actions](https://docs.github.com/en/actions)

---

## 📈 CI/CD & Coverage

| Feature    | Status                                                                                                                                             |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| ✅ Build    | ![Build](https://github.com/aditya-bhuyan/go-book-service/actions/workflows/go-test.yml/badge.svg)                                                 |
| ✅ Coverage | [![codecov](https://codecov.io/gh/aditya-bhuyan/go-book-service/branch/main/graph/badge.svg)](https://codecov.io/gh/aditya-bhuyan/go-book-service) |

---
## Docker 
| Feature       | Status  |
|---------------|---------|
| ✅ Build       | ![Build](https://github.com/aditya-bhuyan/go-book-service/actions/workflows/go-test.yml/badge.svg)
| ✅ Coverage    | [![codecov](https://codecov.io/gh/aditya-bhuyan/go-book-service/branch/main/graph/badge.svg)](https://codecov.io/gh/aditya-bhuyan/go-book-service)
| 🐳 DockerHub   | [![Docker Pulls](https://img.shields.io/docker/pulls/adityapratapbhuyan/go-book-service)](https://hub.docker.com/r/adityapratapbhuyan/go-book-service)

---

## 👤 Author

**Aditya Pratap Bhuyan**
🔗 [LinkedIn](https://linkedin.com/in/adityabhuyan)
📦 [GitHub](https://github.com/aditya-bhuyan)

---

## 📝 License

This project is licensed under the [GNU General Public License v3.0](LICENSE).


---

## 🙌 Contributions

Contributions, issues, and feature requests are welcome!
Please open a PR or create an issue.



---
