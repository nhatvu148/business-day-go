## How to start server

- make server

## How to run test

- make test

## Run docker compose for test

- docker-compose up --build test-server test-db

## Run docker compose for development

- docker-compose up --build dev-server dev-db

## Run docker compose for production

- docker-compose up --build prod-server prod-db

## Create database migration

- migrate create -ext sql -dir db/migration -seq custom_holiday
