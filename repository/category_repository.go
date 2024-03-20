package repository

import "belajar-golang-unit-test/entity"


// implement contract interface in go
type CategoryRepository  interface {
	FindById(id string) *entity.Category
}