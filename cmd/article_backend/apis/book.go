package apis

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/daos"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetBooks is function for endpoint /api/v1/books to get all books
func GetBooks(c *gin.Context) {
	s := services.NewBookService(daos.NewBookDAO())
	if books, err := s.FindAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	}
}
