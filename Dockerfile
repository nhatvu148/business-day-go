# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN cd ./server && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main main.go

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/server/main .

CMD [ "/app/main" ]