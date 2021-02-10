package services

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

type projectDAO interface {
	FindAll() []models.Project
}

type ProjectService struct {
	dao projectDAO
}

// NewProjectService creates a new ProjectService with the given project DAO.
func NewProjectService(dao projectDAO) *ProjectService {
	return &ProjectService{dao}
}

func (s *ProjectService) FindAll() ([]models.Project, error) {
	projects := s.dao.FindAll()
	err := fmt.Errorf("no projects found")
	if len(projects) > 0 {
		return projects, nil
	}
	return projects, err
}
