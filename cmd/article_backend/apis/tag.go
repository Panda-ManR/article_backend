package apis

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/daos"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetTags is function for endpoint /api/v1/tags to get all tags by post_id
func GetTags(c *gin.Context) {
	s := services.NewTagService(daos.NewTagDAO())
	id, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if tags, err := s.FindAll(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tags": tags,
		})
	}
}
