# Simple Bank

## Goals
- Create bank application using Golang and DDD

## Assumptions
- Use DDD for as a system architecture
- Use tools improving software development experience:
  - [sqlc](https://docs.sqlc.dev/en/latest/) - create Goland code from SQL
  - [gomock](https://github.com/uber-go/mock) - create unit testing mocks
  - [Cursor AI](https://www.cursor.com/) - coding assistance, documentation, etc. 
  - [Docker](https://docs.docker.com/) - managing and build of images for service and integration tests
  - [Testcontainer](https://testcontainers.com/) - integration tests
  - [golang-lint](https://golangci-lint.run/) - Golang linters

## Project Structure:
- `internal/domain` contains the core business logic and entities
- `internal/application` contains use cases and application services
- `internal/infrastructure` contains implementations of repositories and external services
- `internal/interfaces` contains API handlers and other interfaces

### Core Domains 
#### Account Management
- Creating, maintaining and tracing account related activities
- Different type of accounts (checking, savings, loan, etc.)
#### Customer Management
- Managing customer information, relationship and interactions
#### Transaction Processing
- handling financial operations (deposit, withdraw, money transfer, etc.)
#### Product Management
- Bank may offer different kind of products or be a broker for some products and services.



