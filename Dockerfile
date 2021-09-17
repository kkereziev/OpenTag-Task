FROM golang:1.16-alpine

COPY . /app

ARG PORT

WORKDIR /app

RUN go mod download\
    && go build -o server -ldflags "-X main.Port=${PORT}"

EXPOSE ${PORT}

CMD ["./server"]