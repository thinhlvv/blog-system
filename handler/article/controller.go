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
		model.BaseResponse
	}
)

// Create return Article by ID.
func (ctrl Controller) Create(c echo.Context) error {
	// validate request
	// service call
	return c.JSON(http.StatusOK, nil)
}
