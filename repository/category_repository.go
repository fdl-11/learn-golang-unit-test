package repository

import "golang-unit-test/entity"

// kontrak
type CategoryRepository interface {
	FindById(id string) *entity.Category		// kalau ada isi, akan return category, kalau tidak ada akan return nil
}