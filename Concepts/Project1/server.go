package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	const port string = ":5000"
	router.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprintf(res,"Up and Running!")
	})

	//The advantage of using Gorilla Mux is that, we can specify the Methods too
	router.HandleFunc("/posts",getPosts).Methods("GET")
	router.HandleFunc("/addPost",addPost).Methods("POST")
	log.Println("Server listening on PORT: ",port)
	http.ListenAndServe(port,router)
}
