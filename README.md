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

## Project Structure Tree
```
.
├── internal/
│   ├── application/
│   │   ├── account/      // Account service for domain and use case definition
│   │   ├── customer/     // t.b.d.
│   │   └── transaction/  // t.b.d.
│   ├── domain/
│   │   ├── account/      // Account domain definition
│   │   ├── customer/     // t.b.d.
│   │   ├── event/        // Base event definition
│   │   └── transaction/  // t.b.d.
│   ├── infra/
│   │   ├── db/           // DB scheme and queries definition
│   │   └── repo/
│   │       ├── query     // Golang code generated with by SQLC related to DB operations
│   │       └── account   // Account repository code
│   ├── interface/
│   │   ├── grpc/         // n.a.
│   │   ├── rest/         // HTTP server, handlers and route definition,
│   │   └── stream/       // n.a.
│   └── tool/
│       └── sqlc/         // SQLC configuration
├── orchestrator/         // t.b.d. Orchestrator responsible for events processing

```

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


