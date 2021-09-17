run:
	go run main.go -port ${port}

compose:
	PORT=${port} docker-compose up -d

compose-down:
	PORT="8" docker-compose down
