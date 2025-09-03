# 📇 Contact Management RESTful API

A simple RESTful API for managing users, contacts, and addresses, built with **Go**, **Fiber**, and **GORM**.  
This project is created **for learning purposes only**.

---

## 📌 Features

**User API:**
- Register a user
- Login & generate token
- Get current user
- Update current user
- Logout (invalidate token)

**Contact API** *(protected)*:
- Create a contact
- Search & list contacts (with pagination & filters)
- Get contact by ID
- Update contact
- Delete contact

**Address API** *(protected)*:
- Create address for a contact
- List all addresses for a contact
- Get address by ID
- Update address
- Delete address

**General Features:**
- Authentication with token (UUID) via Authorization header
- Input validation using `go-playground/validator`
- Consistent JSON response format (`data`, `errors`)
- Layered architecture *(Handler → Usecase → Repository → Entity)*
- OpenAPI 3.0 specification (YAML file included)

---

## 🛠️ Tech Stack

**Language:**
- Go 1.22+

**Framework & Libraries:**
- [Fiber](https://github.com/gofiber/fiber) – Web framework for routing
- [GORM](https://gorm.io/) – ORM for Go
- [PostgreSQL Driver for GORM](https://pkg.go.dev/gorm.io/driver/postgres)
- [Go Playground Validator](https://github.com/go-playground/validator) – Struct validation
- [Logrus](https://github.com/sirupsen/logrus) – Logging
- [Viper](https://github.com/spf13/viper) – Configuration management

**Database:**
- PostgreSQL 14+ (or compatible)

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/jonathangunawan30/golang-contact-management-restful-api.git
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Configure environment

Edit `config.yaml`:

```yaml
server:
  port: "3000"

database:
  url: "postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable"
```

### 4. Run database migration

Make sure PostgreSQL is running and the database exists.

Install `golang-migrate` CLI if you haven’t:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run migration:

```bash
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable" up
```

### 5. Run the application

```bash
go run main.go
```