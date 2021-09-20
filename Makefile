port = 8000

run:
	go run main.go -port ${port}

test:
	go test ./...

compose:
	PORT=${port} docker-compose up -d

compose-down:
	PORT=${port} docker-compose down
