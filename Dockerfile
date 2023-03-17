FROM golang:1.20 as builder

COPY . /app
WORKDIR /app
RUN go mod download

ENV CGO_ENABLED 0
RUN go build -o /assets/in ./src/in
RUN go build -o /assets/out ./src/out
RUN go build -o /assets/check ./src/check

FROM alpine:edge AS resource
RUN apk add --no-cache bash

COPY --from=builder assets/in /opt/resource/in
COPY --from=builder assets/out /opt/resource/out
COPY --from=builder assets/check /opt/resource/check

RUN chmod +x /opt/resource/in
RUN chmod +x /opt/resource/out
RUN chmod +x /opt/resource/check

FROM resource
