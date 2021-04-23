package controller

import (
	"encoding/json"
	"net/http"

	"github.com/iamseki/go-course-api/entity"
	"github.com/iamseki/go-course-api/errors"
	"github.com/iamseki/go-course-api/service"
)

type postController struct {
	svc service.PostService
}

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

func NewPostController(s service.PostService) *postController {
	return &postController{svc: s}
}

func (c *postController) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	//data, err := json.Marshal(posts)
	posts, err := c.svc.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (c *postController) AddPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling the body request"})
		return
	}

	err = c.svc.Validate(&post)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := c.svc.Create(&post)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(result)
}
