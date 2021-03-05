package main

import (
	"fmt"
	"net/http"
)

/**
	If we need to have some code that needs to run for every request, regardless of the route
	it will eventually end up invoking, We need to have something called as a middleware.
 */

func HelloWorld(res http.ResponseWriter,req *http.Request){
	fmt.Fprintf(res,"Hello world")
}

func main(){
	http.HandleFunc("/",HelloWorld)
	http.ListenAndServe(":3000",nil)
}

/**
	Unit testing is meant to run applications through the whole request cycle and not just the
	individual functions.
	Mock-Gen!!
 */