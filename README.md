# go_restful

A simple REST API server which stores map[string]string and provides add, update, delete, get interfaces and a client which uses the server api to access and interact with the server. The mode of communication is JSON.

Build server: go build server.go types.go
Build client: go build client.go types.go

server runs on port 8080 and the client assumes the server is on local host. 

# Usage of ./client:
  -add string
        key,value to be added
  -delete string
        key to be deleted
  -get string
        key to retrieve
  -getAll
        get all keys
  -update string
        key,value to be updated

# Sample client run:

test#./client -add="foo,bar"
Status: 201 Created

test#./client -getAll
Status: 200 OK
[{"key":"foo","value":"bar"},{"key":"abc","value":"def"}]

test#./client -update="abc,xyz"
Status: 200 OK

test#./client -getAll
Status: 200 OK
[{"key":"foo","value":"bar"},{"key":"abc","value":"xyz"}]

test#./client -get="foo"
Status: 200 OK
{"key":"foo","value":"bar"}

test#./client -delete="abc"
Status: 200 OK

test#./client -getAll
Status: 200 OK
[{"key":"foo","value":"bar"}]
test#

# References:

https://golang.org/pkg/net/http/
https://golang.org/pkg/encoding/json/
http://www.gorillatoolkit.org/pkg/mux#api
https://www.easypost.com/docs/api.html
https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
https://gobyexample.com/
https://mholt.github.io/curl-to-go/
www.stackoverflow.com
