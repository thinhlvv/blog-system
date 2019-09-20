package article

import "github.com/thinhlvv/blog-system/model"

// New return article controller service.
func New(app model.App) *Controller {
	repo := NewRepo(app.DB)
	svc := NewService(repo)
	c := NewController(svc, app)
	return c
}
