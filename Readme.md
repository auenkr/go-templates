# Go Templates

This repo contains a collection of Go templates for different use cases. It is intended to be used with the `gonew` CLI.

## Templates

- `simple-http-health`: minimal native `net/http` server with `GET /healthz`

## Usage

Clone a template into a new module:

```sh
gonew github.com/AuenKr/go-templates/simple-http-health github.com/you/new-service
```

Then run it:

```sh
cd new-service
go run .
```
