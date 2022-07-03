## How to start server

- make server

## How to run test

- make test

## Run docker compose for development

- docker-compose -f docker-compose-dev.yml up --build

## Run docker compose for production

- docker-compose up --build

## Migrate database

- migrate create -ext sql -dir db/migration -seq custom_holiday
