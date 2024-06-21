FROM golang:1.22.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

RUN go build -o main .

EXPOSE 8060
CMD ["go", "run", "server.go"]