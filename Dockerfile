FROM golang:1.20 as builder

COPY . /app
WORKDIR /app
RUN go mod download

ENV CGO_ENABLED 0
RUN go build -o /assert/in ./src/in
RUN go build -o /assets/out ./src/out
RUN go build -o /assets/check ./src/check

FROM alpine:edge AS resource
RUN apk add --no-cache bash
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
