.PHONY: server
server: 
	cd ./server && go run main.go

client: 
	cd ./cmd/cli && go run main.go