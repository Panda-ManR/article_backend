package services

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

type sectionDAO interface {
	FindAll(id uint) []models.Section
}

type SectionService struct {
	dao sectionDAO
}

// NewSectionService creates a new SectionService with the given section DAO.
func NewSectionService(dao sectionDAO) *SectionService {
	return &SectionService{dao}
}

func (s *SectionService) FindAll(postId uint) ([]models.Section, error) {
	sections := s.dao.FindAll(postId)
	err := fmt.Errorf("no sections found for post_id: %v", postId)
	if len(sections) > 0 {
		return sections, nil
	}
	return sections, err
}
