package main

import (
	"github.com/iamseki/go-course-api/controller"
	"github.com/iamseki/go-course-api/repository"
	"github.com/iamseki/go-course-api/router"
	"github.com/iamseki/go-course-api/service"
)

func main() {
	const port string = ":8080"

	// postRepository := repository.NewFirestorePostRepository()
	postRepository := repository.NewSQLiteRepository()
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	r := router.NewGorillaMuxRouter()

	r.GET("/posts", postController.GetPosts)
	r.POST("/posts", postController.AddPost)
	r.SERVE(port)
}
