DB_URL=postgresql://postgres:123456789@postgres:5456/custom_holiday?sslmode=disable
LOCAL_DB_URL=postgresql://postgres:123456789@localhost:5456/custom_holiday?sslmode=disable

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
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup-local:
	migrate -path db/migration -database "$(LOCAL_DB_URL)" -verbose up

migratedown-local:
	migrate -path db/migration -database "$(LOCAL_DB_URL)" -verbose down
