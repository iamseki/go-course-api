package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Wonder Woman", Text: "Marvel Movie"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	data, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the body request"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	res.WriteHeader(http.StatusCreated)

	data, err := json.Marshal(post)
	res.Write(data)
}
