package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
	We give the corresponding json id's as well.
 */

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}

var (
	posts []Post
)

func init(){
	posts = []Post{{Id: 1,Title: "A Ray of hope",Text: "A ray of hope is really needed!"}}
}

func getPosts(resp http.ResponseWriter,req *http.Request){
	resp.Header().Set("Content-Type","application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error Marshalling the Post Array"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter,req *http.Request){
	resp.Header().Set("Content-Type","application/json")
	var post Post
	log.Print(req.Body)
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts,post)
	resp.WriteHeader(http.StatusOK)
	result,err := json.Marshal(post)
	resp.Write(result)
}