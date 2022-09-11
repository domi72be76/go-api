package main

import (
	"bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/hello", getUsers).Methods("GET")
	router.HandleFunc("/hello/{Username}", birthDayMsg).Methods("GET")
	router.HandleFunc("/hello/{Username}", addAndUpdatedUser).Methods("PUT")
	return router
}

func TestHome(t *testing.T) {
    request, _ := http.NewRequest("GET", "/", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetUsers(t *testing.T) {
    request, _ := http.NewRequest("GET", "/hello", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestBirthDayMsg(t *testing.T) {
    request, _ := http.NewRequest("GET", "/hello/domi", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestAddAndUpdatedUser(t *testing.T) {
	//data use for to add mew user
	var jsonData = []byte(`{
		"DateOfBirth": "1999-09-22"
	}`)
	b := bytes.NewBuffer(jsonData)

    request, _ := http.NewRequest("PUT", "/hello/domiZ", b)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 204, response.Code, "OK response is expected")
}