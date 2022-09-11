# Go-API

A simple API that help to manage user date of birth records. To do so we use a Goland app and a dummy in-memory 'database'.

## Setup

### First step/Requirements
- install git, go, docker
- clone the go-api repositoty

### Preparatiion
$ export GOPATH=<.../\>go-api\
$ cd $GOPATH\
$ go get github.com/gorilla/mux\
$ go get github.com/stretchr/testify/assert\
$ cd src

### Run without docker locally
$ go build -o ./server .\
$ ./server

### Run with docker locally
$ docker build -t go-api .\
$ docker run -it --rm -p 9000:9000 go-api 

## Use the API

### Add or update a record
$ curl -X PUT http://localhost:9000/hello/<username\> -H "Content-Type: application/json"  -d '{"DateOfBirth": "<YYYY-MM-DD\>"}'

### Get the birthday message
$  curl http://127.0.0.1:9000/hello/<username\>

### Get the complete list of records
$  curl http://127.0.0.1:9000/hello
