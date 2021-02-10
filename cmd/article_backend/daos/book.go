package daos

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

// BookDAO persists book data in database
type BookDAO struct{}

// NewBookDAO creates a new BookDAO
func NewBookDAO() *BookDAO {
	return &BookDAO{}
}

func (dao *BookDAO) FindAll() []models.Book {
	var books []models.Book
	config.Config.DB.Find(&books)
	return books
}
