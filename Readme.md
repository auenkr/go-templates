# Go Templates

This repo contains a collection of Go templates for different use cases. It is intended to be used with the `gonew` CLI.

## Templates

- `simple-http-health`: minimal native `net/http` server with `GET /healthz`

## Usage

Clone a template into a new module:

```sh
gonew github.com/auenkr/go-templates/<template-to-clone> <your-module-path> ./<clone-dir-path>
```

## Example

```sh
gonew github.com/auenkr/go-templates/simple-http-health github.com/auenkr/tmp-health-test ./http-server

cd http-server
go run .
```

Test the health endpoint:

```sh
curl http://localhost:8080/healthz
```

Expected response:

```json
{"status":"ok"}
```

## Notes

- `<your-module-path>` should be a valid Go module path, such as `github.com/you/service`.
- `gonew` expects the new module path to contain a dot, for example `github.com/you/service` or `example.com/service`.
