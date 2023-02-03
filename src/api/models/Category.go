package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories []Category

func NewCategory(db *sql.DB) *Category {
	return &Category{}
}

func (s *Category) GetAllCategories() Categories {
	var categories Categories
	Connection().Find(&categories)
	return categories
}

func (s *Category) CreateCategory() Category {
	s.ID = uuid.New().String()
	Connection().Create(&s)
	return *s
}

func (s *Category) DeleteCategory() Category {
	Connection().Delete(&s)
	return *s
}

func (s *Category) UpdateCategory() Category {
	Connection().Save(&s)
	return *s
}
