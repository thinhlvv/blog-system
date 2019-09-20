package article_test

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thinhlvv/blog-system/handler/article"
	"github.com/thinhlvv/blog-system/testhelper"
)

func TestService(t *testing.T) {
	db := testhelper.NewDB()
	repo := article.NewRepo(db)
	service := article.NewService(repo)

	Convey("Service Feature: Create Article", t, func() {
		req := article.CreateReq{
			Title:   "Title",
			Content: "Content",
			Author:  "Author",
		}

		createdArt, err := service.CreateArticle(req)
		So(err, ShouldBeNil)

		getArt, err := service.GetArticleByID(strconv.Itoa(createdArt.ID))
		So(err, ShouldBeNil)
		So(getArt.ID, ShouldEqual, createdArt.ID)

		_, err = service.GetAllArticles()
		So(err, ShouldBeNil)

		Reset(func() {
			db.Close()
		})
	})
}
