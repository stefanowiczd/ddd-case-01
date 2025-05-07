# Simple Bank

## To the readers

This project represents a comprehensive application of my accumulated expertise and technical proficiency in software development.

I also tried to leveraged advanced AI-powered coding assistance tools to enhance development efficiency and code quality.

## Goals
- Create bank application using Golang and DDD

## Assumptions
- Use DDD for the system architecture
- Use tools improving software development experience:
  - [sqlc](https://docs.sqlc.dev/en/latest/) - create Goland code from SQL
  - [gomock](https://github.com/uber-go/mock) - create unit testing mocks
  - [Cursor AI](https://www.cursor.com/)
    - Code documentation
    - Code review
    - Code auto-completion based on the project structure and code definition
  - [Docker](https://docs.docker.com/) - managing and build of images for service and integration tests
  - [Testcontainer](https://testcontainers.com/) - integration tests
  - [golang-lint](https://golangci-lint.run/) - Golang linters
- [t.b.d.] Orchestrator - apply SAGA pattern for processing events.

## Project Structure Tree
```
.
├── internal/
│   ├── application/
│   │   ├── account/      // Account service for domain and use case definition
│   │   ├── customer/     // Customer service for domain and use case definition
│   │   └── transaction/  // t.b.d.
│   ├── domain/
│   │   ├── account/      // Account domain definition
│   │   ├── customer/     // Customer domain definition
│   │   ├── event/        // Base event definition
│   │   └── transaction/  // t.b.d.
│   ├── infra/
│   │   ├── db/           // DB scheme and queries definition
│   │   └── repo/
│   │       ├── query     // Golang code generated with by SQLC related to DB operations
│   │       └── account   // Account repository code
│   │       └── customer  // Customer repository code
│   ├── interface/
│   │   ├── grpc/         // n.a.
│   │   ├── rest/         // HTTP server, handlers and routes definition,
│   │   └── stream/       // n.a.
│   └── tool/
│       └── sqlc/         // SQLC configuration
├── orchestrator/         // t.b.d. Orchestrator responsible for events processing

```

### Core Domains 
#### Account Management
- Creating, maintaining and tracing account related activities
- Different type of accounts (checking, savings, loan, etc.)
#### Customer Management
- Managing customer information, relationship and interactions
#### [t.b.d.] Transaction Processing
- handling financial operations (deposit, withdraw, money transfer, etc.)
#### [t.b.d.] Product Management
- Bank may offer different kind of products or be a broker for some products and services.


### Orchestrator [t.b.d.]
Implmentation will be realised based on SAGA 1 design pattern.

## Tools
### Project build
```shell
make build
```
### Run unit tests
```shell
make test
```
### Run integration tests
```shell
make test-integration
```
