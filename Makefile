DB_URL=postgresql://postgres:123456789@localhost:5456/custom_holiday?sslmode=disable

server: 
	go run ./cmd/server

test:
	go test -cover -v ./...

test1:
	go test -v ./... -run CustomHolidayHandler

build:
	go build -o main ./cmd/server

build-client:
	cd client; \
	yarn install; \
	NEXT_TELEMETRY_DISABLED=1 yarn run export

pgadmin4:
	docker-compose up pgadmin4 -d

dev-server:
	docker-compose up dev-server dev-db --build

test-server: 
	docker-compose up test-server test-db --build

prod-server: 
	docker-compose up prod-server prod-db --build

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
