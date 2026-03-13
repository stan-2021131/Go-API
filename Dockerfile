FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev git

COPY . .

RUN go mod init ejemplo.com/videojuegos || true
RUN go get github.com/mattn/go-sqlite3

CMD ["go", "run", "main.go"]