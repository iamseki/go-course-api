package service

import (
	"errors"
	"math/rand"

	"github.com/iamseki/go-course-api/entity"
	"github.com/iamseki/go-course-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{ repository repository.PostRepository }

func NewPostService(r repository.PostRepository) PostService {
	return &service{repository: r}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int()
	return s.repository.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.repository.FindAll()
}
