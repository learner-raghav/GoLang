package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		fmt.Println("Hello World ")
		fmt.Fprintf(res,"Hello World")
	})
	http.ListenAndServe("0.0.0.0:3000",nil)
}
