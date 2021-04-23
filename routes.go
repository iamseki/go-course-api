package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/iamseki/go-course-api/entity"
	"github.com/iamseki/go-course-api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	//data, err := json.Marshal(posts)
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the body request"}`))
		return
	}

	post.ID = rand.Int()
	repo.Save(&post)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(post)
}
