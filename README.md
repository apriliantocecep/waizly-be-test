# WAIZLY Test

## Basic Test Answer

For basic test answer, file prefix is `basic_test_`

## Implement Test 2 Answer

The file is `implement_test_2_answer.txt`

## Description

This is golang app for invoice system using clean architecture template.


## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Framework & Library

- Gin (HTTP Framework) : https://github.com/gin-gonic/gin
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Google Wire (Dependency Injection) : https://github.com/google/wire

## Configuration

All configuration is in `.env` file. You can found the .env example file in root `configs` folder.

## Database Migration

All database migration is in `database/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir database/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:@tcp(localhost:3306)/waizly?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations up
```

## Run Application

You can run locally and download the package first with command `go mod download`

### Run unit test

```bash
go test -v ./test/ -coverprofile=coverage.out
```

### Run web server

```bash
go run cmd/web/main.go
```

## Run Application with Docker (Locally)

Make sure Docker Desktop already running

Run this command to running:

```shell
docker compose up --build -d
```

Run this command to remove:

```shell
docker compose down -v
```

# API Spec

Download and Import the collection file into Postman

`Waizly API.postman_collection.json`