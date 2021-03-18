# Todos with Oto & friends

## Dependencies

- go
- docker (for tests)
- postgres

## Setup

- clone this repo
- Install dependencies
  ```
  make install-tools
  ```
- run server
  ```
  go run main.go
  ```
- run tests (require docker running)
  ```
  make test
  ```
- regenerate types after updating defenitions
  ```
  make generate
  ```
