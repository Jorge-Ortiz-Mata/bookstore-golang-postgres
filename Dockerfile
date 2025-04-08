# Dockerfile
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Imagen mínima sin libc
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8000
CMD ["./main"]
