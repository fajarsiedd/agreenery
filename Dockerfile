#phase 1
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

#phase 2
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY google_credentials.json google_credentials.json

COPY .env .env

RUN apk update && apk add --no-cache ca-certificates

RUN apk --no-cache add tzdata

EXPOSE 80

CMD ["./main"]