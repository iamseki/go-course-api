package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iamseki/go-course-api/controller"
	"github.com/iamseki/go-course-api/entity"
	"github.com/iamseki/go-course-api/repository"
	"github.com/iamseki/go-course-api/service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int    = 123
	TITLE string = "title 1"
	TEXT  string = "text 1"
)

var (
	postRepo repository.PostRepository = repository.NewSQLiteRepository()
	postSrv  service.PostService       = service.NewPostService(postRepo)
	sut      controller.PostController = controller.NewPostController(postSrv)
)

func tearDown(postID int) {
	var post entity.Post = entity.Post{
		ID: postID,
	}
	postRepo.Delete(&post)
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&post)
}

func TestAddPost(t *testing.T) {
	// Create new HTTP request
	jsonStr := []byte(`{"title":"` + TITLE + `","text":"` + TEXT + `"}`)
	request, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// http stuff
	handler := http.HandlerFunc(sut.AddPost)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	tearDown(post.ID)
}

func TestGetPosts(t *testing.T) {
	setup()

	request, _ := http.NewRequest("GET", "/posts", nil)
	handler := http.HandlerFunc(sut.GetPosts)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.Equal(t, ID, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// Cleanup database
	tearDown(ID)
}
