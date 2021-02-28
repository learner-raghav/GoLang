package repository

import (

	"../entity"
	firebase "firebase.google.com/go"
	"fmt"
	"context"
	"google.golang.org/api"

)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post,error)
}

type repo struct {

}

// New Post Repository
func NewPostRepository() PostRepository {
	return &repo{}
}

const projectId = "golang1-e63c6"


func (repo *repo) Save(post *entity.Post) (*entity.Post,error) {


}

func (repo *repo) FindAll() ([]entity.Post,error) {

}