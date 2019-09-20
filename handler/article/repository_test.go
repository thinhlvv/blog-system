package article_test

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/blog-system/handler/article"
	"github.com/thinhlvv/blog-system/testhelper"
)

func TestRepository(t *testing.T) {
	db := testhelper.NewDB()

	Convey("Repository Feature: CRUD article", t, func() {
		art := article.Article{
			Title:   "Title",
			Content: "Content",
			Author:  "Author",
		}
		repo := article.NewRepo(db)

		createdArtID, err := repo.CreateArticle(art)
		So(err, ShouldBeNil)

		getArt, err := repo.GetArticleByID(strconv.Itoa(createdArtID))
		So(err, ShouldBeNil)
		So(getArt.ID, ShouldEqual, createdArtID)

		_, err = repo.GetAllArticles()
		So(err, ShouldBeNil)

		Reset(func() {
			db.Close()
		})
	})
}
