package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET","http://localhost:3000",nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	HelloWorld(res,req)
	exp := "Hello World"
	received := res.Body.String()
	if exp != received {
		t.Fatalf("Error occured. Expected: %s, Received:  %s",exp,received)
	}
}
