package article

import (
	"net/http"
	"strconv"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/model"
	"github.com/thinhlvv/blog-system/pkg"
)

// Controller represents handler functions.
type Controller struct {
	service   Service
	validator pkg.RequestValidator
}

// NewController returns controller endpoints.
func NewController(service Service, app model.App) *Controller {
	return &Controller{
		service:   service,
		validator: app.RequestValidator,
	}
}

// GetByID return Article by ID.
func (ctrl Controller) GetByID(c echo.Context) error {
	articleID := c.Param("id")
	if _, err := strconv.Atoi(articleID); err != nil {
		return c.JSON(http.StatusBadRequest, model.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	article, err := ctrl.service.GetArticleByID(articleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, model.BaseResponse{
				Status:  http.StatusOK,
				Message: err.Error(),
				Data:    nil,
			})
		}
		return c.JSON(http.StatusUnprocessableEntity, model.BaseResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    article,
	})
}

// GetAll return Article by ID.
func (ctrl Controller) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

type (
	createReq struct {
		Title   string `json:"title" validate:"max=255,min=1"`
		Content string `json:"content" validate:"min=1"`
		Author  string `json:"author" validate:"max=255,min=1"`
	}

	createRes struct {
		ID int `json:"id"`
	}
)

// Create return Article by ID.
func (ctrl Controller) Create(c echo.Context) error {
	var req createReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := ctrl.validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	article, err := ctrl.service.CreateArticle(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.BaseResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    createRes{ID: article.ID},
	})
}
