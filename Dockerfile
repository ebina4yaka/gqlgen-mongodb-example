FROM golang:1.15.5-alpine3.12 AS builder

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -v


FROM alpine:3.12.0 AS app

COPY --from=builder /build/app /usr/local/bin
CMD app
