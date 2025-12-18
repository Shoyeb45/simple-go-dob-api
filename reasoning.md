# Reasoning Behind the Design and Architecture

## Entry Point

* The entry point of the application is [cmd/server/main.go](./cmd/server/main.go).
* As the name suggests, it is the starting point of the application. It first reads the environment variables.
* Since the application depends on environment variables, they are read and validated at the very beginning.
* After that, it initializes the `uber-zap-logger`.
* This is followed by database initialization, where the application connects to the database and maintains a connection pool.
* These components are initialized only once and reused throughout the application, without the need to reinitialize or reread them.
* This makes later usage of these components very easy and efficient.
* Finally, the server is started on the configured port.

## Database

* For this task, I have chosen PostgreSQL along with SQLC to convert SQL queries into type-safe Go functions.
* SQLC is used to maintain full control over SQL queries while still providing type safety.
* Everything related to SQLC is located in [./sqlc/*](./sqlc).

## Centralized `AppError` and Proper Error Handling

* Error handling is a critical part of any application.
* To handle this properly, I created an `AppError` struct along with several specific error types to format errors consistently and send meaningful responses to the client.
* I also implemented an `ErrorHandler` middleware that catches errors, logs them within the application, and sends appropriate responses to the client.
* Correct HTTP status codes are used for different error types, such as `400` for `BadRequestError`.

## Routes, Handler, Repository, and Service

* A three-layered architecture is used to make the application scalable and maintainable.
* In [./internal/routes](./internal/routes), all routes related to `users` are defined.
* The [repository layer](./internal/repository) is responsible for interacting with the database and exposes functions such as `FindById`, `DeleteById`, and `GetUsers`.
* All business logic resides in the [service layer](./internal/service/).
* Finally, in the [handlers](./internal/handler/) layer, user input is collected and validated, and responses are prepared and sent back to the client.

## App.go

* [app.go](./internal/app/) is located in `./internal/app/`.
* Instead of initializing routes or adding middleware in the main entry point, all application setup is handled in this file.
* This includes route registration and manual dependency injection.
* Dependencies are wired in a clear flow: **repository → service → handler**.

## Docker and Deployment

* A production-ready [Dockerfile](./Dockerfile) is created using a Go Alpine image.
* The application is deployed on Render.
* Application link:
  [https://simple-go-dob-api.onrender.com/](https://simple-go-dob-api.onrender.com/)


