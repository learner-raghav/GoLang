package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// We can create a struct that holds information to be displayed in put HTML file.
type Welcome struct {
	Name string
	Time string
	QuoteToShow Quote
}


//This is the entrypoint to the go application
func main(){

	//We need some random data before.
	quoteNo := rand.Intn(len(quotesArr)-1) + 0
	welcome := Welcome{"Anonymous",time.Now().Format(time.Stamp),quotesArr[quoteNo]}

	//We will tell go where it can find our HTML file.
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	//Our HTML comes with CSS, that go needs to provide when we run the application.
	//We basically have to show the static folder too from here.
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))

	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		if name := req.FormValue("name"); name != "" {
			welcome.Name = name
		}
		welcome.Time = time.Now().Format(time.Stamp)
		welcome.QuoteToShow = quotesArr[rand.Intn(len(quotesArr)-1)]
		//If the error is not equal to nil, it means some problem occurred.
		if err := templates.ExecuteTemplate(res,"welcome-template.html",welcome); err != nil {
			http.Error(res,err.Error(),http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening!!")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000",nil))

}