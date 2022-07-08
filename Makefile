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
	docker-compose up -d pgadmin4

dev-server:
	docker-compose up --build dev-server dev-db 

test-server: 
	docker-compose up --build test-server test-db

prod-server: 
	docker-compose up --build prod-server prod-db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
