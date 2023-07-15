FROM golang:latest AS build

ENV GO111MODULE=on

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -C cmd -o main

FROM alpine:latest

COPY --from=build /app/cmd/main ./

RUN chmod +x ./main

EXPOSE 5001

CMD ./main