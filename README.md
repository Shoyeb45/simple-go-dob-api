# Simple Date of Birth App

## Tech stack



## Set up in local

1. Make sure to have `go-version` >= 1.25.5
2. Install `sqlc`:
    ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

3. Install [`golang-migrate`](https://github.com/golang-migrate/migrate):
    ```bash
    # mac
    brew install golang-migrate

    # Linux
    curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
    sudo mv migrate /usr/local/bin
    ```

4. Generate sqlc database interfaces:
    ```bash
    sqlc generate
    ```

5. Generate initial migration or skip.
    ```bash
    migrate create -ext sql -dir db/migrations -seq init
    ```

6. Push the migration to DB.
    ```bash
    migrate -path ./db/migrations/ -database <DB_URL> up
    ```

7. Run Go Application:
    ```bash
    # development
    go run cmd/server/main.go
    ```