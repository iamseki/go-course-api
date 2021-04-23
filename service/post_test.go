package service_test

import (
	"testing"

	"github.com/iamseki/go-course-api/entity"
	"github.com/iamseki/go-course-api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestCreate(t *testing.T) {
	mockRepo := &MockRepository{}
	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	// Setup Expectations from mocked methods
	mockRepo.On("Save").Return(&post, nil)

	sut := service.NewPostService(mockRepo)

	result, err := sut.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	mockRepo := &MockRepository{}

	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	// Setup Expectations from mocked methods
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	sut := service.NewPostService(mockRepo)
	result, err := sut.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
	assert.Equal(t, 1, result[0].ID)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	sut := service.NewPostService(nil)

	err := sut.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	sut := service.NewPostService(nil)
	post := &entity.Post{ID: 1, Text: "Fake Test"}

	err := sut.Validate(post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestValidateSuccess(t *testing.T) {
	sut := service.NewPostService(nil)
	post := &entity.Post{ID: 1, Title: "Fake title", Text: "Fake Test"}

	err := sut.Validate(post)

	assert.Nil(t, err)
}
