package services

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

type bookDAO interface {
	FindAll() []models.Book
}

type BookService struct {
	dao bookDAO
}

// NewBookService creates a new BookService with the given book DAO.
func NewBookService(dao bookDAO) *BookService {
	return &BookService{dao}
}

func (s *BookService) FindAll() ([]models.Book, error) {
	books := s.dao.FindAll()
	err := fmt.Errorf("no books found")
	if len(books) > 0 {
		return books, nil
	}
	return books, err
}
