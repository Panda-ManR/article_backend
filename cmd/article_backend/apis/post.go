package apis

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/daos"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetPost is function for endpoint /api/v1/posts to get Post by ID
func GetPost(c *gin.Context) {
	s := services.NewPostService(daos.NewPostDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if post, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// GetPosts is function for endpoint /api/v1/posts to get all posts
func GetPosts(c *gin.Context) {
	s := services.NewPostService(daos.NewPostDAO())
	if posts, err := s.FindAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	}
}
