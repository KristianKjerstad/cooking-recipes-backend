FROM golang:1.22.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/air-verse/air@latest
RUN apk add build-base

RUN go build -o main .

EXPOSE 8060
CMD ["air"]