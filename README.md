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

## Architecture and components used

  [Download the high level architecture document (PDF)](./architecture-fake-twitter.pdf)

---

## Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop) (includes Docker Compose)

    You must have an instance of docker running locally in your desktop before running the project.

- [Git](https://git-scm.com/)

---

## Getting Started

### 1. Clone the repository

```sh
git clone https://github.com/juaniferro/fake-twitter.git
cd fake-twitter
```

---

### 2. Environment Variables

Check if `.env` is already in project (basically you shouldn´t push these variables but for the sake of simpicity to try the API we will)

If it is not present, create a `.env` file in the project root:

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

In these cases, the user_id that we pass in the Headers parameter symbolizes the user making the action. They would be the user posting a tweet, following another user (the one in the path param) and getting their timeline

- **Get Timeline**
  ```
  GET /timeline
  Headers: user_id: <user_id>
  ```

  ```bash
    curl --location 'http://localhost:8080/timeline' \
    --header 'user_id: 4'
  ```

- **Post Tweet**
  ```
  POST /tweet
  Headers: user_id: <user_id>
  Body (JSON): { "content": "your tweet" }
  ```

  ```bash
    curl --location 'http://localhost:8080/tweet' \
    --header 'user_id: 1' \
    --header 'Content-Type: application/json' \
    --data '{
     "content" : "tuiteando para probar :)"
    }'
  ```

- **Follow User**
  ```
  POST /follow/{followed_user_id}
  Headers: user_id: <user_id>
  ```

  ```bash
    curl --location --request POST 'http://localhost:8080/follow/2' \
    --header 'user_id: 4'
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

- **Starting Docker issues:**  
  Sometimes, there could be a failure similar to this
  ```
  migrate-1  | error: failed to open database: dial tcp 172.18.0.2:3306: connect: connection refused  
  ```

  if that happens, run 
  ```sh
  docker-compose down -v
  ```
  and then try again
  ```sh
  docker-compose up --build
  ```
- **Port conflicts:**  
  If port 3306 is in use, stop local MySQL or change the port mapping in `docker-compose.yml`.
- **Reset database:**  
  Use `docker-compose down -v` to remove all data and start fresh (will replicate the starting data of the migration though).
- **.env issues:**  
  Ensure `.env` is a file (not a directory) and is in the project root.

---