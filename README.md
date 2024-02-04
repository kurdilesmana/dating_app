# Dating App

## Docker
Clone this repository and run:
```
docker-compose up 
```

Following endpoints:

| Method | Route                | Body                                               |
| ------ | -------------------- |----------------------------------------------------|
| POST   | /register            | `{"username": "dian", "email": "1234@email.com", "password": "0821ss" }`|


## Structure adapt Hexagonal Architecture:
``` bash
dating-app/
|-- cmd/
|   -- main.go
|-- internal/
|   |-- domain/ 
|   |   |-- user.go
|   |-- repository/
|   |   |-- user_repository.go 
|   |-- service/
|   |   |-- user_service.go
|   |-- delivery/
|   |   |-- http/
|   |       |-- handler/
|   |       |   -- user_handler.go
|   |       |-- router/
|   |           |-- user_router.go
|   |-- infrastructure/
|       |-- database/
|           |-- database.go
|-- config/
    |-- config.go
```

- cmd/:
  - main.go: 
    - The entry point of your application. This is where you initialize and wire up your dependencies, configure your application, and start the execution.
- internal/:
  - domain/: Contains the core business logic or domain model of your application.
    - user.go: 
      - Defines the User struct, which represents the core entity in your application.
  - repository/: Contains interfaces or implementations responsible for data access.
    - user_repository.go: 
      - Defines the interface for a user repository, specifying methods like CreateUser, GetUserByID, etc.
  - service/: Houses the application's business logic or use cases.
    - user_service.go: 
      - Implements methods related to user use cases, such as creating a user, updating a user, etc.
  - delivery/: Handles the delivery mechanism of your application, such as HTTP, CLI, or messaging.
    - http/: Specifically handles HTTP-related concerns.
      - handler/: Contains HTTP request handlers, responsible for handling HTTP requests and interacting with the application layer.
        - user_handler.go: 
          - Implements HTTP handlers for user-related operations.
      - router/: Contains the router setup for your HTTP server.
        - user_router.go: 
          - Configures the routes for user-related endpoints.
  - infrastructure/: Houses infrastructure-related concerns, such as databases or external services.
    - database/: Contains code related to database setup and configuration.
      - database.go: 
        - Initializes and configures the database connection.
- config/:
  - config.go: 
    - Contains code related to application configuration, likely using a configuration library like Viper to manage environment variables and configuration files.
