# Simple Books API
Simple Books API (with [OpenAPI 3.0 specification](./api-docs.yml)) implemented in Golang.

## Testing
### Generating Mocks
```shell
go install github.com/vektra/mockery/v2@v2.20.0

go generate ./...
```
### Running Tests
To run tests:
```shell
export BASIC_AUTH_USER="admin"
export BASIC_AUTH_PASSWORD="password"

export DB_USER="admin"
export DB_PASSWORD="password"
export DB_NAME="gorm"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_SSLMODE="disable"

go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### Deploy Development Server (Docker)
Deploy deployment server and PosgreSQL DB with `docker-compose`.
```shell 
cd .docker
docker-compose up --build

# Stop the server and database
docker-compose down
```

### Deploy Development Server
You can also deploy a development server without using `docker-compose`. 
However, you must provide a database server to connect to.
```shell
export BASIC_AUTH_USER="admin"
export BASIC_AUTH_PASSWORD="password"

export DB_USER="admin"
export DB_PASSWORD="password"
export DB_NAME="gorm"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_SSLMODE="disable"

go run .
```

## Usage
### Client Requests
```shell
export BASIC_AUTH_USER="admin"
export BASIC_AUTH_PASSWORD="password"

export BASE_URL="http://localhost:8080"

export API_URL="${BASE_URL}/api/v1"
export DOCS_URL="${BASE_URL}/api-docs"

export HEALTH_URL="${API_URL}/healthz"
export BOOK_URL="${API_URL}/book"

# Check health
curl -v "${HEALTH_URL}"

# Show API Documentation
curl -v "${DOCS_URL}"

# List books
curl -v "${BOOK_URL}" | jq

# Create book
curl -v -X POST \
  -u "${BASIC_AUTH_USER}:${BASIC_AUTH_PASSWORD}" \
  -H "Content-Type: application/json" \
  --data '{"title":"The Power of Geography","author":"Tim Marshall","year":2009}' \
  "${BOOK_URL}" | jq

# Read book
curl -v \
  -u "${BASIC_AUTH_USER}:${BASIC_AUTH_PASSWORD}" \
  "${BOOK_URL}/ID" | jq

# Update book
curl -v -X PUT \
  -u "${BASIC_AUTH_USER}:${BASIC_AUTH_PASSWORD}" \
  -H "Content-Type: application/json" \
  --data '{"title":"The Gods are to blame","author":"John Doe","year":1992}' \
  "${BOOK_URL}/ID" | jq

# Delete book
curl -v -X DELETE \
  -u "${BASIC_AUTH_USER}:${BASIC_AUTH_PASSWORD}" \
  "${BOOK_URL}/ID" | jq
```
