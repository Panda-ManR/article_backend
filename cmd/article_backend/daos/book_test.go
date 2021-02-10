package daos

import (
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewBookDAO()

	books := dao.FindAll()

	assert.Equal(t, 4, len(books))
}
