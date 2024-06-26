package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	// nil dan karena dia not found
	assert.Nil(t, category)
	// expect found
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
