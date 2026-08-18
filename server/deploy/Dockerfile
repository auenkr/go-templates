FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/server

FROM alpine:latest

RUN addgroup -S server && adduser -S -G server app

WORKDIR /app

COPY --from=builder /bin/server /app/server

USER app

EXPOSE 8080

ENTRYPOINT ["/app/server"]
