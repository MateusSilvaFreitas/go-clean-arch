# 🧼 Go Clean Architecture - Foo API

This project is a sample API built with Go using the principles of **Clean Architecture**. It provides basic CRUD operations for the `Foo` entity (Create, Read, Delete).

---

## 🧩 Objective

Demonstrate how to apply **Clean Architecture** with Go in a real-world project.

---

## ⚙️ Technologies

- **Go (Golang)**
- **Gin** - Web framework
- **MySQL** - Relational database
- **Go-SQL-Driver** - MySQL driver for Go

---

## 🧱 Project Structure

```
.
├── internal
│   ├── entity        # Domain entities (pure business rules)
│   ├── usecase       # Application use cases
│   │   └── dto           # Input/output structures
│   ├── interface
│   │   └── controller  # Web adapters (HTTP handlers)
│   │   └── gateway     # Database interface
│   ├── infra
│   │   ├── db          # Repositories and database connection
│   │   ├── router      # Route definitions
│   │   └── dependencies # Dependency injection
│
├── main.go           # Entry point
```

---

## 🧼 Clean Architecture in this context

- **Entity**: Contains entities with validations and pure business logic (`Foo`).
- **Usecase**: Contains application logic (how to create, find, or delete a Foo).
- **Interface**: Implements adapters for use cases, such as HTTP handlers and Database interface.
- **Infra**: Hosts implementations of repositories (MySQL), HTTP router, and dependency setup.

---

## 📦 Implemented Use Cases

- `CreateFoo`: Creates a new Foo.
- `FindFooById`: Retrieves a Foo by ID.
- `DeleteFoo`: Deletes a Foo by ID.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- MySQL

### Database Setup

Create the `foo` database and table:

```sql
CREATE DATABASE foo;

USE foo;

CREATE TABLE foo (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  type VARCHAR(255),
  created_at DATETIME
);
```

You need to modify `mysql.go` to your database settings.

### Running the Project

```bash
go run cmd/go-clean-arch/main.go
```

The app will start on port `8080`.

---

## 📡 API Endpoints

| Method | Endpoint   | Description      |
| ------ | ---------- | ---------------- |
| POST   | `/foo/foo` | Create a new Foo |
| GET    | `/foo/:id` | Get Foo by ID    |
| DELETE | `/foo/:id` | Delete Foo by ID |

---

## 🧪 Testing

The following layers should be covered with unit tests:

- `entity`: entity validations (`Validate()` method).
- `usecase`: business logic (`CreateFoo`, `FindFooById`, `DeleteFoo`).
- `controller`: HTTP handlers (with mocked usecases).
- `db`: optional integration tests using Docker.

---

## 🧪 Unit Test

Unit Testing

🛠️🛠️🛠️🛠️🛠️🛠️

In progress

---
