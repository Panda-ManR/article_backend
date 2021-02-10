package daos

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

// TagDAO persists tag data in database
type SectionDAO struct{}

// NewSectionDAO creates a new SectionDAO
func NewSectionDAO() *SectionDAO {
	return &SectionDAO{}
}

func (dao *SectionDAO) FindAll(postId uint) []models.Section {
	var sections []models.Section
	config.Config.DB.Where("post_id = ?", postId).Find(&sections)
	return sections
}
