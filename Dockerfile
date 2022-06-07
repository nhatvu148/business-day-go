# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/server

# Run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

CMD [ "/app/main" ]