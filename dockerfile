FROM golang:1.22.2

WORKDIR /app/go-inventory

COPY . /app/go-inventory

RUN go mod tidy

CMD ["go", "run", "/app/go-inventory/main.go"]

