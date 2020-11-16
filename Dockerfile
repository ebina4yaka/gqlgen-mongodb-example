FROM golang:1.15.5-alpine3.12 AS builder

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -o server server.go


FROM alpine:3.12.0 AS app

COPY --from=builder /build/server /usr/local/bin
CMD server
