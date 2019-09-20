package article

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thinhlvv/blog-system/model"
)

// Controller represents handler functions.
type Controller struct {
	service Service
}

// NewController returns controller endpoints.
func NewController(service Service) *Controller {
	return &Controller{service: service}
}

// GetByID return Article by ID.
func (ctrl Controller) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// GetAll return Article by ID.
func (ctrl Controller) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

type (
	createReq struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
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
