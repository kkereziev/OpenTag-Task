#default values
containerport = 8000 #this port is used for the container
port = 8000 # this port is used for the docker host

run:
	go run main.go -port ${port}

test:
	go test ./...

compose:
	PORT=${port} CONTAINERPORT=${containerport} docker-compose up -d

compose-down:
	PORT=${port} CONTAINERPORT=${containerport} docker-compose down
