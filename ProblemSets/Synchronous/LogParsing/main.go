package main

var rawLogs = `
[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] GET /usr/local/1231/files HTTP/1.1 204
[03:04:20 UTC] POST /usr/local/1231/files HTTP/1.1 404
[03:04:20 UTC] PATCH /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] DELETE /usr/local/1231/files HTTP/1.1 201
[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] GET /usr/local/1231/files HTTP/1.1 404
[03:04:20 UTC] DELETE /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 500
[03:04:20 UTC] PATCH /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] PATCH /usr/local/1231/files HTTP/1.1 500
[03:04:20 UTC] POST /usr/local/1231/files HTTP/1.1 204
[03:04:20 UTC] GET /usr/local/1231/files HTTP/1.1 200
`

func main() {

}
