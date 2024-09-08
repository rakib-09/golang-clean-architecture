# Golang Clean Architecture

This project is a starter template for building scalable, maintainable, and testable applications using the Clean Architecture pattern in Go. The clean architecture focuses on organizing your code into layers, separating concerns to improve the maintainability and testability of the codebase.

## Project Structure

The project is structured in a way that separates business logic, domain models, and infrastructure concerns. Below is a breakdown of the directory structure:

```plaintext
golang-clean-architecture/
├── cmd/                 # Entry point of the application
│   └── main.go          # Main application file
├── config/              # Configuration settings (e.g., environment variables, config files)
│   └── config.go        # Config loader
├── conn/                # Database and service connections
│   └── db.go            # Database connection setup
├── controllers/         # API controllers for handling HTTP requests
│   └── user_controller.go
├── domain/              # Domain models and business logic
│   └── user.go
├── middlewares/         # Custom middleware (e.g., authentication, logging)
│   └── auth_middleware.go
├── repositories/        # Repositories for interacting with databases or external services
│   └── user_repository.go
├── routes/              # Routing definitions for the web application
│   └── routes.go
├── services/            # Application services implementing business logic
│   └── user_service.go
├── utils/               # Utility functions and helpers
│   └── hash.go
├── worker/              # Background job processing
│   └── job_worker.go
├── Dockerfile           # Docker setup
├── docker-compose.yml   # Docker Compose configuration for multi-container setups
├── go.mod               # Go modules file
├── go.sum               # Go modules checksum file
├── Makefile             # Task runner for automating common tasks
└── README.md            # Project documentation
```

### Layers Overview

1. **cmd/**:
   The main application entry point. This file is responsible for starting the server and wiring up all the layers of the application.

2. **config/**:
   Centralized configuration management. Stores application settings and environment variables.

3. **conn/**:
   Database and external service connections. This is where you configure and manage connections like PostgreSQL, MySQL, Redis, or third-party APIs.

4. **controllers/**:
   Contains API controllers that handle incoming HTTP requests and delegate them to appropriate services.

5. **domain/**:
   Contains the core domain models and business logic interfaces. This layer should remain independent of external services or frameworks.

6. **middlewares/**:
   Custom middleware functions for request handling, such as logging, authentication, and request validation.

7. **repositories/**:
   Implements the repository pattern for interacting with databases or external services. It abstracts the data source from the rest of the application.

8. **routes/**:
   Defines the routing for the API, mapping HTTP routes to controllers.

9. **services/**:
   Implements the application’s business logic by coordinating between controllers, repositories, and other services.

10. **utils/**:
    Contains helper functions and utilities that are used across the application (e.g., password hashing, string manipulation).

11. **worker/**:
    Implements background job processing for tasks that don’t need to block the request-response cycle (e.g., sending emails, processing large data sets).

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (1.19+)
- [Docker](https://www.docker.com/get-started)
- [PostgreSQL](https://www.postgresql.org/download/) or any database of your choice

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/rakib-09/golang-clean-architecture.git
   cd golang-clean-architecture
   ```

2. Build and run the application using Docker:

   ```bash
   docker-compose up --build
   ```

4. Access the application at `http://localhost:8080`.

### Running Tests

You can run the tests using the following command:

```bash
go test ./...
```

### Makefile Commands

The `Makefile` contains predefined commands to simplify your workflow:

- `make build`: Build the application.
- `make run`: Run the application locally.
- `make test`: Run the test suite.
- `make clean`: Clean the project build files.

### Technologies Used

- **Go**: The programming language used to build the application.
- **Docker**: For containerization and development environment setup.
- **PostgreSQL**: Default database, can be replaced by any other DBMS.
- **Asynq**: For background task processing.
- **log/slog**: Used for structured logging.

## Contributing

Feel free to submit issues, fork the repository, and open pull requests for improvements or bug fixes.

```

This `README.md` provides a detailed overview of your project’s architecture, setup, and usage. Let me know if you'd like to make any adjustments!
