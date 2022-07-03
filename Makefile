server: 
	export LOG_TYPE=USER_FRIENDLY && go run ./cmd/server

client: 
	go run ./cmd/server

test:
	go test -cover -v ./...

build:
	go build -o main ./cmd/server

build-client:
	cd client; \
	yarn install; \
	NEXT_TELEMETRY_DISABLED=1 yarn run export

migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456789@localhost:5456/custom_holiday?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456789@localhost:5456/custom_holiday?sslmode=disable" -verbose down
