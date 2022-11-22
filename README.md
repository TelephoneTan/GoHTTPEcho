# GoHTTPEcho

GoLang HTTP server that echoes back whatever you send in HTTP request body.

## Usage

```shell
cd GoHTTPEcho
go run ./main &
curl -X POST --data "hello, world" http://localhost:3000
```