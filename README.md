# Golang microservice template

This application is a template to create microservices in Golang with initial configurations, to minimizes initial work when create a project from scratch.

## Used technologies

- Go
- Docker
- Gofiber v2
- PostgreSQL
- Migrations
- Hexagonal Architecture
- Shell script
- Windows bat script
- Terraform

## How to use

1. You need docker installed and running
2. If you're using linux, execute this command from project root folder: `./scripts/run-local-linux.sh`
3. If you're using windows, execute this command from project root folder: `C:\scripts\run-local-windows.bat`

This will build database docker image and start the application locally

## Migrations

The only rule to follow when using migrations is to put migration file inside `migrations` folder and name it with the following pattern: `timestamp_migration_name.up.sql`

Example: `1680966618_initial.up.sql`

## Endpoints

|HTTP Verb|URN|Description|
|---------|---|-----------|
|GET|/health|Return the application state|
|GET|/users|Return all users|
|POST|/users|Create a user|

## Testing

### Executing unit tests + coverage
```shell
go test -v ./tests/unit -cover -coverpkg=./application/services/... --bench=. --benchmem -coverprofile=coverage.out
```

### Render in HTML the coverage results

```shell
go tool cover -html=coverage.out
```

## How to run it locally?

First you need to execute Docker daemon, it needs to be running.

Bellow is the local database settings

```text
DATABASE_HOST=127.0.0.1
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASS=postgres
DATABASE_NAME=postgres
```

### Command on Windows

Open the command prompt, type command bellow and press enter(inside the project folder)
* `C:\scripts\run-local-windows.bat`

### Command on Unix-like

* `./scripts/run-local-linux.sh

**Created with** :heart: by [Rodolfo Azevedo](https://github.com/rof20004)
