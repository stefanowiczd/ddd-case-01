# Simple Bank

## To the readers

This project represents a comprehensive application of my accumulated expertise and technical proficiency in software development.

I also tried to leveraged advanced AI-powered coding assistance tools to enhance development efficiency and code quality.


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
- [t.b.d.] Apply SAGA pattern for processing events.

## Project Structure:
- `internal/domain` contains the core business logic and entities
- `internal/application` contains use cases and application services
- `internal/infrastructure` contains implementations of repositories and external services
- `internal/interfaces` contains API handlers and other interfaces

### Core Domains 
#### Account Management
- Creating, maintaining and tracing account related activities
- Different type of accounts (checking, savings, loan, etc.)
#### [t.b.d.] Customer Management
- Managing customer information, relationship and interactions
#### [t.b.d.] Transaction Processing
- handling financial operations (deposit, withdraw, money transfer, etc.)
#### [t.b.d.] Product Management
- Bank may offer different kind of products or be a broker for some products and services.

