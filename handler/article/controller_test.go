package article_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/handler/article"
	"github.com/thinhlvv/blog-system/model"
	"github.com/thinhlvv/blog-system/pkg"
	"github.com/thinhlvv/blog-system/testhelper"

	. "github.com/smartystreets/goconvey/convey"
)

var articleJSON = `{
	"title": "Title",
	"content": "content",
	"author":"thinh test"
}`

func TestControllerCreateArticle(t *testing.T) {
	// Setup
	db := testhelper.NewDB()
	e := echo.New()
	app := model.App{
		DB:               db,
		RequestValidator: pkg.NewRequestValidator(),
	}

	handler := article.New(app)

	Convey("Controller Scenario: Create Article successfully", t, func() {
		req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)
		So(err, ShouldBeNil)
		So(rec.Code, ShouldEqual, http.StatusCreated)
	})

	Convey("Controller Scenario: Get Article By ID successfully", t, func() {
		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := handler.GetByID(c)
		So(err, ShouldBeNil)
		So(rec.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Controller Scenario: Get Article By ID successfully", t, func() {
		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetAll(c)
		So(err, ShouldBeNil)
		So(rec.Code, ShouldEqual, http.StatusOK)
	})

}
