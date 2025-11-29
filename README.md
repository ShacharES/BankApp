# BankApp

Learning-focused banking backend with account management, transfers, authentication, and migrations.

## Features
- Create/login users with token-based auth (PASETO)
- CRUD accounts with currency validation
- Atomic transfers with double-entry bookkeeping
- SQLC-generated queries, mocks, and unit tests

## Stack
- Go 1.23, Gin, Viper, sqlc, golang-migrate
- PostgreSQL 14
- Docker/Docker Compose for local orchestration

## Prerequisites
- Go 1.23+ (for local runs)
- Docker + Docker Compose
- `migrate` CLI v4, `sqlc` v2, `mockgen` (optional: only when regenerating SQL or mocks)

## Local Development (Go + Dockerized Postgres)
- Create the shared network once: `docker network create bank-network` (skip if it exists)
- Start Postgres: `make postgres` (or `make rundb` after the first start)
- Create the database: `make createdb`
- Apply migrations: `make migrateup`
- Run the API: `make server` (listens on `0.0.0.0:8080`)

## Docker Compose
- Ensure `app.env` has `DB_SOURCE=postgresql://root:secret@postgres:5432/bankapp?sslmode=disable` so the API container can reach the DB service
- Build and start everything: `docker compose up --build`
- The API exposes `http://localhost:8080`; migrations run on container startup via `start.sh`

## Migrations
- Up: `make migrateup`
- Down last batch: `make migratedown1`
- Create DB locally: `make createdb` / drop: `make dropdb`

## Code Generation & Tests
- Regenerate SQLC code: `make sqlc`
- Regenerate DB mocks: `make mock`
- Run tests: `make test`

## Configuration
- `DB_DRIVER` (postgres)
- `DB_SOURCE` (connection string)
- `SERVER_ADDRESS` (e.g., `0.0.0.0:8080`)
- `TOKEN_SYMMETRIC_KEY` (32-char secret)
- `ACCESS_TOKEN_DURATION` (e.g., `15m`)

## API Quick Reference
- `POST /users` – register; `POST /users/login` – obtain access token
- Auth header: `Authorization: Bearer <token>`
- Accounts (auth): `POST /accounts`, `GET /accounts/:id`, `GET /accounts`, `PUT /accounts/:id`, `DELETE /accounts/:id`
- Transfers (auth): `POST /transfers` (from -> to account)
