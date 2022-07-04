## How to start server

- make server

## How to run test

- make test

## Run docker compose for test

- docker-compose up test-server test-db --build

## Run docker compose for development

- docker-compose up dev-server dev-db --build

## Run docker compose for production

- docker-compose up prod-server prod-db --build

## Create database migration

- migrate create -ext sql -dir db/migration -seq custom_holiday
