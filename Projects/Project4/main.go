package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type api struct {
	db *sql.DB
}

type post struct {
	ID int
	Title string
	Text string
}

func (a *api) posts(res http.ResponseWriter,req *http.Request){

	rows, err := a.db.Query("SELECT id,title,body FROM posts")
	if err != nil {
		a.fail(res,"Failed to fetch posts"+err.Error(),500)
		return
	}
	defer rows.Close()

	var posts []post
	for rows.Next() {
		p := post{}
		if err := rows.Scan(&p.ID,&p.Title,&p.Text); err != nil {
			a.fail(res,"Failed to scan post: "+err.Error(),500)
			return
		}
		posts = append(posts,p)
	}

	if rows.Err() != nil {
		a.fail(res,"Failed to read all posts"+rows.Err().Error(),500)
		return
	}

	data := struct{
		Posts []post
	} {Posts: posts}

	a.ok(res,data)
}


func (a *api) fail(w http.ResponseWriter,msg string,status int){
	w.Header().Set("Content-Type","application/json")
	data := struct {
		Error string
	}{msg}

	res,_ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(res)
}

func (a *api) ok(w http.ResponseWriter,data interface{}){
	w.Header().Set("Content-Type","application/json")
	resp,err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.fail(w,"Something went wrong",500)
		return
	}
	w.Write(resp)
}

//func main(){
//	dbDriver := "mysql"
//	dbUser := "raghav"
//	dbPass := "raghav@123M"
//	dbName := "testDB"
//
//	db,err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
//	if err != nil {
//		panic(err)
//	}
//
//	app := &api{db: db}
//	http.HandleFunc("/posts",app.posts)
//	http.ListenAndServe("0.0.0.0:5000",nil)
//}
