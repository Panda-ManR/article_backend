package daos

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

// ProjectDAO persists project data in database
type ProjectDAO struct{}

// NewProjectDAO creates a new ProjectDAO
func NewProjectDAO() *ProjectDAO {
	return &ProjectDAO{}
}

func (dao *ProjectDAO) FindAll() []models.Project {
	var projects []models.Project
	config.Config.DB.Preload("Tags").Find(&projects)
	return projects
}
