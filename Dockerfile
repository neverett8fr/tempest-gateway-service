# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2-buster AS builder

WORKDIR /tempest-gateway-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN go build -o /tempest-gateway-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /tempest-gateway-service ./

EXPOSE 8080

ENTRYPOINT ["/tempest-gateway-service"]