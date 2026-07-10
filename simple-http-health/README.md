# Simple HTTP Health Server

A minimal Go template for a native `net/http` server with a health endpoint.

## Endpoint

- `GET /healthz` returns `200 OK` with `{"status":"ok"}`

## Run

```sh
go run .
```

The server listens on `:8080` by default. Override it with `ADDR`:

```sh
ADDR=:3000 go run .
```

## Use With gonew

```sh
gonew github.com/auenkr/go-templates/simple-http-health github.com/you/new-service ./new-service
```

Then run it:

```sh
cd new-service
go run .
```

Test the health endpoint:

```sh
curl http://localhost:8080/healthz
```
