package main

import "net/http"

func main(){

	/**
		1. http.Handle - The signature of http.handle is that it needs the string (route) and an
			implementation of interface handler. So, we need to create types that implement the handler
			interface. But the disadvantage is that, we will have to create 1000 objects for 1000 routes
			This is always a bad option. SO, to remove this, we have an alternate option

		2. http.HandlerFunc - Handler func is a user defined type that implements a handler interface
				and it makes any user defined type a handler.

			type HandlerFunc func(ResponseWriter,*request)
			func (f HandlerFunc) ServeHttp(w ResponseWriter,r *Request)
			The advantage is we have no need to create stricts or anything.

		3. If we use http.HandleFunc, the signature of this is such that, it is implcitly accepting a
	 */
	http.Handle("/",http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){

	}))
	//Directly expects a function!!
	//http.HandleFunc()
}