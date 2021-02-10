package daos

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewTagDAO()

	tags := dao.FindAll(1)

	assert.Equal(t, 3, len(tags))
}

func TestTagDAO_FindEmpty(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewTagDAO()

	tags := dao.FindAll(9999)

	assert.Empty(t, tags)
}
