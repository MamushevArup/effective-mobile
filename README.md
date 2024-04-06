# Car Catalog Project

This project is a car catalog application that allows users to perform various operations on car data, such as retrieving, updating, and deleting car records.

## Prerequisites

Before running the project, ensure you have the following dependencies installed:

- [Goose](https://github.com/pressly/goose) or any other migration tool for managing database migrations.
- Go programming language installed on your system.
- Postgres database to connect and store data

## Installation and Setup

### Go to ```cd effective-mobile ```

1. **Migration**
If you have goose installed then
```
goose -dir ./schema "postgres://<username>:<password>@<host>:<port>/<dbname>" up
```
## Please paste the data which actual for you.

2. **Mock server**
   ## To imititate external api I created mock which return static json
   ```
   go run mock/api/server.go 
   ```
3. **Start actual server**
   ## All endpoints documented with swagger
   ```
   go run cmd/main.go
   ```
   
