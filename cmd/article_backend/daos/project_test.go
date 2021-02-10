package daos

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewProjectDAO()

	projects := dao.FindAll()

	fmt.Printf("%v", projects)

	assert.Equal(t, 2, len(projects))
}
