# Transaction

The project is a Golang REST API that uses gorilla/mux as router framework, SQLite as a database and Docker to build.

# Getting started

## Locally

To run project locally you need to export the environment variable `PORT=:8000` or to other port to your choise and you must have installed:

- Golang >=1.16
- SQLite3

```shell
make run-local
```

```shell
make run-tests
```

## Docker

```shell
make build
make docker-image
make docker-run
```

or just

```shell
make docker-up
```

The default PORT to test via Docker is `8080`

to run **tests on Docker**:

```shell
make docker-tests
```

# Usage

## `POST /account`

The body request to **create an account** is:

```shell
{
    "document_number": "70620459000173"
}
```
The request is:

    curl -X POST 'http://localhost:8080/account' -H 'Content-Type: application/json' --data-raw '{
    "document_number": "70620459000173"}'

## `GET /account/{id}`

To **get an account** the request is:

    curl -X GET 'http://localhost:8080/account/1'

## `POST /transaction`

The body request to **create a transaction** is:

```shell
{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": 600.45
}
```
The request is:

    curl -X POST 'http://localhost:8080/transaction' -H 'Content-Type: application/json' --data-raw '{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": 600.45}'