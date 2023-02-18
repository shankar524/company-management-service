[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
# Go Company Management Service
> a simple company management micro-service written in Golang

## Features
- Create/Read/Update/Delete company using API
- Authenticated routes/requests(using JWT)
- Postgres database for storage
- Publish Event (Kafka WIP)

### Stack
- [Postgres](https://www.postgresql.org) for a database
- Golang's [`database/sql`](https://golang.org/pkg/database/sql/) and this [Postgres driver](https://github.com/lib/pq) for database interactions
- [Gin](https://github.com/gin-gonic/gin) for a http framework / routing
- [Docker](https://www.docker.com) for containerization
- Authentication using [JWT tokens](https://jwt.io)


## Installation

### Using the Makefile

#### Available commands

- `make test` Runs all (unit and integraion)tests using docker compose
- `make lint` Lints the whole project using *gometalinter* (which is automatically installed if not already)
- `make coverage` Generates a HTML code coverage report [coverage_report.html](./coverage_report.html)
- `make run-all-local` runs whole system docker compose

## Running locally

- Download the necessary dependencies with [dep](https://github.com/golang/dep)
  ```sh
    go mod download
  ```

- Set the relevant environment variables (see .env.sample)
  ```sh
    cp .env.sample .env
    source .env
  ```

- Launch app with all dependencies
  ```sh
    make run-all-local
  ```


## Running in production

- Set the relevant environment variables (see .env.sample)
  ```sh
    cp .env.sample .env
    source .env
  ```

- Build image(CI)
  ```sh
    docker build . -t "cms:latest"
  ```

- Run image using docker(K8s/ECS)
  ```sh
    docker run cms
  ```

### Resources

#### Application Users
- **GET** `/health` Returns health of server
  Response

  ```json
  {
    "status": "ok"
  }
  ```
- **POST** `/api/company` Create a new company
  Example

  ```sh
    curl --location --request POST 'localhost:8080/api/company/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    	"name": "Apple",
    	"employeeCount": 10,
    	"registered": true,
    	"type": "Corporations"
    }
    '
  ```

  Response

  ```json
  {
    "message": "created",
    "id": "81ba414c-6609-4a80-bfba-85f6f77869df"
  }
  ```


#### Todos (requires authentication)

- **GET** `/api/company/:id` Retrieves a company details

  Example
  ```sh
  curl localhost:8080/api/v1/company/81ba414c-6609-4a80-bfba-85f6f77869df \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNTUzODg4OTkxfQ.eNn7bMfqHwA1ZF8Q87Ut0kdyZPntURuIGNuHMTvefJ8"
  ```
  Response
  ```json
  {
    "id": "d6d26322-3dba-4654-b5fe-bcfa84f2b687",
    "name": "Apple",
    "description": "MVP383",
    "employeeCount": 30,
    "registered": true,
    "type": "Corporations"
  }
  ```
- **PATCH** `/api/company/:id` Retrieves a company details

  Example
  ```sh
  curl  --request PATCH localhost:8080/api/v1/company/81ba414c-6609-4a80-bfba-85f6f77869df \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNTUzODg4OTkxfQ.eNn7bMfqHwA1ZF8Q87Ut0kdyZPntURuIGNuHMTvefJ8"
  ```
  Response
  ```json
  {
    "message": "updated"
  }
  ```
- **DELETE** `/api/company/:id` Removes provided company

  Example
  ```sh
  curl  --request DELETE localhost:8080/api/v1/company/81ba414c-6609-4a80-bfba-85f6f77869df \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiZXhwIjoxNTUzODg4OTkxfQ.eNn7bMfqHwA1ZF8Q87Ut0kdyZPntURuIGNuHMTvefJ8"
  ```
  Response
  ```json
  {
    "message": "deleted"
  }
  ```

### Licence

MIT