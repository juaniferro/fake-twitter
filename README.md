# Fake Twitter API

A simple Twitter-like REST API built with Go, MySQL, and Docker.

---

## Features

- Post tweets
- Follow users
- Get user timeline
- Table-driven unit tests for handlers, services, and repositories
- Database migrations with [migrate/migrate](https://github.com/golang-migrate/migrate)

---

## Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop) (includes Docker Compose)
- [Git](https://git-scm.com/)

---

## Getting Started

### 1. Clone the repository

```sh
git clone <your-repo-url>
cd fake-twitter
```

---

### 2. Environment Variables

Create a `.env` file in the project root:

```
DB_USER=root
DB_PASS=pass
DB_HOST=mysql
DB_PORT=3306
DB_NAME=fake_tw_database
SHOULD_PARSE_TIME=true
PORT=8080
```

---

### 3. Run with Docker Compose

Build and start all services (API, MySQL, migrations):

```sh
docker-compose up --build
```

- The API will be available at [http://localhost:8080](http://localhost:8080)
- MySQL runs in a container and is initialized with migrations and seed data.

To stop and remove all containers and data:

```sh
docker-compose down -v
```

---

### 4. API Usage

#### Example Endpoints

- **Get Timeline**
  ```
  GET /timeline
  Headers: user_id: <user_id>
  ```

- **Post Tweet**
  ```
  POST /tweet
  Headers: user_id: <user_id>
  Body (JSON): { "content": "your tweet" }
  ```

- **Follow User**
  ```
  POST /follow/{followed_user_id}
  Headers: user_id: <user_id>
  ```

Use [Postman](https://www.postman.com/) or `curl` to test endpoints.

---

### 5. Running Tests

To run unit tests (requires Go installed locally):

```sh
go test ./...
```

- Tests use table-driven style and mocks for handlers, services, and repositories.
- Example repository test uses [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock).

---

### 6. Database Migrations

- Migrations are located in the `migrations/` folder.
- They are automatically applied on `docker-compose up` using the `migrate/migrate` container.

---

### 7. Troubleshooting

- **Port conflicts:**  
  If port 3306 is in use, stop local MySQL or change the port mapping in `docker-compose.yml`.
- **Reset database:**  
  Use `docker-compose down -v` to remove all data and start fresh.
- **.env issues:**  
  Ensure `.env` is a file (not a directory) and is in the project root.

---