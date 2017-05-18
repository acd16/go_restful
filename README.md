# go_restful

A simple REST API server which stores map[string]string and provides add, update, delete, get interfaces and a client which uses the server api to access and interact with the server. The mode of communication is JSON.

Build server: go build server.go types.go
Build client: go build client.go types.go

server runs on port 8080 and the client assumes the server is on local host. 


Sample client run:

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
