FROM golang:1.20.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main ./

EXPOSE 5001

CMD ["./main"]