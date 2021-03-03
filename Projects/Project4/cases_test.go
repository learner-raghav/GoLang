package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
	SQL mock is a mock library for implementing the sql/driver.
	It has one and only one purpose. To simulate any sql driver behaviour in tests, without needing
	a real database connection. It helps to maintain correct TDD Workflow.
 */

func (a *api) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestShouldGetPosts(t *testing.T)  {
	db,mock,err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' occured",err)
	}
	defer db.Close()

	//Create app with mocked db, request and response to test.
	app := &api{db: db}
	req,err := http.NewRequest("GET","http://localhost/posts",nil)
	if err != nil {
		t.Fatalf("An error '%f' occured while creating the request",err)
	}

	w := httptest.NewRecorder()

	//Before we fetch we need to put some data into
	rows := sqlmock.NewRows([]string{"id","title","body"}).
		AddRow(1,"post 1","Hello").AddRow(2,"post 2","World")

	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	app.posts(w,req)

	if w.Code != 200 {
		t.Fatalf("Expected status code to be 200, bot got: %d",w.Code)
	}

	data := struct{
		Posts []post
	}{
		Posts: []post{
			{
				ID: 1,
				Text: "Hello",
				Title: "post 1",
			},
			{
				ID: 2,
				Text: "World",
				Title: "post 2",
			},

		},
	}

	app.assertJSON(w.Body.Bytes(),data,t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s",err)
	}
}

func TestShouldRespondWithErrorOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &api{db}
	req, err := http.NewRequest("GET", "http://localhost/posts", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnError(fmt.Errorf("some error"))

	// now we execute our request
	app.posts(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	data := struct {
		Error string
	}{"Failed to fetch postssome error"}
	app.assertJSON(w.Body.Bytes(), data, t)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}