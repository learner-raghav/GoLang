package main

import "net/http"

func main(){
	//Building a simple file server!!
	//Handler interface - ServeHTTP method
	//ResponseWriter, Request objects
	//Write([]byte), Header(int), Header header  -Interface has 3 methods
	http.ListenAndServe("0.0.0.0:3000",http.FileServer(http.Dir("/home/raramuri/Desktop/GoLang/")))
}
